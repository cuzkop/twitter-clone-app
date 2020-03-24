
-- +migrate Up
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

-- +migrate Down
DROP TABLE IF EXISTS `users`;
