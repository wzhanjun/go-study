package errcode

// ErrCode 错误码
type ErrCode int

//go:generate stringer -type ErrCode -linecomment -output code_string.go
const (
	ERR_CODE_OK             ErrCode = 0    // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1001 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2001 // 请求超时
)
