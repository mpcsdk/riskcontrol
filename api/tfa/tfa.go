// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package tfa

import (
	"context"
	
	"riskcontral/api/tfa/v1"
)

type ITfaV1 interface {
	TfaInfo(ctx context.Context, req *v1.TfaInfoReq) (res *v1.TfaInfoRes, err error)
	SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error)
	SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error)
	VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error)
	TfaRequest(ctx context.Context, req *v1.TfaRequestReq) (res *v1.TfaRequestRes, err error)
}


