package e

// Msgs the flags to define error massage string
var Msgs = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_USER:           "该用户不存在",
	ERROR_NOT_EXIST_NOTE:           "该笔记不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// Msg get the message detail by specified error code
func Msg(code int) string {
	msg, ok := Msgs[code]
	if ok {
		return msg
	}

	return Msgs[ERROR]
}
