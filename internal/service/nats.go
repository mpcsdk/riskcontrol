// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"

	"github.com/golang/protobuf/ptypes/empty"
)

type (
	INats interface {
		PerformAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error)
		PerformAllAbi(ctx context.Context, req *v1.AllAbiReq) (res *v1.AllAbiRes, err error)
		PerformAllNftRules(ctx context.Context, req *v1.NftRulesReq) (res *v1.NftRulesRes, err error)
		PerformAllFtRules(ctx context.Context, req *v1.FtRulesReq) (res *v1.FtRulesRes, err error)
		PerformRiskTFA(ctx context.Context, req *v1.TFARiskReq) (res *v1.TFARiskRes, err error)
		PerformSmsCode(ctx context.Context, req *v1.SmsCodeReq) (res *v1.SmsCodeRes, err error)
		PerformMailCode(ctx context.Context, req *v1.MailCodekReq) (res *v1.MailCodekRes, err error)
		PerformVerifyCode(ctx context.Context, req *v1.VerifyCodekReq) (res *v1.VerifyCodeRes, err error)
		PerformRiskTxs(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error)
	}
)

var (
	localNats INats
)

func Nats() INats {
	if localNats == nil {
		panic("implement not found for interface INats, forgot register?")
	}
	return localNats
}

func RegisterNats(i INats) {
	localNats = i
}
