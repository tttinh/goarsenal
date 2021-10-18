package errorcode

type Code string

const (
	OK                   Code = "OK"
	ERROR_INVALID_PARAMS Code = "ERROR_INVALID_PARAMS"
	ERROR_LIST_WAGERS    Code = "ERROR_LIST_WAGERS"
	ERROR_CREATE_WAGER   Code = "ERROR_CREATE_WAGER"
	ERROR_BUY_WAGER      Code = "ERROR_BUY_WAGER"
)
