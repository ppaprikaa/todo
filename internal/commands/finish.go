package commands

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sanzhh/todo/internal/db"
	"github.com/sanzhh/todo/internal/ports"
	"github.com/sanzhh/todo/internal/storage"
	"github.com/spf13/cobra"
)

var (
	finishId string
)

var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "to mark a task as done.",
	Long:  "this subcommand exists in order to mark a task as done.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if strings.TrimSpace(finishId) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "ID: ")
			finishId, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(finishId) == "" {
			_ = cmd.Help()
			return
		}

		db, err := db.NewSQLite()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "error: %v\n", err)
			return
		}
		defer db.Close()

		storage := storage.NewSQLite(db)

		finishCtx, finishCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer finishCancel()

		done := new(bool)
		*done = true

		if err = storage.Update(
			finishCtx,
			ports.UpdateDTO{
				ID:   strings.TrimSpace(finishId),
				Done: done,
			},
		); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "error: %v\n", err)
		}
	},
}

func init() {
	finishCmd.Flags().StringVarP(&finishId, "id", "i", "", "ID of the todo to be updated.")
}
