package def

import "strings"

type ResponseInterface interface {
	Error() *ErrorInfo
	IsError() bool
}

type Response struct {
	Errinf *ErrorInfo `json:"error_info" xml:"errinf"`
}

func (res Response) Error() *ErrorInfo {
	return res.Errinf
}

// IsError (Error() != nil)
func (res Response) IsError() bool {
	return res.Error() != nil
}

type ErrorInfo struct {
	Errcd  string `json:"error_code" xml:"errcd"`
	Errmsg string `json:"error_message" xml:"errmsg"`
}

func (e *ErrorInfo) ErrorCode() string {
	if e != nil {
		return strings.TrimSpace(e.Errcd)
	}
	return ""
}

func (e *ErrorInfo) ErrorMessage() string {
	if e != nil {
		return strings.TrimSpace(e.Errmsg)
	}
	return ""
}
