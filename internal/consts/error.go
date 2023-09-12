package consts

type errCode struct {
	code    int
	message string
	detail  interface{}
}

func (e *errCode) Error() string {
	return e.message
}
func (e *errCode) Message() string {
	return e.message
}
func (e *errCode) Code() int {
	return e.code
}
func (e *errCode) Detail() interface{} {
	return e.detail
}

var (
	CodeNil = &errCode{-1, "nil", nil} // No error code specified.
	CodeOK  = &errCode{0, "ok", nil}   // It is OK.
	///
	CodeTokenUnExist = &errCode{12, "Token UnExist", nil} // The token does not exist.
	CodeTokenInvalid = &errCode{11, "Token Invalid", nil} // The token is invalid.
	///
	CodeVerifyCodeInvalid  = &errCode{21, "Verify Code Invalid", nil}  // The verification code is invalid.
	CodeVerifyCodeNotExist = &errCode{22, "Verify Code NotExist", nil} // The verification code does not exist.
	///
	CodeRiskRuleNotExist = &errCode{31, "Risk Rule NotExist", nil} // The risk rule does not exist.
	///
	CodeAuthFailed = &errCode{41, "auth failed", nil}
	///
	CodeInternalError = &errCode{50, "Internal Error", nil} // An error occurred internally.

)
