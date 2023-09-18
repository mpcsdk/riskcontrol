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
	CodeTokenInvalid  = &errCode{11, "Token Invalid", nil}  // The token is invalid.
	CodeTokenNotExist = &errCode{12, "Token NotExist", nil} // The token does not exist.
	CodeTFANotExist   = &errCode{12, "TFA NotExist", nil}   // The token does not exist.
	///
	CodeRiskVerification   = &errCode{21, "Risk VerificationCode", nil}      // The risk rule does not exist.
	CodeVerifyCodeInvalid  = &errCode{22, "Verify Code Invalid", nil}        // The verification code is invalid.
	CodeVerifyCodeNotExist = &errCode{23, "Verify RiskSerial NotExist", nil} // The verification code does not exist.
	///
	///
	CodeInternalError = &errCode{50, "Internal Error", nil} // An error occurred internally.

)
