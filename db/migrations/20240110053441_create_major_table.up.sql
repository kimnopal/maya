CREATE TABLE
    majors (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        name VARCHAR(100) NOT NULL,
        created_at BIGINT UNSIGNED NOT NULL,
        updated_at BIGINT UNSIGNED NOT NULL,
        faculty_id BIGINT UNSIGNED NOT NULL,
        CONSTRAINT FK_major_faculty FOREIGN KEY (faculty_id) REFERENCES faculties (id)
    ) engine = InnoDB;