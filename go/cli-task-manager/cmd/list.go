/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"task/pkg/utils"

	"github.com/spf13/cobra"
)

var showAll bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List tasks with optional filtering by completion status",
	Run:   listTasks,
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show completed tasks too")
	rootCmd.AddCommand(listCmd)
}

func listTasks(cmd *cobra.Command, args []string) {
	tasks, err := GetDb().GetAllTasks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(utils.FormatTaskList(tasks, showAll))
}
