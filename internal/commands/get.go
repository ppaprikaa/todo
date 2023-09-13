package commands

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sanzhh/todo/internal/db"
	"github.com/sanzhh/todo/internal/output"
	"github.com/sanzhh/todo/internal/storage"
	"github.com/spf13/cobra"
)

var (
	getID     string
	getDate   string
	getOffset int
	getLimit  int
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get todos.",
	Long: `Get todos by id, date or offset and limit.
Hierarchy: id > date > offset & limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		DefID, DefDate := true, true

		db, err := db.NewSQLite()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
			return
		}
		defer db.Close()

		storage := storage.NewSQLite(db)

		for DefDate || DefID {
			switch {
			case strings.TrimSpace(getID) != "":
				getCtx, getCancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer getCancel()
				todo, err := storage.GetByID(getCtx, strings.TrimSpace(getID))
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
					return
				}

				res, err := output.DescribeTodo(todo)
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
					return
				}

				fmt.Fprintln(cmd.OutOrStdout(), res)

				return
			case strings.TrimSpace(getDate) != "":
				getCtx, getCancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer getCancel()
				todos, err := storage.GetByDate(getCtx, strings.TrimSpace(getDate))
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
					return
				}

				res, err := output.Todos(todos)
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
					return
				}

				fmt.Fprintln(cmd.OutOrStdout(), res)
				return
			case getOffset >= 0 && getLimit >= 0:
				getCtx, getCancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer getCancel()

				todos, err := storage.Get(getCtx, uint64(getOffset), uint64(getLimit))
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
					return
				}

				res, err := output.Todos(todos)
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
					return
				}

				fmt.Fprintln(cmd.OutOrStdout(), res)
				return
			case strings.TrimSpace(getID) == "" && DefID:
				fmt.Fprint(cmd.OutOrStdout(), "ID: ")
				scannr := bufio.NewReader(cmd.InOrStdin())

				getID, err = scannr.ReadString('\n')
				if strings.TrimSpace(getID) == "" {
					DefID = false
				}
			case strings.TrimSpace(getDate) == "" && DefDate:
				fmt.Fprint(cmd.OutOrStdout(), "Date: ")
				scannr := bufio.NewReader(cmd.InOrStdin())

				getDate, err = scannr.ReadString('\n')
				if strings.TrimSpace(getDate) == "" {
					DefDate = false
				}
			}
		}
	},
}

func init() {
	getCmd.Flags().StringVarP(&getID, "id", "i", "", "specify id to fetch todo")
	getCmd.Flags().StringVarP(&getDate, "date", "d", "", "specify date to fetch todos")
	getCmd.Flags().IntVarP(&getOffset, "offset", "o", -1, "specify offset")
	getCmd.Flags().IntVarP(&getLimit, "limit", "l", -1, "specify limit")
}
