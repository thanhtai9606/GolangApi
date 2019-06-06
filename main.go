package main

import (
	student_api "app-api/api"
	"app-api/config"
	models "app-api/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	startServer()

}

func startServer() {
	r := mux.NewRouter()                                                 //All
	r.HandleFunc("/api/students", student_api.AllStudent).Methods("GET") //All
	fmt.Println("Server is running in http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
func Demo() {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		courseModel := models.CourseModel{
			Db: db,
		}

		courses, err2 := courseModel.FindAll()

		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, course := range courses {
				fmt.Println(course.ToString())
				fmt.Println("=========================")
			}
		}
	}
}

func Demo2() {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		studentModel := models.StudentModel{
			Db: db,
		}

		students, err2 := studentModel.FindAll()

		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, student := range students {
				fmt.Println(student.ToString())
				fmt.Println("=========================")
			}
		}
	}
}
