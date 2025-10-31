package enum

type TransactionType string

const (
	RealtimeTransaction TransactionType = "realtime"
	BatchTransaction    TransactionType = "batch"
)
