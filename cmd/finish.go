/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish a task.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finish called")
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)
}
