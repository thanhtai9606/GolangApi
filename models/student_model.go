package models

import (
	entities "app-api/entities"
	"database/sql"
)

type StudentModel struct {
	Db *sql.DB
}

func (studentModel StudentModel) FindAll() (student []entities.Student, err error) {
	rows, err := studentModel.Db.Query("select * from Student")

	if err != nil {
		return nil, err
	} else {

		var students []entities.Student

		for rows.Next() {
			var id int
			var firstname string
			var lastname string
			var enrollmentDate string
			err2 := rows.Scan(&id, &firstname, &lastname, &enrollmentDate)

			if err2 != nil {
				return nil, err
			} else {
				student := entities.Student{
					ID:             id,
					FirstName:      firstname,
					LastName:       lastname,
					EnrollmentDate: enrollmentDate,
				}

				students = append(students, student)
			}
		}

		return students, nil
	}
}
