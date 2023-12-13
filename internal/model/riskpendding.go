package model

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type RiskVerifyPendding struct {
	RiskKind RiskKind
	//风控序号
	RiskSerial string
	//用户id
	UserId string
	///
	riskBeforFunc []func(context.Context) error
	// riskVerify    map[model.RiskKind]*riskVerify
	verifier map[VerifyKind]IVerifier
	// sender        map[VerifyKind]sender
	riskAfterFunc []func(context.Context) error
	///
	phoneSender int
	mailSender  int
	///
	dealline *gtime.Time
}

// //

func (s *RiskVerifyPendding) Verifier(kind VerifyKind) IVerifier {
	if v, ok := s.verifier[kind]; ok {
		return v
	}
	return nil
}

func (s *RiskVerifyPendding) Verifiers() map[VerifyKind]IVerifier {
	return s.verifier
}
func (s *RiskVerifyPendding) AddAfterFunc(after func(context.Context) error) {
	if after == nil {
		return
	}
	s.riskAfterFunc = append(s.riskAfterFunc, after)
}
func (s *RiskVerifyPendding) AddBeforFunc(befor func(context.Context) error) {
	if befor == nil {
		return
	}
	s.riskBeforFunc = append(s.riskBeforFunc, befor)
}

func (s *RiskVerifyPendding) AddVerifier(verifier IVerifier) {
	s.verifier[verifier.VerifyKind()] = verifier
}

func (s *RiskVerifyPendding) VerifierCode(code *VerifyCode) (VerifyKind, error) {
	for k, v := range s.verifier {
		if _, err := v.Verify(code); err != nil {
			return k, err
		}
	}
	return "", nil
}

// /
func (s *RiskVerifyPendding) DoFunc(ctx context.Context) (VerifyKind, error) {
	if k, err := s.AllDone(); err != nil {
		return k, err
	} else {
		s.DoBefor(ctx)
		if k, err := s.DoAfter(ctx); err != nil {
			err = gerror.Wrap(err, mpccode.ErrDetails(
				mpccode.ErrDetail("k", k),
			))
			return "", err
		}
		////notice: completion info
		for _, v := range s.verifier {
			err := v.SendCompletion()
			if err != nil {
				g.Log().Errorf(ctx, "%+v", err)
			}
		}
	}
	return "", nil
}

// /
func (s *RiskVerifyPendding) DoBefor(ctx context.Context) (string, error) {
	for _, f := range s.riskBeforFunc {
		f(ctx)
	}
	return "", nil
}

func (s *RiskVerifyPendding) DoAfter(ctx context.Context) (string, error) {
	for _, verifer := range s.verifier {
		if !verifer.IsDone() {
			return string(verifer.VerifyKind()), gerror.NewCode(mpccode.CodeRiskVerifyCodeInvalid)
		}
	}
	//done
	for _, task := range s.riskAfterFunc {
		err := task(ctx)
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
func (s *RiskVerifyPendding) AllDone() (VerifyKind, error) {
	for _, e := range s.verifier {
		if e.IsDone() {
			continue
		}
		return e.VerifyKind(), errRiskNotDone
	}
	return "", nil
}
