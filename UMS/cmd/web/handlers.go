package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vitalis-Maina/internal/data"
	"github.com/gorilla/mux"
)

type Student struct {
	Name         string `json:"student_name"`
	CourseID     int64  `json:"course_id"`
	DepartmentID int64  `json:"department_id"`
}

type Department struct {
	ID   int64  `json:"department_id"`
	Name string `json:"department_name"`
}

type Courses struct {
	ID           int64  `json:"course_id"`
	Name         string `json:"course_name"`
	DepartmentID int64  `json:"department_id"`
}
type Lecturers struct {
	ID           int64  `json:"lecturer_id"`
	Name         string `json:"lecturer_name"`
	DepartmentID int64  `json:"department_id"`
	CourseID     int64  `json:"course_id"`
}

type Units struct {
	ID         int64  `json:"unit_id"`
	Name       string `json:"unit_name"`
	CourseID   int64  `json:"course_id"`
	LecturerID int64  `json:"lecturer_id"`
}
type StudentUnits struct {
	StudentID int64 `json:"student_id"`
	UnitID    int64 `json:"unit_id"`
}

func (app *application) showStudent(w http.ResponseWriter, r *http.Request) {

	//param := flow.Param(r.Context(), "id")
	param := mux.Vars(r)
	studentId := param["id"]
	fmt.Fprintf(w, "displaying student record with id %s", studentId)

}
func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	studentid, err := strconv.Atoi(id["id"])
	if err != nil {
		log.Fatal(err)
	}
	app.models.DeleteStudent(int64(studentid))
	fmt.Fprintf(w, "student with id :%d has been successfully deleted", studentid)
}
func (app *application) deleteDepartment(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	departmentid, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = app.models.DeleteDepartment(departmentid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "department with id: %d has been successfully deleted", departmentid)
}

func (app *application) deleteCourse(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	courseid, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = app.models.DeleteCourse(courseid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "course with id: %d has been successffully deleted", courseid)
}
func (app *application) deleteLecturer(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	lecturerid, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = app.models.DeleteLecturer(lecturerid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "lecturer with id:%d has been successfully deleted", lecturerid)
}
func (app *application) deleteUnit(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	unitid, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = app.models.DeleteUnit(unitid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "unit with id:%d has been successfully deleted ", unitid)
}
func (app *application) deleteStudentUnit(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	studentid, err := strconv.Atoi(param["studentid"])
	if err != nil {
		log.Fatal(err)
	}
	unitid, err := strconv.Atoi(param["unitid"])
	if err != nil {
		log.Fatal(err)
	}
	err = app.models.DeleteStudentUnit(studentid, unitid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Student unit %d - %d has been successfully deleted", studentid, unitid)
}
func (app *application) listData(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error

	switch r.URL.Path {
	case "/v1/students":
		data, err = app.models.GetStudents()
	case "/v1/departments":
		data, err = app.models.GetDepartments()
	case "/v1/courses":
		data, err = app.models.GetCourses()
	case "/v1/lecturers":
		data, err = app.models.GetLecturers()
	case "/v1/units":
		data, err = app.models.GetUnits()
	case "/v1/studentunits":
		data, err = app.models.GetStudentUnits()
	default:
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	j := json.NewEncoder(w)
	j.SetIndent("", " ")
	err = j.Encode(data)

	if err != nil {
		log.Fatal(err)
	}
}
func (app *application) insertStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = app.models.InsertStudent(data.Student(student))
	if err != nil {
		http.Error(w, "Failed to insert student", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Student inserted successfully")
}

func (app *application) insertDepartment(w http.ResponseWriter, r *http.Request) {
	var department Department
	err := json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = app.models.InsertDepartment(data.Department(department))
	if err != nil {
		http.Error(w, "Failed to insert department", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Department inserted successfully")
}

func (app *application) insertCourse(w http.ResponseWriter, r *http.Request) {
	var course Courses
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = app.models.InsertCourse(data.Courses(course))
	if err != nil {
		http.Error(w, "Failed to insert course", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Course inserted successfully")
}

func (app *application) insertLecturer(w http.ResponseWriter, r *http.Request) {
	var lecturer Lecturers
	err := json.NewDecoder(r.Body).Decode(&lecturer)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = app.models.InsertLecturer(data.Lecturers(lecturer))
	if err != nil {
		http.Error(w, "Failed to insert lecturer", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Lecturer inserted successfully")
}

func (app *application) insertUnit(w http.ResponseWriter, r *http.Request) {
	var unit Units
	err := json.NewDecoder(r.Body).Decode(&unit)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = app.models.InsertUnit(data.Units(unit))
	if err != nil {
		http.Error(w, "Failed to insert unit", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Unit inserted successfully")
}

func (app *application) insertStudentUnit(w http.ResponseWriter, r *http.Request) {
	var studentUnit StudentUnits
	err := json.NewDecoder(r.Body).Decode(&studentUnit)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = app.models.InsertStudentUnit(data.StudentUnits(studentUnit))
	if err != nil {
		http.Error(w, "Failed to insert student unit", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "StudentUnit inserted successfully")
}

func (app *application) updateStudentUnit(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	studentid, err := strconv.Atoi(params["studentid"])
	if err != nil {
		log.Fatal(err)
	}
	unitid, err := strconv.Atoi(params["unitid"])
	if err != nil {
		log.Fatal(err)
	}
	var studentUnit StudentUnits
	err = json.NewDecoder(r.Body).Decode(&studentUnit)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}
	err = app.models.UpdateStudentUnits(studentid, unitid, data.StudentUnits(studentUnit))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "studentUnit  %d - %d updated successfully", studentid, unitid)
}

func (app *application) updateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studentid, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var student Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
	}
	err = app.models.UpdateStudent(studentid, data.Student(student))
	if err != nil {
		http.Error(w, "failed to update student", http.StatusInternalServerError)

	}
	fmt.Fprintf(w, "student with id %d updated successfully", studentid)
}

func (app *application) updateUnit(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	unitid, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)

	}
	var unit Units
	err = json.NewDecoder(r.Body).Decode(&unit)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
	}
	err = app.models.UpdateUnit(unitid, data.Units(unit))
	if err != nil {
		http.Error(w, "failed to update unit", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Unit with id %d updated successfully", unitid)
}

func (app *application) updateLecturer(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	lecturerid, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	var lecturer Lecturers
	err = json.NewDecoder(r.Body).Decode(&lecturer)
	if err != nil {
		http.Error(w, "failed to decode request body ", http.StatusBadRequest)
	}
	err = app.models.UpdateLecturer(lecturerid, data.Lecturers(lecturer))
	if err != nil {
		http.Error(w, "failed to update lecturer", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "lecturer with id %d successfully updated", lecturerid)

}

func (app *application) updateDepartment(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	departmentid, err := strconv.Atoi(param["id"])
	if err != nil {
		http.Error(w, "failed to parse request route variable", http.StatusNotAcceptable)
		return
	}
	var department Department
	err = json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)

	}
	err = app.models.UpdateDepartment(departmentid, data.Department(department))
	if err != nil {
		http.Error(w, "failed to update department", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Department with id %d updated successfully ", departmentid)
}

func (app *application) updateCourse(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	courseid, err := strconv.Atoi(param["id"])
	if err != nil {
		http.Error(w, "failed parsing request route variable", http.StatusNotAcceptable)
		return
	}
	var course Courses
	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
	}

	err = app.models.UpdateCourse(courseid, data.Courses(course))
	if err != nil {
		http.Error(w, "failed to update course ", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "course with id %d updated successfully", courseid)
}
