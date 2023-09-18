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
	SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error)
	VerifyMailCode(ctx context.Context, req *v1.VerifyMailCodeReq) (res *v1.VerifyMailCodeRes, err error)
	UpPhone(ctx context.Context, req *v1.UpPhoneReq) (res *v1.UpPhoneRes, err error)
	UpMail(ctx context.Context, req *v1.UpMailReq) (res *v1.UpMailRes, err error)
}


