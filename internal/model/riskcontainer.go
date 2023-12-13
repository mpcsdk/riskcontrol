package model

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

var errRiskNotDone error = errors.New("risk not done")

const (
	VerifierKind_Nil   = "nil"
	VerifierKind_Phone = "Phone"
	VerifierKind_Mail  = "Mail"
)

type RiskPenddingContainer struct {
	lock                     sync.RWMutex
	riskPendding             map[RiskPenndingKey]*RiskVerifyPendding
	ctx                      context.Context
	verificationCodeDuration int
}

func NewRiskPenddingContainer(times int) *RiskPenddingContainer {
	s := &RiskPenddingContainer{
		riskPendding:             make(map[RiskPenndingKey]*RiskVerifyPendding),
		ctx:                      context.Background(),
		verificationCodeDuration: times,
	}
	//
	gtimer.Add(s.ctx, time.Second*time.Duration(times), func(ctx context.Context) {
		s.lock.Lock()
		defer s.lock.Unlock()
		n := gtime.Now()
		for key, risk := range s.riskPendding {
			if risk.dealline.Before(n) {
				g.Log().Info(s.ctx, "RiskPenddingContainer dealline:", key, risk.dealline.Local().String())
				delete(s.riskPendding, key)
			}
		}
	})
	//
	return s
}

// /
// //
func (s *RiskPenddingContainer) GetRiskVerify(userId, riskSerial string) *RiskVerifyPendding {
	key := RiskPenddingKey(userId, riskSerial)
	s.lock.RLock()
	defer s.lock.RUnlock()
	if risk, ok := s.riskPendding[key]; ok {
		return risk
	}
	return nil
}

func (s *RiskPenddingContainer) NewRiskPendding(
	userId, riskSerial string,
	riskKind RiskKind,
) *RiskVerifyPendding {
	risk := s.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		risk = &RiskVerifyPendding{
			RiskKind:   riskKind,
			UserId:     userId,
			RiskSerial: riskSerial,
			verifier:   map[VerifyKind]IVerifier{},
			dealline:   gtime.Now().Add(time.Duration(s.verificationCodeDuration) * time.Second),
			// dealline: gtime.Now(),
		}
		key := RiskPenddingKey(userId, riskSerial)
		s.lock.Lock()
		s.riskPendding[key] = risk
		s.lock.Unlock()
		g.Log().Info(s.ctx, "RiskPenddingContainer new:", key, risk.dealline.String())
	}
	return risk
}
func (s *RiskPenddingContainer) Del(userId, riskSerial string) {
	key := RiskPenddingKey(userId, riskSerial)
	delete(s.riskPendding, key)
}

// var errRiskNotExist error = errors.New("risk not exist")
