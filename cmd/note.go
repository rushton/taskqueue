/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Adds a note to a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("note called")
	},
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
