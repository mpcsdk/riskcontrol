package tfa

// import (
// 	"context"
// 	pendding "riskcontrol/internal/logic/tfa/penddingrisk"
// 	"riskcontrol/internal/logic/tfa/tfaconst"
// 	"sync"
// 	"time"

// 	"github.com/gogf/gf/v2/frame/g"
// 	"github.com/gogf/gf/v2/os/gtime"
// 	"github.com/gogf/gf/v2/os/gtimer"
// 	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
// )

// type RiskPenddingContainer struct {
// 	lock                     sync.RWMutex
// 	riskPendding             map[RiskPenndingKey]*pendding.RiskPendding
// 	ctx                      context.Context
// 	verificationCodeDuration int
// 	////
// }

// func newRiskPenddingContainer(times int) *RiskPenddingContainer {
// 	s := &RiskPenddingContainer{
// 		riskPendding:             make(map[RiskPenndingKey]*pendding.RiskPendding),
// 		ctx:                      context.Background(),
// 		verificationCodeDuration: times,
// 	}
// 	//
// 	gtimer.Add(s.ctx, time.Second*time.Duration(times), func(ctx context.Context) {
// 		s.lock.Lock()
// 		defer s.lock.Unlock()
// 		n := gtime.Now()
// 		for key, risk := range s.riskPendding {
// 			if risk.DealLine().Before(n) {
// 				g.Log().Info(s.ctx, "RiskPenddingContainer dealline:", key, risk.DealLine().Local().String())
// 				delete(s.riskPendding, key)
// 			}
// 		}
// 	})
// 	//
// 	return s
// }

// // ///
// type RiskPenndingKey string

// func riskPenddingKey(userId, riskSerial string) RiskPenndingKey {
// 	return RiskPenndingKey("riskPendding:" + userId + ":" + riskSerial)
// }

// // //
// func (s *RiskPenddingContainer) getRiskVerifier(userId, riskSerial string) *pendding.RiskPendding {
// 	key := riskPenddingKey(userId, riskSerial)
// 	s.lock.RLock()
// 	defer s.lock.RUnlock()
// 	if risk, ok := s.riskPendding[key]; ok {
// 		return risk
// 	}
// 	return nil
// }

// func (s *RiskPenddingContainer) newRiskVerifier(
// 	tfaInfo *entity.Tfa,
// 	riskKind tfaconst.RISKKIND,
// ) *pendding.RiskPendding {
// 	risk := pendding.NewRiskPendding(tfaInfo, riskKind)
// 	key := riskPenddingKey(tfaInfo.UserId, risk.RiskSerial())
// 	s.lock.Lock()
// 	s.riskPendding[key] = risk
// 	s.lock.Unlock()
// 	g.Log().Info(s.ctx, "RiskPenddingContainer new:", key, risk.DealLine().String())
// 	return risk
// }
// func (s *RiskPenddingContainer) del(userId, riskSerial string) {
// 	key := riskPenddingKey(userId, riskSerial)
// 	delete(s.riskPendding, key)
// }
