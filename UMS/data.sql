CREATE TABLE Department (
  ID INT PRIMARY KEY,
  Department_Name VARCHAR(40)
);

CREATE TABLE Courses (
  ID INT PRIMARY KEY,
  course_Name VARCHAR(40),
  Department_id INT,
  FOREIGN KEY (Department_id) REFERENCES Department(ID)
);

CREATE TABLE Lecturers (
  ID INT PRIMARY KEY,
  Lecturer_Name VARCHAR(40),
  Department_id INT,
  Course_id INT,
  FOREIGN KEY (Department_id) REFERENCES Department(ID),
  FOREIGN KEY (Course_id) REFERENCES Courses(ID)
);

CREATE TABLE Units (
  ID INT PRIMARY KEY,
  Unit_Name VARCHAR(40),
  Course_id INT,
  Lecturer_id INT,
  FOREIGN KEY (Course_id) REFERENCES Courses(ID),
  FOREIGN KEY (Lecturer_id) REFERENCES Lecturers(ID)
);

CREATE TABLE Students (
  ID SERIAL PRIMARY KEY,
  Student_Name VARCHAR(40),
  Course_id INT,
  Department_id INT,
  FOREIGN KEY (Course_id) REFERENCES Courses(ID),
  FOREIGN KEY (Department_id) REFERENCES Department(ID)
);

CREATE TABLE StudentUnits (
  Student_id INT,
  Unit_id INT,
  PRIMARY KEY (Student_id, Unit_id),
  FOREIGN KEY (Student_id) REFERENCES Students(ID),
  FOREIGN KEY (Unit_id) REFERENCES Units(ID)
);

--comment
-- Create a trigger function
CREATE OR REPLACE FUNCTION check_department_constraint()
RETURNS TRIGGER AS $$
DECLARE
    student_department_id INT;
    unit_department_id INT;
BEGIN
    SELECT Department_id INTO student_department_id FROM Students WHERE ID = NEW.Student_id;
    SELECT Department_id INTO unit_department_id FROM Courses WHERE ID = (
        SELECT Course_id FROM Units WHERE ID = NEW.Unit_id
    );

    IF student_department_id <> unit_department_id THEN
        RAISE EXCEPTION 'Cannot insert a unit from a different department';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Create a trigger on the StudentUnits table
CREATE TRIGGER enforce_department_constraint
BEFORE INSERT ON StudentUnits
FOR EACH ROW
EXECUTE FUNCTION check_department_constraint();




INSERT INTO Department (ID, Department_Name) VALUES
  (1, 'Computer Science'),
  (2, 'Education'),
  (3, 'Hospitality'),
  (4, 'Business');

INSERT INTO Courses (ID, course_Name, Department_id) VALUES
  (1, 'BsC', 1),
  (2, 'MsC', 1),
  (3, 'Cyber Sec', 1),
  (4, 'English lit', 2),
  (5, 'Mathematics', 2),
  (6, 'Hotel Mgt', 3),
  (7, 'Finance', 4);

INSERT INTO Lecturers (ID, Lecturer_Name, Department_id, Course_id) VALUES
  (1, 'Vitalis', 1, 1),
  (2, 'John', 1, 2),
  (3, 'Steve', 2, 4),
  (4, 'Sarah', 2, 4),
  (5, 'Leon', 2, 5),
  (6, 'Jane', 3, 6),
  (7, 'Peter', 3, 6),
  (8, 'Shon', 4, 7),
  (9, 'Milly', 4, 7),
  (10, 'Mike', 1, 3);

INSERT INTO Units (ID, Unit_Name, Course_id, Lecturer_id) VALUES
  (101, 'Software Engineering', 1, 1),
  (102, 'Computer Repair', 1, 1),
  (103, 'Computer Programming', 1, 1),
  (104, 'Cyber Security', 3, 10),
  (105, 'Ethical Hacking', 3, 10),
  (106, 'Penetrating Testing', 3, 10),
  (107, 'Distributed Systems', 2, 2),
  (108, 'Design Analysis', 2, 2),
  (109, 'Research', 4, 4),
  (110, 'Oral Literature', 4, 3),
  (111, 'Grammar', 4, 4),
  (112, 'Discrete Maths', 5, 5),
  (113, 'Calculus', 5, 5),
  (114, 'Probability & Statistics', 5, 5),
  (115, 'Pastry', 6, 6),
  (116, 'Front Office', 6, 7),
  (117, 'Hotel Mgt', 6, 6),
  (118, 'Legal Aspects', 7, 8),
  (119, 'Business Management', 7, 8),
  (120, 'Accounting', 7, 9),
  (121, 'Business Analytics', 7, 9);

INSERT INTO Students (Student_Name, Department_id, Course_id) VALUES
  ('Julius', 1, 1),
  ('Emily', 1, 1),
  ('Hellen', 1, 1),
  ('Sherry', 1, 3),
  ('Elon', 1, 3),
  ('Caren', 1, 2),
  ('Paul', 2, 4),
  ('Grace', 2, 5),
  ('Ken', 3, 6),
  ('Billy', 4, 7);

INSERT INTO StudentUnits (Student_id, Unit_id) VALUES
  -- Add other units for Julius here
  (1, 104),
  (1, 105),
  -- Add other units for Emily here
  (2, 104),
  (2, 105),
  -- Add other units for Hellen here
  (3, 106),
  
  -- For Sherry
  (4, 104),
  (4, 106),

  -- For Elon
  (5, 104),
  (5, 105),

  -- For Caren
  (6, 107),
  (6, 108),

  -- For Paul
  (7, 109),
  (7, 110),

  -- For Grace
  (8, 113),
  (8, 114),

  -- For Ken
  (9, 115),
  (9, 116),
  (9, 117),

  -- For Billy
  (10, 118),
  (10, 119),
  (10, 120);
