CREATE TABLE `user` (
    `id` INT(11) PRIMARY KEY NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `occupation` VARCHAR(255) NULL,
    `hash_password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `avatar_name` VARCHAR(255) NULL,
    `role` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
);