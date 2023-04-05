-- Create "cards" table
CREATE TABLE `cards` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `password` text NOT NULL, `card` integer NULL, `status` text NULL, `cp` integer NULL, `url` text NULL DEFAULT 'https://card.syui.ai', `created_at` datetime NULL, `user_card` integer NOT NULL, CONSTRAINT `cards_users_card` FOREIGN KEY (`user_card`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create "groups" table
CREATE TABLE `groups` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL, `password` text NOT NULL);
-- Create index "group_name" to table: "groups"
CREATE UNIQUE INDEX `group_name` ON `groups` (`name`);
-- Create "users" table
CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `username` text NOT NULL, `password` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `next` text NULL DEFAULT '20230405', `group_users` integer NULL, CONSTRAINT `users_groups_users` FOREIGN KEY (`group_users`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX `users_username_key` ON `users` (`username`);
-- Create index "user_username" to table: "users"
CREATE UNIQUE INDEX `user_username` ON `users` (`username`);
