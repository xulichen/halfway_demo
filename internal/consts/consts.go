package consts

// default value
const (
	DefaultLogPath     = "/var/log/category.log"
	DefaultServiceName = "xxxService"
	DefaultEnvOnline   = "online"
	DefaultEnvTest     = "test"
	DefaultEnvLocal    = "local"
)

const (
	ResponseCodeSuccess             = 20000
	ResponseCodeNeedLogin           = 40003
	ResponseCodeErrParameter        = 40023 // 参数错误
	ResponseCodeInternalServerError = 50000
)
