DROP TABLE IF EXISTS `users`;

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

INSERT INTO `users` (`id`, `screen_name`, `screen_id`, `email`, `password`, `created_at`, `updated_at`)
VALUES
	(1, 'kazuki', 'cuzkop', 'cuzkop@gmail.com', 'password', '2020-03-25 20:13:56', '2020-03-25 20:13:56'),
	(2, 'hoge', 'hoge', 'hoge@gmail.com', 'password', '2020-03-25 20:13:56', '2020-03-25 20:13:56'),
	(3, 'huga', 'huga', 'huga@gmail.com', 'password', '2020-03-25 20:13:56', '2020-03-25 20:13:56'),
	(4, 'piyo', 'piyo', 'piyo@gmail.com', 'password', '2020-03-25 20:13:56', '2020-03-25 20:13:56'),
    (5, 'hoo', 'hoo', 'hoo@gmail.com', 'password', '2020-03-25 20:13:56', '2020-03-25 20:13:56');

DROP TABLE IF EXISTS `tweets`;

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

INSERT INTO `tweets` (`id`, `user_id`, `text`, `tweet_id`, `is_comment`, `created_at`, `updated_at`, `is_deleted`)
VALUES
	(1, 1, 'test', 0, 0, '2020-03-25 20:15:06', '2020-03-25 20:15:06', 0),
	(2, 2, 'test', 0, 0, '2020-03-25 20:15:06', '2020-03-25 20:15:06', 0),
	(3, 3, 'test', 0, 0, '2020-03-25 20:15:06', '2020-03-25 20:15:06', 0),
	(4, 4, 'test', 0, 0, '2020-03-25 20:16:11', '2020-03-25 20:16:11', 0),
	(5, 1, 'test2', 0, 0, '2020-03-25 20:16:19', '2020-03-25 20:16:19', 0),
	(6, 1, 'testtest', 4, 1, '2020-03-25 20:16:39', '2020-03-25 20:16:39', 0);

DROP TABLE IF EXISTS `favorites`;

CREATE TABLE `favorites` (
  `user_id` int NOT NULL COMMENT 'ユーザーID',
  `tweet_id` int NOT NULL COMMENT 'ツイートID',
  PRIMARY KEY (`user_id`,`tweet_id`)
);

INSERT INTO `favorites` (`user_id`, `tweet_id`)
VALUES
	(1, 3),
	(2, 1);

DROP TABLE IF EXISTS `followers`;

CREATE TABLE `followers` (
  `following_id` int NOT NULL,
  `followed_id` int NOT NULL,
  UNIQUE KEY `following_id` (`following_id`,`followed_id`)
);

INSERT INTO `followers` (`following_id`, `followed_id`)
VALUES
	(1, 2),
	(1, 3),
	(1, 4),
	(2, 3),
	(3, 4);
