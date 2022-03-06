package utils

import "time"

const (
	ERRCODE_CUSTOM              = 40001
	ERRCODE_SERVER_ERROR        = 40002
	ERRCODE_DATABASE_ERROR      = 40003
	ERRCODE_REQUEST_PARAM_ERROR = 40004
	ERRCODE_NOT_LOGINED         = 40005
)

const (
	ERRMSG_CUSTOM              = "custom error"
	ERRMSG_SERVER_ERROR        = "server error"
	ERRMSG_DATABASE_ERROR      = "database error"
	ERRMSG_PARAM_ERROR         = "param error"
	ERRMSG_REQUEST_PARAM_ERROR = "request param error"
	ERRMSG_NOT_LOGINED         = "not logined"
)

const (
	TOKEN_IN_HEADER         = "X-Auth-Token"
	USER_TOKEN_REDIS_PREFIX = "USER:TOKEN:"
	TOKEN_EXPIRE            = 60 * time.Minute
)

const (
	EXTRA_INFO = "extra-info"
	LOGINED    = "logined"
)
