CREATE TABLE campaign_images (
    `id` INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `campaign_id` INT(11) NOT NULL REFERENCES campaign (id) ON DELETE CASCADE,
    `file_name` VARCHAR(255) NULL,
    `is_primary` TINYINT NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
);