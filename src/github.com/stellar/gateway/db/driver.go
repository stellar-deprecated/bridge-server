package db

type Driver interface {
	Init(url string) (err error)
	MigrateUp() (migrationsApplied int, err error)

	InsertReceivedPayment(object *ReceivedPayment) (id int64, err error)
	UpdateReceivedPayment(object *ReceivedPayment) (err error)
	GetLastReceivedPayment() (*ReceivedPayment, error)

	InsertSentTransaction(object *SentTransaction) (id int64, err error)
	UpdateSentTransaction(object *SentTransaction) (err error)
}
