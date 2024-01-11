CREATE TABLE
    post_tag (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        post_id BIGINT UNSIGNED NOT NULL,
        tag_id BIGINT UNSIGNED NOT NULL,
        CONSTRAINT FK_post_tag_post FOREIGN KEY (post_id) REFERENCES posts (id),
        CONSTRAINT FK_post_tag_tag FOREIGN KEY (tag_id) REFERENCES tags (id)
    ) engine = InnoDB;