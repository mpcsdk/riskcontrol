package pendding

import (
	"context"
	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/conf"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/rand"
)

type RiskPendding struct {
	riskKind tfaconst.RISKKIND
	//风控序号
	riskSerial string
	//用户id
	UserId string
	///
	riskBeforFunc []func(context.Context) error
	// riskVerify    map[model.RiskKind]*riskVerify
	verifier   map[tfaconst.VERIFYKIND]tfaconst.IVerifier
	verifyKind []string
	// sender        map[VerifyKind]sender
	riskAfterFunc []func(context.Context) error
	///
	phoneSender int
	mailSender  int
	///
	dealline *gtime.Time
}

func NewRiskPendding(tfaInfo *entity.Tfa, kind tfaconst.RISKKIND, data *v1.RequestData) *RiskPendding {
	if tfaInfo == nil {
		return nil
	}
	riskSerial := rand.GenNewSid()
	risk := &RiskPendding{
		riskKind:   kind,
		UserId:     tfaInfo.UserId,
		riskSerial: riskSerial,
		verifier:   map[tfaconst.VERIFYKIND]tfaconst.IVerifier{},
		dealline:   gtime.Now().Add(penddingRiskeDuration),
		// dealline: gtime.Now(),
	}
	risk.build(tfaInfo, data)
	return risk
	// key := RiskPenddingKey(userId, riskSerial)
	// s.lock.Lock()
	// s.riskPendding[key] = risk
	// s.lock.Unlock()
	// g.Log().Info(s.ctx, "RiskPenddingContainer new:", key, risk.dealline.String())
	// }
	// return risk
}

// //

func (s *RiskPendding) RiskKind() tfaconst.RISKKIND {
	return s.riskKind
}
func (s *RiskPendding) VerifyKind() []string {
	return s.verifyKind
}
func (s *RiskPendding) RiskSerial() string {
	return s.riskSerial
}
func (s *RiskPendding) DealLine() *gtime.Time {
	return s.dealline
}
func (s *RiskPendding) GetVerifier(kind tfaconst.VERIFYKIND) tfaconst.IVerifier {
	if v, ok := s.verifier[kind]; ok {
		return v
	}
	return nil
}

func (s *RiskPendding) GetVerifiers() map[tfaconst.VERIFYKIND]tfaconst.IVerifier {
	return s.verifier
}
func (s *RiskPendding) AddAfterFunc(after func(context.Context) error) {
	if after == nil {
		return
	}
	s.riskAfterFunc = append(s.riskAfterFunc, after)
}
func (s *RiskPendding) AddBeforFunc(befor func(context.Context) error) {
	if befor == nil {
		return
	}
	s.riskBeforFunc = append(s.riskBeforFunc, befor)
}

func (s *RiskPendding) AddVerifier(verifier tfaconst.IVerifier) {
	s.verifier[verifier.VerifyKind()] = verifier
	s.verifyKind = make([]string, len(s.verifier))
	i := 0
	for k, _ := range s.verifier {
		s.verifyKind[i] = string(k)
		i++
	}
}

func (s *RiskPendding) VerifierCode(ctx context.Context, code *model.VerifyCode) (tfaconst.VERIFYKIND, error) {
	for k, v := range s.verifier {
		if _, err := v.Verify(ctx, code); err != nil {
			return k, err
		}
	}
	return "", nil
}

// /
func (s *RiskPendding) DoFunc(ctx context.Context) (tfaconst.VERIFYKIND, error) {
	if k, err := s.AllDone(); err != nil {
		return k, err
	} else {
		s.DoBefor(ctx)
		if _, err := s.DoAfter(ctx); err != nil {

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
func (s *RiskPendding) DoBefor(ctx context.Context) (string, error) {
	for _, f := range s.riskBeforFunc {
		f(ctx)
	}
	return "", nil
}

func (s *RiskPendding) DoAfter(ctx context.Context) (string, error) {
	for _, verifer := range s.verifier {
		if !verifer.IsDone() {
			return string(verifer.VerifyKind()), mpccode.CodeRiskVerifyCodeInvalid()
		}
	}
	//done
	for _, task := range s.riskAfterFunc {
		err := task(ctx)
		if err != nil {
			g.Log().Warning(ctx, "DoAfter:", "err:", err)
			return "", mpccode.CodeInternalError()
		}
	}
	return "", nil
}
func (s *RiskPendding) AllDone() (tfaconst.VERIFYKIND, error) {
	for _, e := range s.verifier {
		if e.IsDone() {
			continue
		}
		return e.VerifyKind(), tfaconst.ErrRiskNotDone
	}
	return "", nil
}

var penddingRiskeDuration time.Duration = time.Minute

func init() {
	penddingRiskeDuration = time.Duration(conf.Config.UserRisk.PenddingRiskeDuration) * time.Second
}
