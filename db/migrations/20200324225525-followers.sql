
-- +migrate Up
CREATE TABLE `followers` (
  `following_id` int NOT NULL,
  `followed_id` int NOT NULL,
  UNIQUE KEY `following_id` (`following_id`,`followed_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `followers`;