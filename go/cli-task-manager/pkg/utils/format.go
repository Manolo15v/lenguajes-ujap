package utils

import (
	"fmt"
	"strings"

	"task/pkg/models"
)

func FormatTaskList(tasks []models.Task, includeCompleted bool) string {
	if len(tasks) == 0 {
		return "No tasks found."
	}

	var result strings.Builder

	// Header
	result.WriteString(fmt.Sprintf("%-4s %-10s %-50s %-10s %s\n", "ID", "STATUS", "DESCRIPTION", "PRIORITY", "CREATED"))
	result.WriteString(strings.Repeat("-", 90) + "\n")

	// Tasks
	for _, task := range tasks {
		if !includeCompleted && task.State == models.StateCompleted {
			continue
		}

		result.WriteString(fmt.Sprintf("%-4d %-10s %-50s %-10s %s\n",
			task.ID,
			task.State,
			Truncate(task.Description, 50),
			task.Priority,
			task.CreatedAt.Format("2006-01-02")))
	}

	return result.String()
}

func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return ""
	}
	return s[:max-3] + "..."
}
