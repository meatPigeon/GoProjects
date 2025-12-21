CREATE TABLE groups (
                        group_id   SERIAL PRIMARY KEY,
                        group_name VARCHAR(50) NOT NULL
);

CREATE TABLE students (
                          student_id   SERIAL PRIMARY KEY,
                          student_name VARCHAR(50) NOT NULL,
                          birth_date   DATE NOT NULL,
                          gender       VARCHAR(6),
                          created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          group_id     INT REFERENCES groups(group_id)
);

ALTER TABLE students
DROP COLUMN created_at

CREATE TABLE class_schedule (
                                class_id    SERIAL PRIMARY KEY,
                                group_id    INT REFERENCES groups(group_id),
                                class_name  VARCHAR(100) NOT NULL,
                                class_date  DATE NOT NULL,
                                start_time  TIME NOT NULL,
                                end_time    TIME NOT NULL,
                                room        VARCHAR(20),
                                teacher_id  INT
);

SELECT * FROM students
WHERE gender = 'Female'
ORDER BY birth_date DESC