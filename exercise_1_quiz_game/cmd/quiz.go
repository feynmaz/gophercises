/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
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

		numCorrect := 0

		timeout, _ := cmd.Flags().GetInt("timeout")
		fmt.Printf("Timeout: %d seconds\n", timeout)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()

	SolveQuiz:
		for idx, task := range tasks {
			select {

			case <-ctx.Done():
				fmt.Printf("Timeout exceeded: %d \n", timeout)
				break SolveQuiz

			default:
				promptWord := promptui.Prompt{
					Label: fmt.Sprintf("Problem #%d: %s", idx+1, task[0]),
				}
				word, err := promptWord.Run()
				if err != nil {
					log.Fatal("Prompt failed:", err)
					return
				}

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

func init() {
	rootCmd.AddCommand(quizCmd)
	quizCmd.Flags().StringP("file", "f", "problems.csv", "Quiz file path. Defaults to 'problems.csv'")
	quizCmd.Flags().IntP("timeout", "t", 30, "Quiz solution time limit in seconds. Defaults to 30")
	quizCmd.Flags().BoolP("shuffle", "s", false, "Whether to shuffle the quiz order. Defaults to false")
}
