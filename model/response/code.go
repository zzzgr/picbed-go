package response

// ResCode 业务返回码
type ResCode int64

const (
	CodeFail            ResCode = -1  // 响应失败
	CodeSuccess         ResCode = 200 // 响应成功
	BadRequest          ResCode = 400 // 错误请求
	ServerInternalError ResCode = 500 // 服务器内部错误

	FileNotFound     ResCode = 100001
	UploaderNotFound ResCode = 100002
)

var ResCodeMap = map[ResCode]string{

	CodeFail:            "处理失败",
	CodeSuccess:         "处理成功",
	BadRequest:          "请求错误",
	ServerInternalError: "服务器内部错误",

	// 100000
	FileNotFound:     "请选择文件",
	UploaderNotFound: "敬请期待",
}

func (c ResCode) Msg() string {
	msg, ok := ResCodeMap[c]
	if !ok {
		msg = ResCodeMap[ServerInternalError]
	}
	return msg
}
