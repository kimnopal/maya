CREATE TABLE
    users (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(100) UNIQUE NOT NULL,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at BIGINT UNSIGNED NOT NULL,
        updated_at BIGINT UNSIGNED NOT NULL,
        role_id BIGINT UNSIGNED NOT NULL,
        major_id BIGINT UNSIGNED NOT NULL,
        CONSTRAINT FK_user_role FOREIGN KEY (role_id) REFERENCES roles (id),
        CONSTRAINT FK_user_major FOREIGN KEY (major_id) REFERENCES majors (id)
    ) engine = InnoDB;