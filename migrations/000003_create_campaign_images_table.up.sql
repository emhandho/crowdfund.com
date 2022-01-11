CREATE TABLE `campaign_images` (
    `id` INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `campaign_id` INT(11) NOT NULL,
    `file_name` VARCHAR(255) NULL,
    `is_primary` TINYINT NOT NULL,
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE
);