package handler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

const (
	queueURL = "https://us-west-2.queue.amazonaws.com/097605708335/fanout"
)

func HandleRequest(ctx context.Context, eventList events.SQSEvent) (string, error) {
	start := time.Now()

	var wg sync.WaitGroup
	var errHolder atomic.Value
	for i := range eventList.Records {
		wg.Add(1)
		go func(record *events.SQSMessage) {
			defer wg.Done()
			handleSingleEventWrapper(ctx, record, &errHolder)
			inc("event-cnt")
		}(&eventList.Records[i])
	}
	wg.Wait()

	if storedErr := errHolder.Load(); storedErr != nil {
		return "", fmt.Errorf("failed to handle event: %+v", storedErr)
	}
	duration := time.Since(start)
	return fmt.Sprintf("Processed msg: %d. time: %s", len(eventList.Records), duration), nil
}

func handleSingleEventWrapper(ctx context.Context, evt *events.SQSMessage, errHolder *atomic.Value) {
	task, err := UnmarshalTask(ctx, evt.Body)
	if err != nil {
		log.Printf("[ERROR] failed to unmarshal SQS task: %s, err: %+v\n", evt.Body, err)
		return
	}
	if err := handleSingleTask(ctx, &task); err != nil {
		errHolder.Store(err)
	}
}

func handleSingleTask(ctx context.Context, task *TaskSplit) error {
	if needSplit(ctx, task) {
		return splitWork(ctx, *task)
	} else {
		return handleLeafEventsInBatch(ctx, task)
	}
}

func needSplit(ctx context.Context, task *TaskSplit) bool {
	return task.End-task.Start > int64(float64(internalBatchSize)*1.6)
}

func handleLeafEventsInBatch(ctx context.Context, task *TaskSplit) error {
	log.Printf("event: %s, start: %d, end: %d\n", task.EventID, task.Start, task.End)
	var wg sync.WaitGroup
	n := task.End - task.Start
	if n <= 0 {
		return nil
	}
	errChan := make(chan error, n)
	// TODO: this is generated id.
	for id := task.Start; id < task.End; id++ {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()

			evt := instantiatecommEvent(ctx, task, fmt.Sprintf("%d", id))
			if err := handleSingleEvent(ctx, &evt); err != nil {
				errChan <- err
			}
			inc("comm-cnt")
		}(id)
	}

	wg.Wait()
	var err error
	select {
	case err = <-errChan:
	default:
	}

	return err
}

func instantiatecommEvent(ctx context.Context, task *TaskSplit, id string) (commEvent CommEvent) {
	commEvent = task.Payload
	commEvent.Owner = id

	// TODO: double check if this is valid assumption.
	commEvent.EventID = task.EventID
	commEvent.CampaignID = task.EventID

	return
}

func handleSingleEvent(ctx context.Context, evt *CommEvent) error {
	if isDebugMode {
		fmt.Printf("[DEBUG] handle event: %+v\n", evt)
	}

	// Mimic Idempotentcy checking.
	if _, err := checkIdempotency(ctx, evt); err != nil {
		log.Printf("[ERROR] failed to check idempotentcy for msg: %s(%s), err: %+v\n",
			evt.EventID, evt.Owner, err)
		return nil
	}

	if err := sendToSinkSqs(ctx, evt); err != nil {
		// Preent nothing happens, but log the error
		log.Printf("[WARN] failed to send msg for source id: %s(%s), err: %+v\n",
			evt.EventID, evt.Owner, err)
	}
	return nil
}
