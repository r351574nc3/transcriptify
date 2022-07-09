package transcript

type Transcript struct {
	GradeYears []*GradeYear
}

type GradeYear struct {
	Courses []*Course
	Year int
}

type Course struct {
	Name string
	Grade string
	Credit int
}