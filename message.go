package kernel

var messages = map[uint]string{
	SuccessNum:       "ok",
	Forbidden:        "forbidden",
	PageNotFound:     "page not found",
	MethodNotAllowed: "method not allowed",
	SystemRunErr:     "system run error",
	ScriptRunErr:     "script run error",
	Idempotent:       "idempotent",
	Conflict:         "conflict",
	ArgsErr:          "args error",
	ResultErr:        "result error",
	JsonSyntaxErr:    "json syntax error",
	XmlSyntaxErr:     "xml syntax error",
	UnknownErr:       "unknown error",
}

func GetMessage(code uint) string {
	if r, ok := messages[code]; ok {
		return r
	}
	return ""
}
