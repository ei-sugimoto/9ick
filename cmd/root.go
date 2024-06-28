package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "9ick",
	Short: "A CLI tool for generating project templates",
	Long:  `A CLI tool that helps you generate project templates quickly.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to mycli!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
