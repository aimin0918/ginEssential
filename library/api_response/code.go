package api_response

const (
	ReturnSucessCode = 1
	ReturnErrorCode  = -1

	CodeInternal = 100

	CodeSuccess = 200

	CodeInvalidParams = 400
	CodeSignErr       = 401
)

var MessageMap = map[int]string{
	CodeSuccess:       "success",
	CodeInvalidParams: "参数错误",
	CodeSignErr:       "验签失败",
	CodeInternal:      "系统在开小差,请稍后再试",
}
