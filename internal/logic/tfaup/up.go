package tfaup

import (
	"context"
	"errors"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sTFA struct {
	ctx context.Context
	// riskClient riskv1.UserClient
	pendding map[string]func()
}

func (s *sTFA) VerifyCode(ctx context.Context, token string, kind, code string) error {
	// 验证验证码
	var err error
	if task, ok := s.pendding[token+kind+code]; ok {
		task()
		delete(s.pendding, token+kind+code)
		return err
	}
	return errors.New("nocode")
}
func (s *sTFA) upPhone(token, phone string) error {
	_, err := dao.Tfa.Ctx(s.ctx).Where(dao.Tfa.Columns().Token, token).Update(entity.Tfa{Phone: phone})
	return err
}
func (s *sTFA) UpPhone(ctx context.Context, token string, phone string) error {

	//todo:
	rst, err := service.Risk().PerformRisk(ctx, "phone", nil)
	if err != nil {
		s.pendding[token+"kind"+"phone"] = func() {
			s.UpPhone(ctx, token, phone)
		}
		return err
	}
	g.Log().Info(ctx, "UpPhone risk:", rst)
	err = s.upPhone(token, phone)
	return err
}
func (s *sTFA) upMail(token, mail string) error {
	_, err := dao.Tfa.Ctx(s.ctx).Where(dao.Tfa.Columns().Token, token).Update(entity.Tfa{Mail: mail})
	return err
}
func (s *sTFA) UpMail(ctx context.Context, token string, mail string) error {
	//todo:
	rst, err := service.Risk().PerformRisk(ctx, "mail", nil)
	if err != nil {
		s.pendding[token+"kind"+"mail"] = func() {
			s.upMail(token, mail)
		}
		return err
	}
	g.Log().Info(ctx, "UpPhone risk:", rst)
	s.upMail(token, mail)
	return err
}

func new() *sTFA {

	ctx := gctx.GetInitCtx()
	// addr, err := gcfg.Instance().Get(ctx, "etcd.address")
	// if err != nil {
	// 	panic(err)
	// }
	// grpcx.Resolver.Register(etcd.New(addr.String()))
	// conn, err := grpcx.Client.NewGrpcClientConn("rulerpc")
	// if err != nil {
	// 	panic(err)
	// }
	// client := risk.NewUserClient(conn)
	return &sTFA{
		ctx: ctx,
	}

}

func init() {
	service.RegisterTFA(new())
}
