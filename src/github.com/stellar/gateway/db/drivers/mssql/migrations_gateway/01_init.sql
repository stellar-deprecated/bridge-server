-- +migrate Up
BEGIN TRANSACTION
GO
CREATE TABLE ReceivedPayment
	(
	id bigint NOT NULL IDENTITY (1, 1),
	operation_id varchar(255) NOT NULL,
	processed_at datetime NOT NULL,
	paging_token varchar(255) NOT NULL,
	status varchar(255) NOT NULL
	)  ON [PRIMARY]
GO
ALTER TABLE ReceivedPayment ADD CONSTRAINT
	PK_ReceivedPayment PRIMARY KEY CLUSTERED 
	(
	id
	) WITH( STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]

GO
ALTER TABLE ReceivedPayment SET (LOCK_ESCALATION = TABLE)
GO
CREATE TABLE SentTransaction
	(
	id bigint NOT NULL IDENTITY (1, 1),
	transaction_id varchar(255) NOT NULL,
	status varchar(10) NOT NULL,
	source varchar(56) NOT NULL,
	submitted_at datetime NOT NULL,
	succeeded_at datetime NOT NULL,
	ledger bigint NOT NULL,
	envelope_xdr text NOT NULL,
	result_xdr varchar(255) NOT NULL
	)  ON [PRIMARY]
	 TEXTIMAGE_ON [PRIMARY]
GO
ALTER TABLE SentTransaction ADD CONSTRAINT
	PK_SentTransaction PRIMARY KEY CLUSTERED 
	(
	id
	) WITH( STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]

GO
ALTER TABLE SentTransaction SET (LOCK_ESCALATION = TABLE)
GO
COMMIT

-- +migrate Down
