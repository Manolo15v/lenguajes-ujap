package cmd

import (
	"fmt"
	"log"
	"strings"

	"task/pkg/models"

	"github.com/spf13/cobra"
)

var priorityStr string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long:  "Add a new task with description and optional priority (high, medium, low)",
	Args:  cobra.MinimumNArgs(1),
	Run:   addTask,
}

func init() {
	addCmd.Flags().StringVarP(&priorityStr, "priority", "p", "medium", "task priority (high, medium, low)")
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	description := strings.Join(args, " ")

	conn := GetDb()

	priority := models.Priority(priorityStr)

	if !models.IsValidPriority(priority) {
		log.Fatalf("Invalid priority: %s. Must be one of: high, medium, low", priorityStr)
	}

	task := models.NewTask(description, priority)

	id, err := conn.InsertTask(task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added task with ID: %d\n", id)
}
