CREATE TABLE
    contacts (
        id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
        value VARCHAR(100) NOT NULL,
        created_at BIGINT UNSIGNED NOT NULL,
        updated_at BIGINT UNSIGNED NOT NULL,
        contact_type_id BIGINT UNSIGNED NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        CONSTRAINT FK_contact_contact_type FOREIGN KEY (contact_type_id) REFERENCES contacts (id),
        CONSTRAINT FK_contact_user FOREIGN KEY (user_id) REFERENCES users (id)
    ) engine = InnoDB;