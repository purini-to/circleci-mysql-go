
-- +migrate Up
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- +migrate Down
DROP TABLE IF EXISTS `users`;
