package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSONReader type
type JSONReader struct {
	file     string
	filename string
}

// Questions from json file
type Questions struct {
	Questions []Question `json:"questions"`
}

// Question key value
type Question struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func (j *JSONReader) read() {
	// open json file
	jsonFile, err := os.Open("problems.json")
	// check error
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read the json file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize questions list
	var questions Questions

	json.Unmarshal(byteValue, &questions)

	problems := parseRow(questions)

	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if problem.a == answer {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func parseRow(questions Questions) []Problem {
	problems := make([]Problem, len(questions.Questions))
	for i, question := range questions.Questions {
		problems[i] = Problem{
			q: question.Q,
			a: question.A,
		}
	}
	return problems
}
