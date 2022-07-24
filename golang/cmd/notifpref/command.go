package notifpref

import (
	"context"
	"log"

	"github.com/redisliu/dev-env/golang/pkg/notifpref/handler"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "notifpref",
	Short: "notifpref",
	Long:  `notifpref`,
	Run:   notifPrefRun,
}

var (
	flagStartIndex uint64
	flagNum        uint64
)

func init() {
	Command.Flags().Uint64VarP(&flagStartIndex, "start", "s", 0, "start index of the user id")
	Command.Flags().Uint64VarP(&flagNum, "num", "n", 10, "number of users to be handled")
}

func notifPrefRun(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	log.Printf("start: %d, num: %d, %v", flagStartIndex, flagNum, ctx)
	for i := uint64(0); i < flagNum; i++ {
		userId := flagStartIndex + i
		out, err := handler.HandleRequest(ctx, userId)
		log.Printf("out: %s, err: %v", out, err)
	}
}
