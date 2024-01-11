CREATE TABLE
    faculties (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        name VARCHAR(100) NOT NULL,
        created_at BIGINT UNSIGNED NOT NULL,
        updated_at BIGINT UNSIGNED NOT NULL
    ) engine = InnoDB;