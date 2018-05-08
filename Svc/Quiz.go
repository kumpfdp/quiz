package Svc

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

type Quiz interface {
	LoadQuestions()
	NumberOfQuestions() int
	Start()
}

type quiz struct {
	QuestionCsvPath    string
	Questions          []Question
	CorrectAnswerCount int
	TimeLimit          int
}

func NewQuiz(questionCsvPath string, timeLimit int) Quiz {
	return &quiz{
		QuestionCsvPath: questionCsvPath,
		TimeLimit:       timeLimit,
	}
}

func (q *quiz) LoadQuestions() {
	// read in csv
	file, err := os.Open(q.QuestionCsvPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// build out question and answers
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		q.Questions = append(q.Questions, NewQuestion(record[0], record[1]))
	}
}

func (q *quiz) NumberOfQuestions() int {
	return len(q.Questions)
}

func (q *quiz) Start() {
	quizChan := make(chan bool)
	timerChan := make(chan bool)

	go StartTimer(time.Second*time.Duration(q.TimeLimit), timerChan)
	go StartQuiz(q, quizChan)

	select {
	case <-timerChan:
		fmt.Println("\n *** Time's up! *** ")
		break
	case <-quizChan:
		break
	}

	fmt.Printf("%v of %v correct!", q.CorrectAnswerCount, q.NumberOfQuestions())
}

func StartTimer(d time.Duration, ch chan<- bool) {
	defer close(ch)

	timerChan := time.NewTimer(d).C
	<-timerChan

	ch <- true
}

func StartQuiz(q *quiz, ch chan<- bool) {
	defer close(ch)

	fmt.Println("Good luck!")

	for _, quest := range q.Questions {
		quest.AskQuestion()
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
			break
		}
		answer = strings.Trim(answer, "\n")

		// evaluate the result
		if quest.EvaluateAnswer(answer) {
			q.CorrectAnswerCount = q.CorrectAnswerCount + 1
		}
	}

	ch <- true
}
