CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `passkey` longtext,
  `balance` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`username`, `passkey`, `balance`) VALUES
('johndoe', '$2a$14$uQpPUQ5NgR7M1BEZeHiOcO0syd8oHPTPNTfDHcVtbwO5n4M57EWo6', 100000000);
-- passkey: secret

CREATE USER 'coba'@'localhost' IDENTIFIED BY 'cobapass';
GRANT ALL PRIVILEGES ON * . * TO 'coba'@'localhost';
FLUSH PRIVILEGES;