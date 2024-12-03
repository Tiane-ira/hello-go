package code

var (
	SUCCESS = AppCode{0, "success"}

	UnknownErr   = AppCode{9997, "unknown err: %s"}
	ParamInvalid = AppCode{9998, "req param invalid"}
	ServerErr    = AppCode{9999, "server error"}
)
