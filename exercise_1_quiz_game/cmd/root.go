/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "exercise_1_quiz_game",
	Short: "Exercise #1: Quiz Game",
	Long: `Create a program that will read in a quiz provided via a CSV file
	and will then give the quiz to a user keeping track 
	of how many questions they get right and how many they get incorrect`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


