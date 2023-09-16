package commands

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ppaprikaa/todo/internal/db"
	"github.com/ppaprikaa/todo/internal/storage"
	"github.com/spf13/cobra"
)

var deleteID string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo.",
	Long:  "Delete todo by id using args or flags.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if strings.TrimSpace(deleteID) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			deleteID, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(deleteID) == "" {
			_ = cmd.Help()
			return
		}

		db, err := db.NewSQLite()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer db.Close()

		storage := storage.NewSQLite(db)

		deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer deleteCancel()

		if err := storage.Delete(deleteCtx, deleteID); err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "error: %v", err)
			return
		}

		fmt.Fprint(cmd.OutOrStdout(), "[!] DELETED\n")
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "specify todo to delete by it's ID.")
}
