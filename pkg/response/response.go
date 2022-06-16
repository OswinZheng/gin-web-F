package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func result(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func (res *Response) WithMsg(message string) *Response {
	res.Msg = message
	return res
}

func (res *Response) WithData(data interface{}) *Response {
	res.Data = data
	return res
}
