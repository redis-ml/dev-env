package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/redisliu/dev-env/golang/cmd/sqsfanoutworker/handler"
)

var (
	flagFilename string
)

func main() {
	ctx := context.Background()

	flag.StringVar(&flagFilename, "filename", "", "filename")
	flag.Parse()

	eventList, err := load(ctx)
	if err != nil {
		log.Fatalf("failed to load event list: %+v", err)
	}

	// Tunes the debugging level.
	handler.SetDebugMode(true)
	handler.SetUseFakeSink(true)

	out, err := handler.HandleRequest(ctx, eventList)
	log.Printf("out: %s, err: %v", out, err)

	// wait for debugging.
	sleepDuration := 1 * time.Hour
	log.Printf("sleeping for %s, ctrl+c to break and quit", sleepDuration)
	time.Sleep(sleepDuration)
}

func load(ctx context.Context) (eventList events.SQSEvent, err error) {
	bb, err := os.ReadFile(flagFilename)
	if err != nil {
		return events.SQSEvent{}, fmt.Errorf("failed to read file: %+v", err)
	}
	err = json.Unmarshal(bb, &eventList)
	return
}
