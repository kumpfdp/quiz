package main

import (
	"github.com/kumpfdp/quiz/Svc"
)

const (
	CsvPath = "problems.csv"
	TimeLimit = 10
)

var correctAnswerCount int
var incorrectAnswerCount int

func main() {
	quiz := Svc.NewQuiz(CsvPath, TimeLimit)
	quiz.LoadQuestions()

	quiz.Start()

}
