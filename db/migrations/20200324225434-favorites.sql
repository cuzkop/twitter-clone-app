
-- +migrate Up
CREATE TABLE `favorites` (
  `user_id` int NOT NULL COMMENT 'ユーザーID',
  `tweet_id` int NOT NULL COMMENT 'ツイートID',
  PRIMARY KEY (`user_id`,`tweet_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `favorites`;