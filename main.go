package main

import (
	"bufio"
	"fmt"
	"github.com/kumpfdp/quiz/Svc"
	"log"
	"os"
	"strings"
)

const (
	CsvPath   = "problems.csv"
	TimeLimit = 10
)

func main() {
	quiz := Svc.NewQuiz(CsvPath, TimeLimit)
	quiz.LoadQuestions()

	fmt.Printf("This is a timed quiz. You have %v seconds to finish. \n", TimeLimit)
	fmt.Print("Ready to begin? (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	r, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if strings.Trim(r, "\n") == "n" {
		fmt.Println("Ending the quiz... ")
		return
	}

	quiz.Start()

}
