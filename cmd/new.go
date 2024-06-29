package cmd

import (
	reader "github.com/ei-sugimoto/9ick/internal/pkg"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project from a template",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]

		if !checkTemplateName(templateName) {
			cmd.PrintErr("Template not found: ", templateName)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func checkTemplateName(templateName string) bool {
	r := reader.NewReader()
	templatesList, err := r.List()
	if err != nil {
		return false
	}
	for _, template := range templatesList {
		if template == templateName {
			return true
		}
	}
	return false
}
