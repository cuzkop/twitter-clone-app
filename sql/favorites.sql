CREATE TABLE `favorites` (
  `user_id` int NOT NULL COMMENT 'ユーザーID',
  `tweet_id` int NOT NULL COMMENT 'ツイートID',
  PRIMARY KEY (`user_id`,`tweet_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;