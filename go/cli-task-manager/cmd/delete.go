/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task-id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run:   deleteTask,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTask(cmd *cobra.Command, args []string) {
	id, idErr := strconv.Atoi(args[0])
	if idErr != nil {
		log.Fatal(idErr)
	}

	conn := GetDb()

	deleteErr := conn.DeleteTask(id)
	if deleteErr != nil {
		log.Fatal(deleteErr)
	}
}
