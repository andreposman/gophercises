package _1_quiz

import (
	"bufio"
	"fmt"
	"github.com/andreposman/gophercises/01-quiz/domain"
	"github.com/andreposman/gophercises/pkg/utils"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type QuizCSV struct {
	Operation string `csv:"operation"`
	Result    string `csv:"result"`
}

func Quiz() {
	utils.PrintName("01", "Quiz Game")
	filePath := "01-quiz/"
	fileName := "prob.csv"
	correctAnswers := 0
	incorrectAnswers := 0

	data, err := utils.ReadCSV(filePath, fileName)
	if err != nil {
		log.Error().Msgf("Error reading file: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("------------------------")
	fmt.Println("Starting Game")
	fmt.Println("------------------------")

	for {
		quizData, _ := parseData(data)
		for i := 0; i <= len(quizData.Question)-1; i++ {
			fmt.Printf("\n-> Question: How much is %v:\n", quizData.Question[i])
			fmt.Printf("-> ")
			text, _ := reader.ReadString('\n')
			text = strings.Trim(text, "\n")

			if text == quizData.Answer[i] {
				correctAnswers++
				if i == len(quizData.Question)-1 {
					printFinalMessage(correctAnswers, incorrectAnswers, quizData)
					fmt.Println()
					os.Exit(0)
				}
			} else {
				incorrectAnswers++
				if i == len(quizData.Question)-1 {
					printFinalMessage(correctAnswers, incorrectAnswers, quizData)
					fmt.Println()
					os.Exit(0)
				}
			}
		}

	}

}

func printFinalMessage(correctAnswers, incorrectAnswers int, quizData domain.QuizCSV) {
	fmt.Printf("\n\n------------------------------------------------------------------------------------------------\n")
	if correctAnswers >= incorrectAnswers {
		fmt.Printf("\n             Contratulations! You got %d/%d of the questions right!\n", correctAnswers, len(quizData.Question))
	}
	if correctAnswers < incorrectAnswers {
		fmt.Printf("\n             Sorry! You only got %d/%d of the questions right.\n", correctAnswers, len(quizData.Question))
	}
	fmt.Printf("\n------------------------------------------------------------------------------------------------")

}

func parseData(data [][]string) (domain.QuizCSV, error) {
	var quizData domain.QuizCSV
	for _, item := range data {
		for pos, str := range item {
			if pos == 0 {
				quizData.Question = append(quizData.Question, str)
			} else {
				quizData.Answer = append(quizData.Answer, str)
			}
		}
	}
	return quizData, nil
}
