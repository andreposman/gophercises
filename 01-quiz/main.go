package _1_quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andreposman/gophercises/01-quiz/domain"
	"github.com/andreposman/gophercises/pkg/utils"
	"github.com/rs/zerolog/log"
)

type QuizCSV struct {
	Operation string `csv:"operation"`
	Result    string `csv:"result"`
}

func Quiz() {
	utils.PrintName("01", "Quiz Game")
	filePath := "01-quiz/"
	fileName := "prob.csv"

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
				fmt.Println("You are correct!")
				if i == len(quizData.Question)-1 {
					fmt.Println("------------------------------------------------")
					fmt.Println("Congratulations! You reached the end.")
					fmt.Println("------------------------------------------------")
					os.Exit(0)
				}
			} else {
				fmt.Println("You are wrong! Try again")
				break
			}

		}
	}

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
	//log.Info().Msgf("Quiz - Questions: %v", quizData.Question)
	//log.Info().Msgf("Quiz - Answers: %v", quizData.Answer)
	return quizData, nil
}
