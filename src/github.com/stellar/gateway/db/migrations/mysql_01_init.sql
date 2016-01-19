-- +migrate Up
CREATE TABLE `ReceivedPayment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operation_id` varchar(255) NOT NULL,
  `processed_at` datetime NOT NULL,
  `paging_token` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `SentTransaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` varchar(10) NOT NULL,
  `source` varchar(56) NOT NULL,
  `submitted_at` datetime NOT NULL,
  `succeeded_at` datetime DEFAULT NULL,
  `ledger` bigint(20) DEFAULT NULL,
  `envelope_xdr` text NOT NULL,
  `result_xdr` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE `ReceivedPayment`;
DROP TABLE `SentTransaction`;
