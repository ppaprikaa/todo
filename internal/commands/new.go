package commands

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ppaprikaa/todo/internal/db"
	"github.com/ppaprikaa/todo/internal/models"
	"github.com/ppaprikaa/todo/internal/ports"
	"github.com/ppaprikaa/todo/internal/storage"
	"github.com/spf13/cobra"
)

var (
	newName        string
	newDescription string
	newDate        string
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates todos",
	Long: `You can create todo using either flags or arguments.
You must fill name.`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if strings.TrimSpace(newName) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "Name: ")
			newName, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(newName) == "" {
			cmd.Help()
			return
		}

		if strings.TrimSpace(newDescription) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "Description(Skippable): ")
			newDescription, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		if strings.TrimSpace(newDate) == "" {
			scannr := bufio.NewReader(cmd.InOrStdin())
			fmt.Fprint(cmd.OutOrStdout(), "Date(Skippable): ")
			newDate, err = scannr.ReadString('\n')
			if err != nil {
				_ = cmd.Help()
				return
			}
		}

		db, err := db.NewSQLite()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "error: %v", err)
			return
		}
		defer db.Close()

		storage := storage.NewSQLite(db)

		newCtx, newCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer newCancel()

		data := ports.InsertDTO{
			Name:        newName,
			Description: newDescription,
		}

		if strings.TrimSpace(newDate) != "" {
			data.Date = &newDate
		}

		storage.Insert(newCtx, data)
	},
}

func init() {
	newCmd.Flags().StringVarP(&newName, "name", "n", "", "fill this to fill name")
	newCmd.Flags().StringVarP(&newDescription, "description", "b", "", "fill this to fill description")
	newCmd.Flags().StringVarP(&newDate, "date", "d", "", fmt.Sprintf("fill this to fill date (date format: %s)", models.TodoDateFormat))
}
