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
	updateId          string
	updateName        string
	updateDescription string
	updateDate        string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update todo.",
	Long:  "Update todo by specifying id and filling some fields.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if strings.TrimSpace(updateId) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "ID: ")
			updateId, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(updateId) == "" {
			_ = cmd.Help()
			return
		}

		if strings.TrimSpace(updateName) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "Name: ")
			updateName, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(updateName) == "" {
			cmd.Help()
			return
		}

		if strings.TrimSpace(updateDescription) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "Description(Skippable): ")
			updateDescription, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(updateDate) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "Date(Skippable): ")
			updateDate, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		updateId = strings.TrimSpace(updateId)
		updateName = strings.TrimSpace(updateName)
		updateDescription = strings.TrimSpace(updateDescription)
		updateDate = strings.TrimSpace(updateDate)

		data := &ports.UpdateDTO{}

		if updateId != "" {
			data.ID = updateId
		}

		if updateName != "" {
			data.Name = &updateName
		}

		if updateDescription != "" {
			data.Description = &updateDescription
		}

		if updateDate != "" {
			data.Date = &updateDate
		}

		db, err := db.NewSQLite()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
			return
		}
		defer db.Close()

		storage := storage.NewSQLite(db)

		updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer updateCancel()

		err = storage.Update(updateCtx, *data)
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateId, "id", "i", "", "ID of the todo to be updated.")
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "Name for todo.")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "b", "", "Description for todo.")
	updateCmd.Flags().StringVarP(&updateDate, "date", "d", "", "Date for todo.")
}
