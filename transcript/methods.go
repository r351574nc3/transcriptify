package transcript

import (
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"text/template"
)

var (
	tmpl *template.Template
)

func (t *Transcript) New() *Transcript {
	if t == nil {
		return new(Transcript).Init()
	}
	return t
}

func (t *Transcript) Init() *Transcript {
	t.GradeYears = make([]*GradeYear, 0)
	return t
}

func (t *Transcript) Append(year *GradeYear) *Transcript {
	t.GradeYears = append(t.GradeYears, year)
	return t
}

func (t *Transcript) Process(filenames ...string) *Transcript {
	tmpl = template.New(path.Base(filenames[0]))
	tmpl = template.Must(tmpl.ParseFiles(filenames...))
	err := tmpl.Execute(os.Stdout, t)
	if err != nil {
		log.Println("executing template:", err)
	}
	return t
}

func (t *Transcript) GetGradeYear(year string, calendar_year string) *GradeYear {
	var retval *GradeYear
	gy, err := strconv.Atoi(year)
	if err != nil {
		return nil
	}

	retval = t.LookupGradeYear(gy)

	if retval != nil {
		return retval
	}

	// If doesn't already exist, add it
	retval = retval.New(gy, calendar_year)
	t.Append(retval)
	return retval
}

// Finds existing GradeYear
func (t *Transcript) LookupGradeYear(year int) *GradeYear {
	for _, gy := range t.GradeYears {
		if gy.Year == year {
			return gy
		}
	}
	return nil
}

func (c *Course) New() *Course {
	if c == nil {
		return new(Course)
	}
	return c
}

func (gy *GradeYear) New(year int, calendar_year string) *GradeYear {
	if gy == nil {
		return new(GradeYear).Init(year, calendar_year)
	}
	return gy
}

func (gy *GradeYear) Init(year int, calendar_year string) *GradeYear {
	gy.Courses = make([]*Course, 0)
	gy.Year = year
	gy.CalendarYear = calendar_year
	return gy
}


func (gy *GradeYear) Append(course *Course) *GradeYear {
	gy.Courses = append(gy.Courses, course)
	return gy
}

func calcGpaCredit(grade string) (int, error) {
	var gpacredit int
	var err error
    switch grade {
		case "A":
			gpacredit = 4
		case "B":
			gpacredit = 3
		case "C":
			gpacredit = 2
		case "D":
			gpacredit = 1
		case "F":
			gpacredit = 0
		default:
			err = errors.New("Not a valid grade")
	}	
	return gpacredit, err
}

func (gy *GradeYear) Recalc() *GradeYear {
	gy.Earn = 0

	var total_credits float32
	var total_courses float32
	for _, course := range gy.Courses {
		gpacredit, err := calcGpaCredit(course.Grade)

		if err != nil {
			continue
		}
		gy.Earn = gy.Earn + course.Credit
		total_credits = total_credits + float32(gpacredit)
		total_courses = total_courses + 1.0
	}

	gy.GPA = total_credits / total_courses
	gy.GpaPoints = total_credits
	return gy
}


func (t *Transcript) Recalc() *Transcript {
	for _, year := range t.GradeYears {
		t.TotalCredits = t.TotalCredits + float32(year.Earn)
		t.GpaPoints = t.GpaPoints + year.GpaPoints
	
	}
	t.Gpa = (t.GpaPoints / (t.TotalCredits * 4.0))*4.0 
	return t
}