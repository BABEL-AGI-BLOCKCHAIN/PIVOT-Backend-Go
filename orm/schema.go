package orm

const (
	// Success indicates that the operation was successful.
	Success = 0
	// InternalServerError represents a fatal error occurring on the server.
	InternalServerError = 500
)

type EventType int

const (
	EventTypeCreateTopic EventType = iota + 1
	EventTypeInvest
	EventTypeWithdraw
	EventTypeWithdrawCommission
)

type ProcessStatus int

const (
	UnProcessed ProcessStatus = iota + 1
	Processed
	ProcessFailed
)
