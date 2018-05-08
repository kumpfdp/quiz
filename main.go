package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/kumpfdp/quiz/Svc"
	"log"
	"os"
	"strings"
)

const (
	DefaultCsvPath   = "problems.csv"
	DefaultTimeLimit = 30
)

func main() {
	f := flag.String("file", DefaultCsvPath, "Path to CSV quiz file")
	tl := flag.Int("timelimit", DefaultTimeLimit, "Time limit to complete the quiz")
	flag.Parse()

	fmt.Println(*f)

	quiz := Svc.NewQuiz(*f, *tl)
	err := quiz.LoadQuestions()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("This is a timed quiz. You have %v seconds to finish. \n", *tl)
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
