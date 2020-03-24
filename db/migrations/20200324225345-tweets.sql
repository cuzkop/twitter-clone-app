
-- +migrate Up
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
)

-- +migrate Down
DROP TABLE IF EXISTS `tweets`;