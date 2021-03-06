package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// CsvReader type
type CsvReader struct {
	file     io.Reader
	Filename string
}

func (c *CsvReader) open() {
	name := flag.String("csv", c.Filename, "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*name)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *name))
	}
	c.file = file
}

func (c *CsvReader) read() {
	c.open()
	r := csv.NewReader(c.file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// Problem consist of question and answer
type Problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
