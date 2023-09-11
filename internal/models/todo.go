package models

const (
	TodoDateFormat = "02-01-2006"
)

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Done        bool   `json:"done"`
}
