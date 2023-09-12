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

type SendSmsCodeReq struct {
	g.Meta `path:"/sendSmsCode" tags:"sendSmsCode" method:"post" summary:"sendSmsCode"`
	Token  string `json:"token"`
	Kind   string `json:"kind"`
}
type SendSmsCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Code   string
}
type VerifySmsCodeReq struct {
	g.Meta `path:"/verifySmsCode" tags:"verifySmsCode" method:"post" summary:"verifySmsCode"`
	Token  string `json:"token"`
	Kind   string `json:"kind"`
	Code   string `json:"code"`
}
type VerifySmsCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Code   string
}

////

type SendMailOTPReq struct {
	g.Meta `path:"/sendMailOTP" tags:"sendMailOTP" method:"post" summary:"sendMailOTP"`
	Token  string `json:"token"`
	Kind   string `json:"kind"`
}
type SendMailOTPRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Code   string
}
type VerifyMailOTPReq struct {
	g.Meta `path:"/verifyMailOTP" tags:"verifyMailOTP" method:"post" summary:"verifyMailOTP"`
	Token  string `json:"token"`
	Code   string `json:"code"`
	Kind   string `json:"kind"`
}
type VerifyMailOTPRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Code   string
}

// /
// //
type CreateTFAReq struct {
	g.Meta `path:"/createTFA" tags:"createTFA" method:"post" summary:"createTFA"`
	Token  string `json:"token"`
	Phone  string `json:"phone"`
	Mail   string `json:"mail"`
}
type CreateTFARes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UpPhoneReq struct {
	g.Meta `path:"/upPhone" tags:"upPhone" method:"post" summary:"upPhone"`
	Token  string `json:"token"`
	Phone  string `json:"phone"`
}
type UpPhoneRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
type UpMailReq struct {
	g.Meta `path:"/upMail" tags:"upMail" method:"post" summary:"upMail"`
	Token  string `json:"token"`
	Mail   string `json:"mail"`
}
type UpMailRes struct {
	g.Meta `mime:"text/html" example:"string"`
}