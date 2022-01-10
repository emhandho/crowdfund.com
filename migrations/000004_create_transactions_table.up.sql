CREATE TABLE transactions (
    `id` INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `campaign_id` INT(11) NOT NULL REFERENCES campaign (id) ON DELETE CASCADE,
    `user_id` INT(11) NOT NULL REFERENCES user (id) ON DELETE CASCADE,
    `amount` INT(11) NOT NULL,
    `status` VARCHAR(255) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
);