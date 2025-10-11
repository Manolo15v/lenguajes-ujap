/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [task-id]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run:   completeTask,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeTask(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	conn := GetDb()

	completeErr := conn.CompleteTask(id)
	if completeErr != nil {
		log.Fatal(completeErr)
	}
}
