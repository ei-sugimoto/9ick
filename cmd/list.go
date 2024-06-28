package cmd

import (
	"fmt"

	reader "github.com/ei-sugimoto/9ick/internal/template"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Run: func(cmd *cobra.Command, args []string) {
		r := reader.NewReader()

		templates, err := r.List()
		if err != nil {
			fmt.Println("Error listing templates:", err)
		}
		for _, tmpl := range templates {
			fmt.Println(tmpl)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
