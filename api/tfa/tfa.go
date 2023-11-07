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
	SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error)
	VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error)
	TfaRequest(ctx context.Context, req *v1.TfaRequestReq) (res *v1.TfaRequestRes, err error)
	DialCode(ctx context.Context, req *v1.DialCodeReq) (res *v1.DialCodeRes, err error)
}


