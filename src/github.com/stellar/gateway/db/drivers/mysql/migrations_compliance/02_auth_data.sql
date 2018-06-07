-- +migrate Up
CREATE TABLE `AuthData` (
  `id` varchar(255) NOT NULL,
  `domain` varchar(255) NOT NULL,
  `auth_data` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `AuthData`;
