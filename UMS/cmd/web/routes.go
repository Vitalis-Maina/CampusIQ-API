package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	// Home route
	router.HandleFunc("/", home).Methods("GET")

	// Student routes
	studentRoutes := router.PathPrefix("/v1/students").Subrouter()
	studentRoutes.HandleFunc("", app.listData).Methods("GET")
	studentRoutes.HandleFunc("/{id}", app.showStudent).Methods("GET")
	studentRoutes.HandleFunc("", app.insertStudent).Methods("POST")
	studentRoutes.HandleFunc("/{id}", app.updateStudent).Methods("PUT")
	studentRoutes.HandleFunc("/{id}", app.deleteStudent).Methods("DELETE")

	// Department routes
	departmentRoutes := router.PathPrefix("/v1/departments").Subrouter()
	departmentRoutes.HandleFunc("", app.listData).Methods("GET")
	departmentRoutes.HandleFunc("", app.insertDepartment).Methods("POST")
	departmentRoutes.HandleFunc("/{id}", app.updateDepartment).Methods("PUT")
	departmentRoutes.HandleFunc("/{id}", app.deleteDepartment).Methods("DELETE")

	// Course routes
	courseRoutes := router.PathPrefix("/v1/courses").Subrouter()
	courseRoutes.HandleFunc("", app.listData).Methods("GET")
	courseRoutes.HandleFunc("", app.insertCourse).Methods("POST")
	courseRoutes.HandleFunc("/{id}", app.updateCourse).Methods("PUT")
	courseRoutes.HandleFunc("/{id}", app.deleteCourse).Methods("DELETE")

	// Lecturer routes
	lecturerRoutes := router.PathPrefix("/v1/lecturers").Subrouter()
	lecturerRoutes.HandleFunc("", app.listData).Methods("GET")
	lecturerRoutes.HandleFunc("", app.insertLecturer).Methods("POST")
	lecturerRoutes.HandleFunc("/{id}", app.updateLecturer)
	lecturerRoutes.HandleFunc("/{id}", app.deleteLecturer).Methods("DELETE")

	// Unit routes
	unitRoutes := router.PathPrefix("/v1/units").Subrouter()
	unitRoutes.HandleFunc("", app.listData).Methods("GET")
	unitRoutes.HandleFunc("", app.insertUnit).Methods("POST")
	unitRoutes.HandleFunc("/{id}", app.updateUnit)
	unitRoutes.HandleFunc("/{id}", app.deleteUnit).Methods("DELETE")

	// StudentUnit routes
	studentUnitRoutes := router.PathPrefix("/v1/studentunits").Subrouter()
	studentUnitRoutes.HandleFunc("", app.listData).Methods("GET")
	studentUnitRoutes.HandleFunc("", app.insertStudentUnit).Methods("POST")
	studentUnitRoutes.HandleFunc("/{studentid}/{unitid}", app.updateStudentUnit).Methods("PUT")
	studentUnitRoutes.HandleFunc("/{studentid}/{unitid}", app.deleteStudentUnit).Methods("DELETE")

	return router
}
