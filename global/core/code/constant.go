package code

var (
	SUCCESS = AppCode{0, "success"}

	UNKONW_ERR    = AppCode{9997, "unkown err: %s"}
	PARAM_INVALID = AppCode{9998, "req param invalid"}
	SERVER_ERR    = AppCode{9999, "server error"}
)
