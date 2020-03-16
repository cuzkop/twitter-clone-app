CREATE TABLE `followers` (
  `following_id` int NOT NULL,
  `followed_id` int NOT NULL,
  UNIQUE KEY `following_id` (`following_id`,`followed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;