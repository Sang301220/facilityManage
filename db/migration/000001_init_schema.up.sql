-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL
);

CREATE TABLE `statusDevices` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL
);

CREATE TABLE `devices` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `serial` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `category_id` int NOT NULL,
  `device_status_id` int NOT NULL,
  `created_at` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp
);

CREATE TABLE `requestStatus` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL
);

CREATE TABLE `roles` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `phone` varchar(100) NOT NULL,
  `role_id` int NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp
);

CREATE TABLE `requests` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `request_status_id` int NOT NULL,
  `user_id` int NOT NULL,
  `borrow_reason` text NOT NULL,
  `device_id` int NOT NULL,
  `device_receive_desire_date` timestamp NOT NULL,
  `device_receive_return_date` timestamp NOT NULL,
  `created_at` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp
);

CREATE TABLE `notices` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `status` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  `due_date` timestamp NOT NULL
);

CREATE TABLE `sessions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `client_ip` varchar(255) NOT NULL,
  `is_blocked` boolean NOT NULL DEFAULT false,
  `created_at` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `expires_at` timestamp NOT NULL
);

CREATE TABLE `noticeStatus` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL
);

CREATE INDEX `categories_index_0` ON `categories` (`name`);

CREATE INDEX `statusDevices_index_1` ON `statusDevices` (`name`);

CREATE INDEX `devices_index_2` ON `devices` (`name`);

CREATE INDEX `requestStatus_index_3` ON `requestStatus` (`name`);

CREATE INDEX `roles_index_4` ON `roles` (`name`);

CREATE INDEX `users_index_5` ON `users` (`name`);

CREATE INDEX `requests_index_6` ON `requests` (`user_id`);

CREATE INDEX `notices_index_7` ON `notices` (`status`);

CREATE INDEX `noticeStatus_index_8` ON `noticeStatus` (`name`);

ALTER TABLE `devices` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

ALTER TABLE `devices` ADD FOREIGN KEY (`device_status_id`) REFERENCES `statusDevices` (`id`);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `requests` ADD FOREIGN KEY (`request_status_id`) REFERENCES `requestStatus` (`id`);

ALTER TABLE `requests` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `requests` ADD FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`);

ALTER TABLE `notices` ADD FOREIGN KEY (`status`) REFERENCES `noticeStatus` (`id`);

ALTER TABLE `notices` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `sessions` ADD FOREIGN KEY (`name`) REFERENCES `users` (`name`);

-- SEED

INSERT INTO roles (name) VALUES ('admin');
INSERT INTO roles (name) VALUES ('it');
INSERT INTO roles (name) VALUES ('user');

INSERT INTO users (name, email, phone, role_id, password) VALUES ('Thanh Sang', 'sang@gmail.com', '0123456789', 1, '81dc9bdb52d04dc20036dbd8313ed055');
INSERT INTO users (name, email, phone, role_id, password) VALUES ('Clean Bandit', 'clean@gmail.com', '0987654321', 2, '81dc9bdb52d04dc20036dbd8313ed055');
INSERT INTO users (name, email, phone, role_id, password) VALUES ('Dua Lipa', 'dua@gmail.com', '0159763248', 3, '81dc9bdb52d04dc20036dbd8313ed055');
