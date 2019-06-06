package api

import (
	"app-api/config"
	"app-api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func AllStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		studentModel := models.StudentModel{
			Db: db,
		}

		students, err2 := studentModel.FindAll()

		if err2 != nil {
			json.NewEncoder(w).Encode(err2)
		} else {
			json.NewEncoder(w).Encode(students)
		}

	}

}
