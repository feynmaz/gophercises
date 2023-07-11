/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Launch quiz",
	Long:  `Launch quiz from csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("quiz called")

		inputFile, _ := cmd.Flags().GetString("file")
		log.Println("file: " + inputFile)

		f, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		csvReader := csv.NewReader(f)
		tasks, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		numCorrect := 0
		for idx, task := range tasks {
			promptWord := promptui.Prompt{
				Label: fmt.Sprintf("Problem #%d: %s", idx+1, task[0]),
			}
			word, err := promptWord.Run()
			if err != nil {
				log.Fatal("Prompt failed:", err)
				return
			}
			if word == task[1] {
				fmt.Println("The answer is right")
				numCorrect++

			} else {
				fmt.Println("The answer is wrong. Correct answer is " + task[1])
			}
		}
		fmt.Printf("%d correct answers out of %d tasks \n", numCorrect, len(tasks))
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
	quizCmd.Flags().StringP("file", "f", "problems.csv", "Quiz file path. Defaults to 'problems.csv'")
}
