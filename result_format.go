package kernel

import "errors"

// Result -> Json
func ResultFormat(r *Result) ([]byte, error) {
	if r == nil {
		return nil, errors.New("r can't be nil")
	}

	if r.Success() {
		return JsonSuccessEncode(r.GetData())
	} else {
		return JsonErrorEncode(r.GetCode(), r.GetMessage(), r.GetData())
	}
}
