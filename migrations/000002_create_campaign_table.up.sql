CREATE TABLE `campaign` (
    `id` INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `user_id` INT(11) NOT NULL,  
    `name` VARCHAR(255) NOT NULL,
    `short_description` VARCHAR(255) NULL,
    `description` TEXT NULL,
    `perks` TEXT NULL,
    `backer_count` INT(11) NULL,
    `goal_amount` INT(11) NULL,
    `current_amount` INT(11) NULL,
    `slug` VARCHAR(255) NULL,
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);