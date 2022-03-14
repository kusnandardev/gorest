package entity

type Transaction struct {
	Id              int
	TransactionType string
	SourceId        string
	DestinationId   string
	Amount          int
	TransactionDate int64 //write unixtime
}
