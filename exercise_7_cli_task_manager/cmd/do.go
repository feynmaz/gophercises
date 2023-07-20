package cmd

import (
	"strings"

	"github.com/feynmaz/gophercises/exercise_7_cli_task_manager/storage"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  `Mark a task on your TODO list as complete`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			description := strings.Join(args, " ")
			storage.CompleteTask(description)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
