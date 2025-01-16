package orm

const (
	// Success indicates that the operation was successful.
	Success = 0
	// InternalServerError represents a fatal error occurring on the server.
	InternalServerError = 500
)

type TokenType int

const (
	ETH TokenType = iota
	ERC20
	ERC721
	ERC1155
)

type EventType int

const ()
