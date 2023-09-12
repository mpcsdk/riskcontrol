package sms

import (
	"context"
	"riskcontral/common"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
)

type sSmsCode struct {
	sms  *huawei
	pool *grpool.Pool
}

func (s *sSmsCode) sendCode(ctx context.Context, receiver, code string) error {

	resp, status, err := s.sms.sendSms(receiver, code)
	g.Log().Info(ctx, "sendcode:", resp, status, err)
	///
	return err
}

func (s *sSmsCode) SendCode(ctx context.Context, receiver string) (string, error) {

	code := common.RandomDigits(6)
	///
	return code, s.sendCode(ctx, receiver, code)
}

// func (s *sSmsCode) Verify(ctx context.Context, sid, code string) error {
// 	c, err := service.Cache().Get(ctx, sid+"smscode")
// 	if err == nil {
// 		if c.String() != code {
// 			return errors.New("verfiy fauild")
// 		}
// 	}
// 	//todo: faild stat
// 	stat, err := service.Cache().Get(ctx, sid+"sms")
// 	if err != nil {
// 		return err
// 	}
// 	if stat.String() == "err" {
// 		estr, err := service.Cache().Get(ctx, sid+"smserr")
// 		if err != nil {
// 			return err
// 		}
// 		return errors.New(estr.String())
// 	}

//		status, err := service.Cache().Get(ctx, sid+"smsstatus")
//		if err != nil {
//			return err
//		}
//		///
//		fmt.Println(status.String())
//		return nil
//	}
func new() *sSmsCode {

	return &sSmsCode{
		pool: grpool.New(10),
		sms:  newhuawei(),
	}
}

func init() {
	service.RegisterSmsCode(new())
}
