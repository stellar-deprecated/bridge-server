-- +migrate Up
CREATE TABLE `AuthorizedTransaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_id` char(64) NOT NULL,
  `memo` varchar(64) NOT NULL,
  `transaction_xdr` text NOT NULL,
  `authorized_at` datetime NOT NULL,
  `data` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `AllowedFI` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `domain` varchar(255) NOT NULL,
  `public_key` char(56) NOT NULL,
  `allowed_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `domain` (`domain`),
  UNIQUE KEY `public_key` (`public_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `AllowedUser` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fi_name` varchar(255) NOT NULL,
  `fi_domain` varchar(255) NOT NULL,
  `fi_public_key` char(56) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `allowed_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `fi_public_key_user_id` (`fi_public_key`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `AuthorizedTransaction`;
DROP TABLE `AllowedFI`;
DROP TABLE `AllowedUser`;
