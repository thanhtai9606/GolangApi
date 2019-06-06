package entities

import (
	"fmt"
)

type Student struct {
	ID             int    `json:"id"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	EnrollmentDate string `json:enrollmentDate`
}

type Course struct {
	CourseID int    `json:"courseID"`
	Title    string `json:"title"`
	Credits  int    `json:"credits"`
}

type Enrollment struct {
	EnrollmentID int    `json:enrollmentID`
	CourseID     int    `json:"courseID"`
	StudentID    int    `json:"studentID"`
	Grade        string `json:"grade"`
}

func (student Student) ToString() string {
	return fmt.Sprintf("id: %d\nfirstname: %s\nlastname: %s\nenrollmentDate: %s ", student.ID, student.FirstName, student.LastName, student.EnrollmentDate)
}

func (course Course) ToString() string {
	return fmt.Sprintf("id: %d\ntitle: %s\ncredits: %d ", course.CourseID, course.Title, course.Credits)
}
