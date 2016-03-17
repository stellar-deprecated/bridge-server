-- +migrate Up
CREATE TABLE `AuthorizedTransaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_id` char(64) NOT NULL,
  `memo` varchar(64) NULL,
  `transaction_xdr` text NOT NULL,
  `authorized_at` datetime NOT NULL,
  `data` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `AuthorizedTransaction`;
