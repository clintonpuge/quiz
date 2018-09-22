package main

import (
	"github.com/clintonpuge/math-quiz/app"
)

func main() {
	csv := app.CsvReader{
		Filename: "problems.csv",
	}
	app.ReadQuestions(&csv)
}
