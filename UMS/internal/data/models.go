package data

import (
	"context"
	"database/sql"
	"log"
	"time"
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
	ID           int64  `jspn:"course_id"`
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
type UmsModel struct {
	DB *sql.DB
}

func NewModels(db *sql.DB) UmsModel {
	return UmsModel{
		DB: db,
	}
}

func (u UmsModel) InsertStudent(student Student) error {
	query := `
		INSERT INTO Students(student_name, course_id, department_id)
		VALUES ($1, $2, $3)
	`

	_, err := u.DB.Exec(query, student.Name, student.CourseID, student.DepartmentID)
	if err != nil {
		log.Println("Failed to insert student:", err)
		return err
	}

	return nil
}

func (u UmsModel) InsertDepartment(department Department) error {
	query := `
		INSERT INTO Department(id, department_name)
		VALUES($1, $2)
	`

	_, err := u.DB.Exec(query, department.ID, department.Name)
	if err != nil {
		log.Println("Failed to insert department:", err)
		return err
	}

	return nil
}

func (u UmsModel) InsertCourse(course Courses) error {
	query := `
		INSERT INTO Courses(id, course_name, department_id)
		VALUES($1, $2, $3)
	`

	_, err := u.DB.Exec(query, course.ID, course.Name, course.DepartmentID)
	if err != nil {
		log.Println("Failed to insert course:", err)
		return err
	}

	return nil
}

func (u UmsModel) InsertLecturer(lecturer Lecturers) error {
	query := `
		INSERT INTO Lecturers(id, lecturer_name, department_id, course_id)
		VALUES($1, $2, $3, $4)

	`

	_, err := u.DB.Exec(query, lecturer.ID, lecturer.Name, lecturer.DepartmentID, lecturer.CourseID)
	if err != nil {
		log.Println("Failed to insert lecturer:", err)
		return err
	}

	return nil
}

func (u UmsModel) InsertUnit(unit Units) error {
	query := `
		INSERT INTO Units(id, unit_name, course_id, lecturer_id)
		VALUES($1, $2, $3, $4)
	`

	_, err := u.DB.Exec(query, unit.ID, unit.Name, unit.CourseID, unit.LecturerID)
	if err != nil {
		log.Println("Failed to insert unit:", err)
		return err
	}

	return nil
}

func (u UmsModel) InsertStudentUnit(studentUnit StudentUnits) error {
	query := `
		INSERT INTO StudentUnits(student_id, unit_id)
		VALUES($1, $2)
	`

	_, err := u.DB.Exec(query, studentUnit.StudentID, studentUnit.UnitID)
	if err != nil {
		log.Println("Failed to insert student unit:", err)
		return err
	}

	return nil
}

func (u UmsModel) GetStudents() ([]Student, error) {
	query := `SELECT student_name, course_id, department_id FROM Students`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []Student{}

	for rows.Next() {
		var s Student

		err := rows.Scan(&s.Name, &s.CourseID, &s.DepartmentID)
		if err != nil {
			return nil, err
		}

		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (u UmsModel) GetDepartments() ([]Department, error) {
	query := `SELECT id, department_name FROM Department`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	departments := []Department{}

	for rows.Next() {
		var d Department

		err := rows.Scan(&d.ID, &d.Name)
		if err != nil {
			return nil, err
		}

		departments = append(departments, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

// Implement the other methods for Courses, Lecturers, Units, and StudentUnits similarly.
func (u UmsModel) GetCourses() ([]Courses, error) {
	query := `SELECT id, course_name, department_id FROM Courses`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Courses{}

	for rows.Next() {
		var c Courses

		err := rows.Scan(&c.ID, &c.Name, &c.DepartmentID)
		if err != nil {
			return nil, err
		}

		courses = append(courses, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (u UmsModel) GetLecturers() ([]Lecturers, error) {
	query := `SELECT id, lecturer_name, department_id, course_id FROM Lecturers`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lecturers := []Lecturers{}

	for rows.Next() {
		var l Lecturers

		err := rows.Scan(&l.ID, &l.Name, &l.DepartmentID, &l.CourseID)
		if err != nil {
			return nil, err
		}

		lecturers = append(lecturers, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lecturers, nil
}

func (u UmsModel) GetUnits() ([]Units, error) {
	query := `SELECT id, unit_name, course_id, lecturer_id FROM Units`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	units := []Units{}

	for rows.Next() {
		var u Units

		err := rows.Scan(&u.ID, &u.Name, &u.CourseID, &u.LecturerID)
		if err != nil {
			return nil, err
		}

		units = append(units, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return units, nil
}

func (u UmsModel) GetStudentUnits() ([]StudentUnits, error) {
	query := `SELECT student_id, unit_id FROM StudentUnits`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	studentUnits := []StudentUnits{}

	for rows.Next() {
		var su StudentUnits

		err := rows.Scan(&su.StudentID, &su.UnitID)
		if err != nil {
			return nil, err
		}

		studentUnits = append(studentUnits, su)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return studentUnits, nil
}

func (u UmsModel) DeleteStudent(studentID int64) error {

	query := `
		DELETE FROM Students WHERE ID= $1
		RETURNING id
	`
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancle()

	_, err := u.DB.ExecContext(ctx, query, studentID)
	if err != nil {
		log.Fatal("failed to delete student", err)

	}
	return nil

}

func (u UmsModel) DeleteDepartment(DepartmentID int) error {
	query := `
		DELETE FROM Department WHERE ID=$1
	`

	_, err := u.DB.Exec(query, DepartmentID)
	if err != nil {
		log.Fatal("failed to delete department", err)
	}
	return nil
}

func (u UmsModel) DeleteCourse(courseid int) error {
	query := `
		DELETE FROM Courses WHERE ID=$1
	`
	_, err := u.DB.Exec(query, courseid)
	if err != nil {
		log.Fatal("failed to delete course", err)
	}
	return nil
}
func (u UmsModel) DeleteLecturer(lecturerID int) error {
	query := `
	
		DELETE FROM Lecturers WHERE ID=$1
	`
	_, err := u.DB.Exec(query, lecturerID)
	if err != nil {
		log.Fatal("failed to delete lecturer", err)
	}
	return nil
}

func (u UmsModel) DeleteUnit(unitID int) error {
	query := `
		DELETE FROM Units WHERE ID=$1
	`
	_, err := u.DB.Exec(query, unitID)
	if err != nil {
		log.Fatal("failed to delete unit", err)
	}
	return nil
}

func (u UmsModel) DeleteStudentUnit(studentID, unitID int) error {
	query := `
	
		DELETE FROM StudentUnits WHERE student_id = $1 AND unit_id = $2
	`
	_, err := u.DB.Exec(query, studentID, unitID)
	if err != nil {
		log.Fatal("failed to delete student unit", err)
	}
	return nil
}

func (u UmsModel) UpdateStudentUnits(studentid int, unitid int, s StudentUnits) error {

	query := `
		UPDATE StudentUnits SET student_id=$1,unit_id=$2
		 WHERE student_id=$3 AND unit_id=$4
	`
	_, err := u.DB.Exec(query, s.StudentID, s.UnitID, studentid, unitid)
	if err != nil {
		log.Fatal("failed to update studentunits")
	}
	return nil
}

func (u UmsModel) UpdateStudent(studentid int, s Student) error {

	query := `
		UPDATE Students SET student_name=$1,course_id=$2,department_id=$3 WHERE id=$4
	`
	_, err := u.DB.Exec(query, s.Name, s.CourseID, s.DepartmentID, studentid)

	if err != nil {
		log.Fatal("failed to update student", err)
	}
	return nil
}

func (u UmsModel) UpdateUnit(unitid int, s Units) error {

	query := `
	
	UPDATE Units SET id=$1,unit_name=$2,course_id=$3,lecturer_id=$4 WHERE
	id=$5
	`
	_, err := u.DB.Exec(query, s.ID, s.Name, s.CourseID, s.LecturerID, unitid)
	if err != nil {
		log.Fatal("failed to update unit ", err)
	}
	return nil
}

func (u UmsModel) UpdateLecturer(lecturerid int, l Lecturers) error {
	query := `
	
	UPDATE Lecturers SET id=$1,lecturer_name=$2,department_id=$3,course_id=$4
	WHERE id=$5
	`
	_, err := u.DB.Exec(query, l.ID, l.Name, l.DepartmentID, l.CourseID, lecturerid)
	if err != nil {
		log.Fatal("failed to update lecturer ", err)
	}
	return nil
}

func (u UmsModel) UpdateDepartment(departmentid int, d Department) error {
	query := `
		UPDATE Department SET id=$1,department_name=$2
		WHERE id=$3
	`
	_, err := u.DB.Exec(query, d.ID, d.Name, departmentid)
	if err != nil {
		log.Fatal("failed to update department ", err)
	}
	return nil
}

func (u UmsModel) UpdateCourse(courseid int, c Courses) error {
	query := `

	UPDATE Courses SET id=$1,course_name=$2,department_id=$3
	WHERE id=$4
	`
	_, err := u.DB.Exec(query, c.ID, c.Name, c.DepartmentID, courseid)
	if err != nil {
		log.Fatal("failed to update course ", err)
	}
	return nil
}
