package orm

const (
	// Success indicates that the operation was successful.
	Success = 0
	// InternalServerError represents a fatal error occurring on the server.
	InternalServerError = 500
	// ErrParameterInvalidNo represents an error when the parameters are invalid.
	ErrParameterInvalidNo = 40001
	// ErrGetL2ClaimableWithdrawalsError represents an error when trying to get L2 claimable withdrawal transactions.
	ErrGetL2ClaimableWithdrawalsError = 40002
	// ErrGetTxsError represents an error when trying to get transactions by address.
	ErrGetTxsError = 40003
)

type TokenType int

const (
	ETH TokenType = iota
	ERC20
	ERC721
	ERC1155
	RED
)

type TxType int

const (
	TxTypeUnknown TxType = iota
	TxTypeDeposit
	TxTypeWithdraw
	TxTypeRefund
)

type TxStatusType int

// Constants for TxStatusType.
const (
	TxStatusTypeSent TxStatusType = iota
	TxStatusTypeConsumed
	TxStatusTypeDropped
)

// MessageType represents the type of message.
type MessageType int

// Constants for MessageType.
const (
	MessageTypeUnknown MessageType = iota
	MessageTypeL1SentMessage
	MessageTypeL2SentMessage
)

type EventType int

const (
	QueueTransaction EventType = iota + 1 // 1. QueueTransaction (L1DepositMsgSent)
	L2RelayedMessage                      // 2. L2RelayedMessage (L2DepositMsgConsumed)
	SentMessage                           // 3. SentMessage (L2withdrawMsgSent)
	L1RelayedMessage                      // 4. L1RelayedMessage (L2DepositMsgConsumed)
)
