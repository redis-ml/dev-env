package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	internalBatchSize = 1000

	ProducerBatchSize = 10
)

var isDebugMode = false

type FileSubset struct {
	ObjectLocation string `json:"object_location,omitempty"`
	ObjectSize     int64  `json:"object_size,omitempty"`
	OffsetStart    int64  `json:"offset_start,omitempty"`
	OffsetEnd      int64  `json:"offset_end,omitempty"`
}

// Size:
//  - object_size: int64
//  - consumer batch: 1000 (SQS event batch size)
//  - internal batch size: 1000 (go routine)

type TaskSplit struct {
	Locations []FileSubset `json:"locations,omitempty"`
	Payload   CommEvent    `json:"payload,omitempty"`
	EventID   string       `json:"event_id,omitempty"`
	Start     int64        `json:"start,omitempty"`
	End       int64        `json:"end,omitempty"`
}

func splitWork(ctx context.Context, parent TaskSplit) error {
	subTasks := getSubTasks(ctx, parent)
	return scheduleSubTask(ctx, subTasks)
}

func getSubTasks(ctx context.Context, parent TaskSplit) []TaskSplit {
	num := (parent.End-parent.Start)/internalBatchSize + 1
	splits := make([]TaskSplit, 0, num)
	for i := int64(0); i < num; i++ {
		increment := int64(i) * internalBatchSize
		newStart, newEnd := parent.Start+increment, parent.Start+increment+internalBatchSize
		if newStart >= parent.End {
			break
		}
		if newEnd > parent.End {
			newEnd = parent.End
		}
		split := newTaskSplit(parent, newStart, newEnd)
		splits = append(splits, split)
	}

	return splits
}

func newTaskSplit(parent TaskSplit, start, end int64) TaskSplit {
	// TODO: double check if we need better deep-copy.
	newTask := parent
	// Update start and end.
	newTask.Start = start
	newTask.End = end
	return newTask
}

func scheduleSubTask(ctx context.Context, tasks []TaskSplit) error {
	payloads := make([]string, 0, len(tasks))
	for _, task := range tasks {
		p, err := json.Marshal(task)
		if err != nil {
			log.Printf("[ERROR] failed to marshal task: %+v err: %v", task, err)
			return err
		}
		payloads = append(payloads, string(p))
	}

	if isDebugMode {
		log.Printf("[DEBUG] local test, skip sending to SQS")
		log.Printf("[DEBUG] payloads: %+v\n", payloads)
		return nil
	}

	sqsClient := newSQSClient()

	entries := resetEntries()
	for i, payload := range payloads {
		if i > 0 && i%ProducerBatchSize == 0 {
			err := sendBatch(ctx, sqsClient, entries)
			if err != nil {
				log.Printf("[ERROR] failed to send batch: %+v err: %v", entries, err)
				return err
			}
			entries = resetEntries()
		}
		entries = append(entries, newBatchEntry(i, payload))
	}
	// Handle remaining entries if any.
	err := sendBatch(ctx, sqsClient, entries)
	if err != nil {
		log.Printf("[ERROR] failed to send batch: %+v err: %v", entries, err)
		return err
	}

	return nil
}

func resetEntries() []*sqs.SendMessageBatchRequestEntry {
	return make([]*sqs.SendMessageBatchRequestEntry, 0, ProducerBatchSize)
}

func newBatchEntry(id int, body string) *sqs.SendMessageBatchRequestEntry {
	return &sqs.SendMessageBatchRequestEntry{
		Id:          aws.String(fmt.Sprintf("%d", id)),
		MessageBody: aws.String(body),
	}
}

func sendBatch(ctx context.Context, sqsClient *sqs.SQS, entries []*sqs.SendMessageBatchRequestEntry) error {
	if len(entries) == 0 {
		return nil
	}
	log.Printf("to send %d entries, sample: %s\n", len(entries), *entries[0].MessageBody)
	batchInput := &sqs.SendMessageBatchInput{
		QueueUrl: aws.String(queueURL),
		Entries:  entries,
	}
	_, err := sqsClient.SendMessageBatchWithContext(ctx, batchInput)
	return err
}

func UnmarshalTask(ctx context.Context, input string) (TaskSplit, error) {
	var taskSplit TaskSplit
	err := json.Unmarshal([]byte(input), &taskSplit)
	return taskSplit, err
}
