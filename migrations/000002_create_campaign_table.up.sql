CREATE TABLE `campaign` (
    `id` INT(11) PRIMARY KEY NOT NULL,
    `user_id` INT(11) NOT NULL REFERENCES user (id) ON DELETE CASCADE,  
    `name` VARCHAR(255) NOT NULL,
    `short_description` VARCHAR(255) NULL,
    `description` TEXT NULL,
    `perks` TEXT NULL,
    `backer_count` INT(11) NULL,
    `goal_amount` INT(11) NULL,
    `current_amount` INT(11) NULL,
    `slug` VARCHAR(255) NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
);