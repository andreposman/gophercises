package _1_quiz

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/andreposman/gophercises/pkg/utils"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

type QuizCSV struct {
	Question []string
	Answer   []string
}

func Quiz() {
	filePath := "01-quiz/"
	fileName := "prob.csv"
	var quizTimeFlag int
	var duration time.Duration = 5
	isTimerDone := make(chan bool, 1)
	reader := bufio.NewReader(os.Stdin)

	flag.IntVar(&quizTimeFlag, "time", 0, "duration of the timer passed via flag")
	flag.Parse()

	if quizTimeFlag > 0 {
		duration = time.Duration(quizTimeFlag)
	}

	data, err := utils.ReadCSV(filePath, fileName)
	if err != nil {
		log.Error().Msgf("Error reading file: %v", err)
	}

	//cli for
	for {
		startGame(duration, reader)
		handleGameLogic(data, reader, duration, isTimerDone)
	}

}

func startGame(duration time.Duration, reader *bufio.Reader) {
	utils.PrintLineDivider(90)
	utils.PrintName("01", "Quiz Game")
	utils.PrintLineDivider(90)

	fmt.Printf("> PRESS ANY KEY TO START THE GAME!")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	fmt.Printf("\nYou will have %d seconds to finish the quiz, before the timer expires!\n", duration)
}

func handleUserInput(reader *bufio.Reader) string {
	fmt.Printf("> ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	return text
}

func handleGameLogic(data [][]string, reader *bufio.Reader, duration time.Duration, isTimerDone chan bool) {
	var correctAnswers, incorrectAnswers int
	gameStartTime := time.Now()
	quizData, _ := parseData(data)

	go quizTimer(duration, gameStartTime, isTimerDone, correctAnswers, incorrectAnswers, len(quizData.Question))

	for i := 0; i <= len(quizData.Question)-1; i++ {
		fmt.Printf("\n(%d) Question: How much is %v ?\n", i+1, quizData.Question[i])
		text := handleUserInput(reader)

		if text == quizData.Answer[i] {
			correctAnswers++
			handleGameOver(i, quizData, correctAnswers, incorrectAnswers, gameStartTime)
		} else {
			incorrectAnswers++
			handleGameOver(i, quizData, correctAnswers, incorrectAnswers, gameStartTime)
		}
	}
}

func quizTimer(duration time.Duration, gameStartTime time.Time, isTimerDone chan bool, correctAnswers, incorrectAnswers int, numQuestions int) {
	timer := time.NewTimer(duration * time.Second)
	<-timer.C
	select {
	case isTimerDone <- true:
		handleTimeExpired(gameStartTime, correctAnswers, incorrectAnswers, numQuestions, duration)
		close(isTimerDone)
	}
}

func handleTimeExpired(gameStartTime time.Time, correctAnswers, incorrectAnswers int, numQuestions int, duration time.Duration) {
	println()
	charCount := 90
	utils.PrintLineDivider(charCount)
	fmt.Printf(" Your time is over! You took %.1f seconds to finish.\n", time.Since(gameStartTime).Seconds())
	printFinalMessage(correctAnswers, incorrectAnswers, numQuestions, duration.Seconds())

	os.Exit(0)
}

func handleGameOver(i int, quizData QuizCSV, correctAnswers, incorrectAnswers int, gameStartTime time.Time) {
	var quizTotalTimeDuration time.Duration
	if i == len(quizData.Question)-1 {
		quizTotalTimeDuration = time.Since(gameStartTime)
		printFinalMessage(correctAnswers, incorrectAnswers, len(quizData.Question), quizTotalTimeDuration.Seconds())

		os.Exit(0)
	}
}

func printFinalMessage(correctAnswers, incorrectAnswers, numQuestions int, quizTotalTimeDuration float64) {
	var prefix string
	chartCount := 90

	utils.PrintLineDivider(chartCount)

	if correctAnswers > incorrectAnswers {
		prefix = strings.ToUpper(" congratulations!!")
	} else {
		prefix = "Maybe next time you can do better."
	}

	fmt.Printf(prefix+" You got %d out of %d of the questions right! You took %.2f seconds to finish.\n", correctAnswers, numQuestions, quizTotalTimeDuration)
	utils.PrintLineDivider(chartCount)
}

func parseData(data [][]string) (QuizCSV, error) {
	var quizData QuizCSV
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
