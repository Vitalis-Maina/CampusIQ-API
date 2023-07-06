# University Management System (UMS)

University Management System (UMS) is a web application that performs CRUD (Create, Read, Update, Delete) operations for managing university-related data. It allows users to manage students, departments, courses, lecturers, units, and student-unit relationships.

## GitHub Repository

The source code for the UMS project can be found on GitHub: [UMS GitHub Repository](https://github.com/Vitalis-Maina/Projects)

## Cloning the Repository

To clone the UMS project repository using SSH, follow these steps:

1. Open your terminal or command-line interface.

2. Navigate to the directory where you want to clone the repository.

3. Copy the SSH URL of the repository from the GitHub page:

git@github.com:Vitalis-Maina/Projects.git

4. Run the following command to clone the repository:

git clone git@github.com:Vitalis-Maina/Projects.git


5. Provide your SSH passphrase, if prompted.

6. The repository will be cloned to your local machine, including the UMS folder.

## Technologies Used

The University Management System is built using the following technologies:

- Go: Programming language used for backend development.
- Gorilla Mux: A powerful HTTP router and URL matcher for building Go web applications. It is used for handling routing in the UMS project.
- PostgreSQL: A popular open-source relational database management system used for storing and managing data in the UMS project.

## Usage

The UMS provides different endpoints to interact with the system. Below are the available endpoints and their corresponding operations:

### Students

- `GET /v1/students`: Retrieves a list of all students.
- `GET /v1/students/{id}`: Retrieves details of a specific student.
- `POST /v1/students`: Creates a new student.
- `PUT /v1/students/{id}`: Updates information for a specific student.
- `DELETE /v1/students/{id}`: Deletes a specific student.

### Departments

- `GET /v1/departments`: Retrieves a list of all departments.
- `POST /v1/departments`: Creates a new department.
- `PUT /v1/departments/{id}`: Updates information for a specific department.
- `DELETE /v1/departments/{id}`: Deletes a specific department.

### Courses

- `GET /v1/courses`: Retrieves a list of all courses.
- `POST /v1/courses`: Creates a new course.
- `PUT /v1/courses/{id}`: Updates information for a specific course.
- `DELETE /v1/courses/{id}`: Deletes a specific course.

### Lecturers

- `GET /v1/lecturers`: Retrieves a list of all lecturers.
- `POST /v1/lecturers`: Creates a new lecturer.
- `PUT /v1/lecturers/{id}`: Updates information for a specific lecturer.
- `DELETE /v1/lecturers/{id}`: Deletes a specific lecturer.

### Units

- `GET /v1/units`: Retrieves a list of all units.
- `POST /v1/units`: Creates a new unit.
- `PUT /v1/units/{id}`: Updates information for a specific unit.
- `DELETE /v1/units/{id}`: Deletes a specific unit.

### Student-Unit Relationships

- `GET /v1/studentunits`: Retrieves a list of all student-unit relationships.
- `POST /v1/studentunits`: Creates a new student-unit relationship.
- `PUT /v1/studentunits/{studentid}/{unitid}`: Updates the relationship between a specific student and unit.
- `DELETE /v1/studentunits/{studentid}/{unitid}`: Deletes the relationship between a specific student and unit.

## Contributions

Contributions to the University Management System project are welcome! If you encounter any issues or have suggestions for improvements, please feel free to submit an issue or a pull request on the [UMS GitHub Repository](https://github.com/Vitalis-Maina/Projects).

## Acknowledgements

- This application was created to demonstrate CRUD functionality for managing university-related data.
- The design and implementation of the application are credited to Vitalis Maina.

