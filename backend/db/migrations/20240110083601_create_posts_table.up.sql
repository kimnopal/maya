CREATE TABLE
    posts (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        title VARCHAR(255) NOT NULL,
        code VARCHAR(100) UNIQUE NOT NULL,
        description LONGTEXT NOT NULL,
        created_at BIGINT UNSIGNED NOT NULL,
        updated_at BIGINT UNSIGNED NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        post_category_id BIGINT UNSIGNED NOT NULL,
        CONSTRAINT FK_post_user FOREIGN KEY (user_id) REFERENCES users (id),
        CONSTRAINT FK_post_post_category FOREIGN KEY (post_category_id) REFERENCES post_categories (id)
    ) engine = InnoDB;