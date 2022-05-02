package utils

import "time"

const (
	ERRCODE_CUSTOM              = 40001
	ERRCODE_SERVER_ERROR        = 40002
	ERRCODE_DATABASE_ERROR      = 40003
	ERRCODE_REQUEST_PARAM_ERROR = 40004
	ERRCODE_NOT_LOGINED         = 40005
	ERRCODE_USER_NOT_EXIST      = 40006
	ERRCODE_INVALID_SESSION     = 40101
	SUCCESS_CODE_CUSTOM         = 2000
)

const (
	ERRMSG_CUSTOM              = "custom error"
	ERRMSG_SERVER_ERROR        = "server error"
	ERRMSG_DATABASE_ERROR      = "database error"
	ERRMSG_PARAM_ERROR         = "param error"
	ERRMSG_REQUEST_PARAM_ERROR = "request param error"
	ERRMSG_NOT_LOGINED         = "not logined"
	ERRMSG_USER_NOT_EXIST      = "user not exist"
	ERRMSG_INVALID_SESSION     = "invalid session"
	ERRMSG_INVALID_Id          = "invalid id"
	SUCCESS_MSG_CUSTOM         = "success"
)

const (
	TOKEN_IN_HEADER         = "X-Auth-Token"
	USER_TOKEN_REDIS_PREFIX = "USER:TOKEN:"
	USER_SID_REDIS_PREFIX   = "USER:SID:"
	TOKEN_EXPIRE            = 24 * time.Hour
	REFRESH_EXPIRE          = 2 * TOKEN_EXPIRE
)

const (
	EXTRA_INFO = "extra-info"
	LOGINED    = "logined"
)

const (
	SESSION_COOKIE_NAME = "SID"
	CLIENT_KEY          = "client"
	UID_KEY             = "UID"
)
