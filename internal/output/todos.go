package output

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/sanzhh/todo/internal/models"
)

func Todos(todos []models.Todo) (string, error) {
	var (
		format = "%s\t%s\t%s\t%v\t\n"
		buf    = &bytes.Buffer{}
		tw     = tabwriter.NewWriter(buf, 0, 8, 2, ' ', 0)
	)

	_, err := fmt.Fprintf(tw, format, "ID", "Name", "Date", "Done")
	if err != nil {
		return "", err
	}
	_, err = fmt.Fprintf(tw, format, "--", "----", "----", "----")
	if err != nil {
		return "", err
	}

	for _, todo := range todos {
		_, err = fmt.Fprintf(tw, format, todo.ID, todo.Name, todo.Date, todo.Done)
		if err != nil {
			return "", err
		}
	}
	err = tw.Flush()
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func DescribeTodo(todo *models.Todo) (string, error) {
	var (
		buf = &bytes.Buffer{}
	)

	metadata := fmt.Sprintf("- ID: %s - Name: %s - Date: %s - Done: %v -\n", todo.ID, todo.Name, todo.Date, todo.Done)
	_, err := fmt.Fprintf(buf, metadata)
	if err != nil {
		return "", err
	}

	delimeterBuilder := strings.Builder{}
	metarunes := []rune(metadata)

	for i := 0; i < (len(metarunes)-11)/2; i++ {
		delimeterBuilder.WriteRune('-')
	}
	delimeterBuilder.WriteString("DESCRIPTION")
	// -11 for "DESCRIPTION" and -1 for \n in metadata
	for i := 0; i < (len(metarunes)-11-1)/2; i++ {
		delimeterBuilder.WriteRune('-')
	}

	_, err = fmt.Fprintf(buf, "%s\n", delimeterBuilder.String())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(buf, "\t%s\n", todo.Description)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
