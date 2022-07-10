package transcript

type Transcript struct {
	GradeYears []*GradeYear
	TotalCredits  float32
	GpaPoints float32
	Gpa float32
}

type GradeYear struct {
	Courses []*Course
	Year int
	CalendarYear string
	GPA float32
	Earn int
	GpaPoints float32
	
}

type Course struct {
	Name string
	Grade string
	Credit int
}