package main

import (
	// "bufio"
	"encoding/csv"
	// "fmt"
	"io"
	"log"
	"os"
	// "strings"
)

const (
	YEAR_COL   = 0
	COURSE_COL = 1
	GRADE_COL  = 2
	CREDIT_COL = 3
)

func main() {
	csvFile, err := os.Open("transcript.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvFile)

	first_rec := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if first_rec {
			first_rec = false
			continue
		}

		year   := record[YEAR_COL]
		course := record[COURSE_COL]
		grade  := record[GRADE_COL]
		credit := record[CREDIT_COL]

		log.Printf("%s,%s,%s,%s\n", year, course, grade, credit)
	}

	log.Println("Works!")
	if err := csvFile.Close(); err != nil {
		log.Fatal(err)
	}
}
