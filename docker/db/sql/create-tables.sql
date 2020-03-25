CREATE TABLE `users` (
  `id` bigint NOT NULL,
  `screen_name` varchar(255) NOT NULL DEFAULT '' COMMENT '表示名',
  `screen_id` varchar(15) NOT NULL,
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'email',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'パスワード',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `screen_id` (`screen_id`),
  UNIQUE KEY `screen_id_2` (`screen_id`)
);

CREATE TABLE `tweets` (
  `id` bigint NOT NULL,
  `user_id` int NOT NULL,
  `text` varchar(140) NOT NULL DEFAULT '',
  `tweet_id` int NOT NULL DEFAULT '0',
  `is_comment` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
);

CREATE TABLE `favorites` (
  `user_id` int NOT NULL COMMENT 'ユーザーID',
  `tweet_id` int NOT NULL COMMENT 'ツイートID',
  PRIMARY KEY (`user_id`,`tweet_id`)
);

CREATE TABLE `followers` (
  `following_id` int NOT NULL,
  `followed_id` int NOT NULL,
  UNIQUE KEY `following_id` (`following_id`,`followed_id`)
);