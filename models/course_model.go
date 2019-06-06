package models

import (
	entities "app-api/entities"
	"database/sql"
)

type CourseModel struct {
	Db *sql.DB
}

func (courseModel CourseModel) FindAll() (course []entities.Course, err error) {
	rows, err := courseModel.Db.Query("select * from course")

	if err != nil {
		return nil, err
	} else {

		var courses []entities.Course

		for rows.Next() {
			var id int
			var title string
			var credits int
			err2 := rows.Scan(&id, &title, &credits)

			if err2 != nil {
				return nil, err
			} else {
				course := entities.Course{
					CourseID: id,
					Title:    title,
					Credits:  credits,
				}

				courses = append(courses, course)
			}
		}

		return courses, nil
	}
}
