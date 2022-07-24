package main

import (
	"context"
	"log"

	"github.com/redisliu/dev-env/golang/cmd/notifpref"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func initConfig() {
	// TODO: implement
}

func init() {
	RootCmd.AddCommand(notifpref.Command)
}

func main() {
	ctx := context.Background()
	cobra.OnInitialize(initConfig)

	if err := RootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("cmd failed with err: %+v", err)
	}
}
