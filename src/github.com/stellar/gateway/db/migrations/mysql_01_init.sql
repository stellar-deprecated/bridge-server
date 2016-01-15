-- +migrate Up
CREATE TABLE `SentTransaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` varchar(10) NOT NULL,
  `source` varchar(56) NOT NULL,
  `submitted_at` datetime NOT NULL,
  `succeded_at` datetime DEFAULT NULL,
  `operation_type` varchar(20) NOT NULL,
  `ledger` bigint(20) DEFAULT NULL,
  `enveloper_xdr` text NOT NULL,
  `result_xdr` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `SentTransaction`;
