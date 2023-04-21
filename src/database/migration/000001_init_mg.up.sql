CREATE TABLE Students (
    id CHAR(36) NOT NULL PRIMARY KEY DEFAULT UUID() DEFAULT UUID(),
    fullname varchar(255) NOT NULL,
    gender ENUM('0', '1'),
    age INT,
    regist_date DATE
);
CREATE TABLE Majors (
    id CHAR(36) NOT NULL PRIMARY KEY DEFAULT UUID(),
    major VARCHAR(255)
);
CREATE TABLE Hobbies (
    id CHAR(36) NOT NULL PRIMARY KEY DEFAULT UUID(),
    hobby VARCHAR(255)
);
CREATE TABLE Student_Hobby (
    id CHAR(36) NOT NULL PRIMARY KEY DEFAULT UUID(),
    hobby_id CHAR(36) REFERENCES Hobbies(id),
    student_id CHAR(36) REFERENCES Students(id)
);
CREATE TABLE Student_Major (
    id CHAR(36) NOT NULL PRIMARY KEY DEFAULT UUID(),
    student_id CHAR(36) REFERENCES Students(id),
    major_id CHAR(36) REFERENCES Majors(id)
);
INSERT INTO Majors (major) VALUES ('Computer Science'), ('Psychology'), ('Economics'), ('Chemical Engineering'), ('Marine Science'), ('Communication'), ('Statistic');
INSERT INTO Hobbies (hobby) VALUES ('Cooking'), ('Coding'), ('Traveling'), ('Writing'), ('Diving'), ('Swimming'), ('Dancing'), ('Cycling'), ('Photography'), ('Painting'), ('Chess'), ('Playing a musical instrument'), ('Camping'), ('Archery');