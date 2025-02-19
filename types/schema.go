package types

const (
	Sepolia = 11155111
	// Success indicates that the operation was successful.
	Success = 0
	// InternalServerError represents a fatal error occurring on the server.
	InternalServerError = 500
	// ErrParameterInvalidNo represents an error when the parameters are invalid.
	ErrParameterInvalidNo = 40001
	// ErrGetL2ClaimableWithdrawalsError represents an error when trying to get L2 claimable withdrawal transactions.
	ErrGetTxsError = 40003
)

type ProcessStatus int

const (
	UnProcessed ProcessStatus = iota + 1
	Processed
	ProcessFailed
)

type EventType int

const (
	CreateTopic EventType = iota + 1
	Invest
	Withdraw
	WithdrawCommission
)
