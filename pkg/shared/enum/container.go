package enum

type ContainerName string

const (
	CustomerContainer    ContainerName = "customer-handler"
	JWTContainer         ContainerName = "jwt-handler"
	TransactionContainer ContainerName = "transaction-handler"

	CacheContainer ContainerName = "cache-handler"
)
