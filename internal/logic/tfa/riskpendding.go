package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// //
// //
type riskVerifyPendding struct {
	RiskKind RiskKind
	//风控序号
	RiskSerial string
	//用户id
	UserId string
	///
	riskBeforFunc []func(context.Context) error
	// riskVerify    map[RiskKind]*riskVerify
	verifier map[VerifyKind]IVerifier
	// sender        map[VerifyKind]sender
	riskAfterFunc []func(context.Context) error
	///
	phoneSender int
	mailSender  int
	///
	dealline *gtime.Time
}

func (s *riskVerifyPendding) Verifiers() map[VerifyKind]IVerifier {
	return s.verifier
}
func (s *riskVerifyPendding) AddAfterFunc(after func(context.Context) error) {
	if after == nil {
		return
	}
	s.riskAfterFunc = append(s.riskAfterFunc, after)
}
func (s *riskVerifyPendding) AddBeforFunc(befor func(context.Context) error) {
	if befor == nil {
		return
	}
	s.riskBeforFunc = append(s.riskBeforFunc, befor)
}

func (s *riskVerifyPendding) AddVerifier(verifier IVerifier) {
	s.verifier[verifier.VerifyKind()] = verifier
}
func (s *riskVerifyPendding) Verifier(verifyKind VerifyKind) IVerifier {
	if v, ok := s.verifier[verifyKind]; ok {
		return v
	}
	return nil
}

func (s *riskVerifyPendding) VerifierCode(code *model.VerifyCode) (VerifyKind, error) {
	for k, v := range s.verifier {
		if _, err := v.Verify(code); err != nil {
			return k, err
		}
	}
	return "", nil
}

// /
func (s *riskVerifyPendding) DoFunc(ctx context.Context) (VerifyKind, error) {
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
			v.SendCompletion()
		}
	}
	return "", nil
}

// /
func (s *riskVerifyPendding) DoBefor(ctx context.Context) (string, error) {
	for _, f := range s.riskBeforFunc {
		f(ctx)
	}
	return "", nil
}

func (s *riskVerifyPendding) DoAfter(ctx context.Context) (string, error) {
	for _, verifer := range s.verifier {
		if !verifer.IsDone() {
			return string(verifer.VerifyKind()), gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
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
func (s *riskVerifyPendding) AllDone() (VerifyKind, error) {
	for _, e := range s.verifier {
		if e.IsDone() {
			continue
		}
		return e.VerifyKind(), errRiskNotDone
	}
	return "", nil
}
