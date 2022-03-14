package response

type TransferResponseDto struct {
	TransactionId   int
	DestinationId   string
	DestinationName string
	Amount          int
	Balance         int
}
