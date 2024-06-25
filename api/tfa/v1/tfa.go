package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TfaInfoReq struct {
	g.Meta `path:"/tfaInfo" tags:"tfaInfo" method:"post" summary:"tfaInfo"`
	Token  string `json:"token"`
}
type TfaInfoRes struct {
	g.Meta       `mime:"text/html" example:"string"`
	Phone        string `json:"phone"`
	UpPhoneTime  string `json:"upPhoneTime"`
	Mail         string `json:"mail"`
	UpMailTime   string `json:"upMailTime"`
	TxNeedVerify bool   `json:"txNeedVerify"`
}

// /
// //
type SendSmsCodeReq struct {
	g.Meta     `path:"/sendSmsCode" tags:"sendSmsCode" method:"post" summary:"sendSmsCode"`
	Token      string `json:"token"`
	RiskSerial string `json:"riskSerial"`
	Phone      string `json:"phone"`
}
type SendSmsCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

////

type SendMailCodeReq struct {
	g.Meta     `path:"/sendMailCode" tags:"sendMailCode" method:"post" summary:"sendMailCode"`
	Token      string `json:"token"`
	RiskSerial string `json:"riskSerial"`
	Mail       string `json:"mail"`
}
type SendMailCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type VerifyCodeReq struct {
	g.Meta     `path:"/verifyCode" tags:"verifyCode" method:"post" summary:"verifyCode"`
	Token      string `json:"token"`
	RiskSerial string `json:"riskSerial"`
	PhoneCode  string `json:"phoneCode"`
	MailCode   string `json:"mailCode"`
}
type VerifyCodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// /

type RequestData struct {
	Enable      bool   `json:"enable"`
	ChainId     string `json:"chainId"`
	SignDataStr string `json:"signData"`
	UserId      string `json:"userId"`
}
type TfaRequestReq struct {
	g.Meta   `path:"/tfaRequest" tags:"tfaRequest" method:"post" summary:"tfaRequest"`
	CodeType string       `json:"codeType" "in:bindPhone,bindMail,updatePhone,updateMail"`
	Data     *RequestData `json:"data"`
	Token    string       `json:"token"`
}
type TfaRequestRes struct {
	g.Meta     `mime:"text/html" example:"string"`
	Ok         int32    `json:"ok"`
	RiskSerial string   `json:"riskSerial"`
	VList      []string `json:"vlist"`
}
