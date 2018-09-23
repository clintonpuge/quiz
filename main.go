package main

import (
	"github.com/clintonpuge/math-quiz/app"
)

func main() {
	// csv := app.CsvReader{
	// 	Filename: "problems.csv",
	// }
	json := app.JSONReader{}
	app.ReadQuestions(&json)
}
