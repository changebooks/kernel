package kernel

import "sync"

type Result struct {
	mu      sync.Mutex  // protects following fields
	code    uint        // 错误码，0：正确、> 0：错误
	message string      // 错误消息
	data    interface{} // 内容
}

func NewSuccess(data interface{}) *Result {
	return NewResult(SuccessNum, GetMessage(SuccessNum), data)
}

func NewSystemRunErr(data interface{}) *Result {
	return NewResult(SystemRunErr, GetMessage(SystemRunErr), data)
}

func NewPageNotFound(data interface{}) *Result {
	return NewResult(PageNotFound, GetMessage(PageNotFound), data)
}

func NewMethodNotAllowed(data interface{}) *Result {
	return NewResult(MethodNotAllowed, GetMessage(MethodNotAllowed), data)
}

func NewArgsErr(data interface{}) *Result {
	return NewResult(ArgsErr, GetMessage(ArgsErr), data)
}

func NewError(err error, data interface{}) *Result {
	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = GetMessage(SystemRunErr)
	}

	return NewResult(SystemRunErr, message, data)
}

func NewResult(code uint, message string, data interface{}) *Result {
	return &Result{
		code:    code,
		message: message,
		data:    data,
	}
}

func (x *Result) Success() bool {
	return x.code == SuccessNum
}

func (x *Result) Error() bool {
	return x.code != SuccessNum
}

func (x *Result) SystemRunErr() bool {
	return x.code == SystemRunErr
}

func (x *Result) GetCode() uint {
	return x.code
}

func (x *Result) SetCode(n uint) *Result {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.code = n
	return x
}

func (x *Result) GetMessage() string {
	return x.message
}

func (x *Result) SetMessage(s string) *Result {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.message = s
	return x
}

func (x *Result) GetData() interface{} {
	return x.data
}

func (x *Result) SetData(o interface{}) *Result {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.data = o
	return x
}
