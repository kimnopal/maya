CREATE TABLE
    appliers (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        created_at BIGINT UNSIGNED NOT NULL,
        updated_at BIGINT UNSIGNED NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        post_id BIGINT UNSIGNED NOT NULL,
        CONSTRAINT FK_applier_user FOREIGN KEY (user_id) REFERENCES users (id),
        CONSTRAINT FK_applier_post FOREIGN KEY (post_id) REFERENCES posts (id)
    ) engine = InnoDB;