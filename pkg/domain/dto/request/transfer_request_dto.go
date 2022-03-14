package request

type TransferRequestDto struct {
	SourceId      string
	DestinationId string `json:"destination_id"`
	Amount        int
}
