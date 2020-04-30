package kernel

const (
	SuccessNum       = 0    // OK
	IgnoreNum        = 200  // -- 禁用，易误解成Http.SC_OK --
	Forbidden        = 403  // 无访问权限
	PageNotFound     = 404  // 页面不存在
	MethodNotAllowed = 405  // 请求方式不支持
	SystemRunErr     = 500  // 系统运行异常
	ScriptRunErr     = 501  // 脚本运行失败
	Idempotent       = 1000 // 幂等
	Conflict         = 1001 // 重复请求
	ArgsErr          = 1101 // 参数错误
	ResultErr        = 1102 // 结果错误
	JsonSyntaxErr    = 1201 // Json解析异常
	XmlSyntaxErr     = 1202 // Xml解析异常
	UnknownErr       = 2008 // 未知错误
)
