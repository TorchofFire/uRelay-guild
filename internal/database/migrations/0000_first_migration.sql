CREATE TABLE `migrations` (
  `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO `migrations` (`id`) VALUES (0000);

CREATE TABLE `users` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `public_key` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `join_date` INT UNSIGNED NOT NULL,
  INDEX (`public_key`)
);

DELIMITER //
CREATE TRIGGER before_insert_users
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
  SET NEW.join_date = UNIX_TIMESTAMP();
END//
DELIMITER ;

CREATE TABLE `guild_categories` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `display_priority` SMALLINT UNSIGNED NOT NULL DEFAULT 0
);

CREATE TABLE `guild_channels` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `channel_type` VARCHAR(255) NOT NULL,
  `display_priority` SMALLINT UNSIGNED NOT NULL DEFAULT 0,
  `category_id` BIGINT UNSIGNED,
  FOREIGN KEY (`category_id`) REFERENCES `guild_categories`(`id`) ON DELETE CASCADE, 
  INDEX (`category_id`)
);

CREATE TABLE `guild_messages` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `sender_id` BIGINT UNSIGNED NOT NULL,
  `message` TEXT NOT NULL,
  `channel_id` BIGINT UNSIGNED NOT NULL,
  `sent_at` INT UNSIGNED NOT NULL,
  FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`channel_id`) REFERENCES `guild_channels`(`id`) ON DELETE CASCADE,
  INDEX (`sender_id`),
  INDEX (`channel_id`)
);

DELIMITER //
CREATE TRIGGER before_insert_guild_messages
BEFORE INSERT ON guild_messages
FOR EACH ROW
BEGIN
  SET NEW.sent_at = UNIX_TIMESTAMP();
END//
DELIMITER ;

-- REVERT MIGRATION

DROP TRIGGER IF EXISTS before_insert_guild_messages;
DROP TABLE `guild_messages`;
DROP TABLE `guild_channels`;
DROP TABLE `guild_categories`;
DROP TRIGGER IF EXISTS before_insert_users;
DROP TABLE `users`;
DROP TABLE `migrations`;