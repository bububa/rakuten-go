package enum

// MessgeType はメッセージ種別です。以下のいずれかが入力されます。
type MessageType string

const (
	// MessageType_INFO
	MessageType_INFO MessageType = "INFO"
	// MessageType_ERROR
	MessageType_Error MessageType = "ERROR"
	// MessageType_WARNING
	MessageType_WARNING MessageType = "WARNING"
)
