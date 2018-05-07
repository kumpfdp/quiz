package quiz

import (
	"github.com/kumpfdp/quiz/Quiz"
)

const (
	csvPath = "problems.csv"
)

var correctAnswerCount int
var incorrectAnswerCount int

func main() {
	quiz := Quiz.NewQuiz(csvPath)
	quiz.LoadQuestions()

	quiz.Start()

}
