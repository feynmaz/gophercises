/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Launch quiz",
	Long:  `Launch quiz from csv file`,
	Run: func(cmd *cobra.Command, args []string) {
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
		doShuffle, _ := cmd.Flags().GetBool("shuffle")
		if doShuffle {
			dest := make([][]string, len(tasks))
			perm := rand.Perm(len(tasks))
			for i, v := range perm {
				dest[v] = tasks[i]
			}
			tasks = dest
		}
		limit, _ := cmd.Flags().GetInt("limit")
		if limit > 0 && limit < len(tasks) {
			tasks = tasks[:limit]
		}

		numCorrect := 0

		timeout, _ := cmd.Flags().GetInt("timeout")
		fmt.Printf("Timeout: %d seconds\n", timeout)

		timer := time.NewTimer(time.Duration(timeout) * time.Second)
		answerCh := make(chan string, 1)

	SolveQuiz:
		for idx, task := range tasks {
			go answer(idx, task, answerCh)

			select {

			case <-timer.C:
				fmt.Printf("Timeout exceeded: %d \n", timeout)
				close(answerCh)
				break SolveQuiz

			case word := <-answerCh:
				answer := strings.Trim(word, " ")
				if answer == task[1] {
					fmt.Println("The answer is right")
					numCorrect++

				} else {
					fmt.Println("The answer is wrong. Correct answer is " + task[1])
				}
			}
		}

		fmt.Printf("%d correct answers out of %d tasks \n", numCorrect, len(tasks))
	},
}

func answer(idx int, task []string, answerCh chan<- string) {
	promptWord := promptui.Prompt{
		Label: fmt.Sprintf("Problem #%d: %s", idx+1, task[0]),
	}
	word, err := promptWord.Run()
	if err != nil {
		log.Fatal("Prompt failed:", err)
		return
	}
	answerCh <- word
}

func init() {
	rootCmd.AddCommand(quizCmd)
	quizCmd.Flags().StringP("file", "f", "problems.csv", "Quiz file path. Defaults to 'problems.csv'")
	quizCmd.Flags().IntP("timeout", "t", 30, "Quiz solution time limit in seconds. Defaults to 30")
	quizCmd.Flags().BoolP("shuffle", "s", false, "Whether to shuffle the quiz order. Defaults to false")
	quizCmd.Flags().IntP("limit", "l", 0, "Count of tasks to set")
}
