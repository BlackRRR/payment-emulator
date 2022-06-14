package model

type Result int32

const (
	ResultOK  Result = 200
	ResultERR Result = 400

	StatusNew     = "new"
	StatusSuccess = "success"
	StatusFailure = "failure"
	StatusError   = "error"
)
