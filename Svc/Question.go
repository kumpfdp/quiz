package Svc

import (
	"fmt"
	"strings"
)

type Question interface {
	AskQuestion()
	EvaluateAnswer(userAnswer string) bool
}

type question struct {
	Question string
	Answer   string
}

func NewQuestion(q, a string) Question {
	return &question{
		Question: q,
		Answer:   a,
	}
}

func (q question) EvaluateAnswer(userAnswer string) bool {
	return strings.Compare(strings.ToUpper(q.Answer), strings.ToUpper(userAnswer)) == 0
}

func (q question) AskQuestion() {
	fmt.Printf("%v ", q.Question)
}
