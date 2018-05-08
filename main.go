package main

import (
	"github.com/kumpfdp/quiz/Svc"
)

const (
	csvPath = "problems.csv"
)

var correctAnswerCount int
var incorrectAnswerCount int

func main() {
	quiz := Svc.NewQuiz(csvPath)
	quiz.LoadQuestions()

	quiz.Start()

}
