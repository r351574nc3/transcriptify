package transcript

import (
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

func (t *Transcript) GetGradeYear(year string) *GradeYear {
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
	retval = retval.New(gy)
	t.Append(retval)
	return retval
}

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

func (gy *GradeYear) New(year int) *GradeYear {
	if gy == nil {
		return new(GradeYear).Init(year)
	}
	return gy
}

func (gy *GradeYear) Init(year int) *GradeYear {
	gy.Courses = make([]*Course, 0)
	gy.Year = year
	return gy
}


func (gy *GradeYear) Append(course *Course) *GradeYear {
	gy.Courses = append(gy.Courses, course)
	return gy
}
