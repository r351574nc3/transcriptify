package main

import (
	// "bufio"
	"encoding/csv"
	// "fmt"
	"io"
	"log"
	"os"
	"strconv"
	"github.com/r351574nc3/transcriptify/transcript"
)

const (
    CALY_COL   = 0
	YEAR_COL   = 1
	COURSE_COL = 2
	GRADE_COL  = 3
	CREDIT_COL = 4
	NA         = -1
)

var (
	records [][4]string
)

func main() {
	csvFile, err := os.Open("transcript.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvFile)

	first_rec := true
	t := transcript.Setup()
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

		year := t.GetGradeYear(record[YEAR_COL], record[CALY_COL])

		if year == nil {
			continue // Empty line
		}

		course    := record[COURSE_COL]
		grade     := record[GRADE_COL]
		credit, err := strconv.Atoi(record[CREDIT_COL])
		if err != nil {
			credit = NA
		}

		var c *transcript.Course
		c = c.New()
		c.Name   = course
		c.Grade  = grade
		c.Credit = credit
		year.Append(c)
		year.Recalc()
	}
	t.Recalc( )
	t.Process("tmpl/transcript.html.tmpl")

	if err := csvFile.Close(); err != nil {
		log.Fatal(err)
	}
}
