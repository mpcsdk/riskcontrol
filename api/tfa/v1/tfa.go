package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type TFAInfoReq struct {
	g.Meta `path:"/tfaInfo" tags:"tfaInfo" method:"post" summary:"tfaInfo"`
	Token  string `json:"token"`
}
type TFAInfoRes struct {
	g.Meta      `mime:"text/html" example:"string"`
	Phone       string      `json:"phone"`
	UpPhoneTime *gtime.Time `json:"upPhoneTime"`
	Mail        string      `json:"mail"`
	UpMailTime  *gtime.Time `json:"upMailTime"`
}

// /
// //
type SendSmsCodeReq struct {
	g.Meta     `path:"/sendSmsCode" tags:"sendSmsCode" method:"post" summary:"sendSmsCode"`
	Token      string `json:"token"`
	RiskSerial string `json:"riskSerial"`
}
type SendSmsCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

////

type SendMailCodeReq struct {
	g.Meta     `path:"/sendMailCode" tags:"sendMailCode" method:"post" summary:"sendMailCode"`
	Token      string `json:"token"`
	RiskSerial string `json:"riskSerial"`
}
type SendMailCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type VerifyReq struct {
	RiskSerial string `json:"riskSerial"`
	Code       string `json:"code"`
}
type VerifyCodeReq struct {
	g.Meta    `path:"/verifyCode" tags:"verifyCode" method:"post" summary:"verifyCode"`
	Token     string       `json:"token"`
	VerifyReq []*VerifyReq `json:"codes"`
}
type VerifyCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// /

type UpPhoneReq struct {
	g.Meta `path:"/upPhone" tags:"upPhone" method:"post" summary:"upPhone"`
	Token  string `json:"token"`
	Phone  string `json:"phone"`
}
type UpPhoneRes struct {
	g.Meta     `mime:"text/html" example:"string"`
	RiskSerial string `json:"riskSerial"`
}
type UpMailReq struct {
	g.Meta `path:"/upMail" tags:"upMail" method:"post" summary:"upMail"`
	Token  string `json:"token"`
	Mail   string `json:"mail"`
}
type UpMailRes struct {
	g.Meta     `mime:"text/html" example:"string"`
	RiskSerial string `json:"riskSerial"`
}

// /
type DialCodeReq struct {
	g.Meta `path:"/dialCode" tags:"dialCode" method:"post" summary:"dialCode"`
}
type DialCode struct {
	Name     string `json:"name"`
	En       string `json:"en"`
	DialCode string `json:"dial_code"`
}
type DialCodeRes struct {
	g.Meta    `mime:"text/html" example:"string"`
	DialCodes []*DialCode `json:"dial_codes"`
}

// notice: debug
type CreateTFAReq struct {
	g.Meta `path:"/createTFA" tags:"createTFA" method:"post" summary:"createTFA"`
	Token  string `json:"token"`
	Phone  string `json:"phone"`
	Mail   string `json:"mail"`
}
type CreateTFARes struct {
	g.Meta `mime:"text/html" example:"string"`
}
