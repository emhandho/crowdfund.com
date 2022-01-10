CREATE TABLE `transactions` (
    `id` INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `campaign_id` INT(11) NOT NULL,
    `user_id` INT(11) NOT NULL,
    `amount` INT(11) NOT NULL,
    `status` VARCHAR(255) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`campaign_id`) REFERENCES `campaign` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);