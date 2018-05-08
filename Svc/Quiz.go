package Svc

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var reader bufio.Reader

type Quiz interface {
	LoadQuestions()
	NumberOfQuestions() int
	Start()
}

type quiz struct {
	QuestionCsvPath    string
	Questions          []Question
	CorrectAnswerCount int
}

func NewQuiz(questionCsvPath string) Quiz {
	return &quiz{
		QuestionCsvPath: questionCsvPath,
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
	fmt.Println("Starting the quiz... Good luck!")

	reader := bufio.NewReader(os.Stdin)

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

	fmt.Printf("%v of %v correct!", q.CorrectAnswerCount, q.NumberOfQuestions())
}
