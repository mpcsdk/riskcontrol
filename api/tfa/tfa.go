// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package tfa

import (
	"context"
	
	"riskcontral/api/tfa/v1"
)

type ITfaV1 interface {
	TFAInfo(ctx context.Context, req *v1.TFAInfoReq) (res *v1.TFAInfoRes, err error)
	SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error)
	VerifySmsCode(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error)
	SendMailOTP(ctx context.Context, req *v1.SendMailOTPReq) (res *v1.SendMailOTPRes, err error)
	VerifyMailOTP(ctx context.Context, req *v1.VerifyMailOTPReq) (res *v1.VerifyMailOTPRes, err error)
	CreateTFA(ctx context.Context, req *v1.CreateTFAReq) (res *v1.CreateTFARes, err error)
	UpPhone(ctx context.Context, req *v1.UpPhoneReq) (res *v1.UpPhoneRes, err error)
	UpMail(ctx context.Context, req *v1.UpMailReq) (res *v1.UpMailRes, err error)
}


