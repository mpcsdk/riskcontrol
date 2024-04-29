package tfa

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// var limitSendInterval = time.Second * 60
var limitSendPhoneDurationCnt = 50
var limitSendPhoneDuration = time.Minute
var limitSendMailDurationCnt = 10
var limitSendMailDuration = time.Minute

func (s *sTFA) delTimeOut(dts []*gtime.Time, limitDuration time.Duration) []*gtime.Time {
	i := 0
	n := gtime.Now()

	beforTime := n.Add(-limitDuration)
	for _, st := range dts {
		if st.After(beforTime) {
			dts[i] = st
			i++
		}
	}
	return dts[:i]
}

func (s *sTFA) limitSendPhoneCnt(ctx context.Context, tokenId string, phone string) error {
	key := phone + "limitSendPhone"
	sendtimes := []*gtime.Time{}
	if v, err := s.cache.Get(ctx, key); err != nil {
		g.Log().Warning(ctx, "limitSendPhone:", "tokenId:", tokenId, "phone", phone, "err", err)
		return mpccode.CodeLimitSendMailCode()
	} else if !v.IsEmpty() {
		err := v.Structs(&sendtimes)
		if err != nil {
			g.Log().Warning(ctx, "limitSendPhone:", "tokenId:", tokenId, "phone", phone, "err", err)
			return mpccode.CodeLimitSendMailCode()
		}
		////
		if len(sendtimes) >= limitSendPhoneDurationCnt {
			sendtimes = s.delTimeOut(sendtimes, limitSendPhoneDuration)
		}

		if len(sendtimes) >= limitSendPhoneDurationCnt {
			g.Log().Info(ctx, "limitSendPhone:", "tokenId:", tokenId, "phone", phone)
			return mpccode.CodeLimitSendPhoneCode()
		}
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	} else {
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	}
}
func (s *sTFA) limitSendMailCnt(ctx context.Context, tokenId string, mail string) error {
	key := mail + "limitSendMail"
	sendtimes := []*gtime.Time{}
	if v, err := s.cache.Get(ctx, key); err != nil {
		g.Log().Warning(ctx, "limitSendMailCnt:", "tokenId:", tokenId, "mail", mail, "err", err)
		return mpccode.CodeLimitSendMailCode()
	} else if !v.IsEmpty() {
		err := v.Structs(&sendtimes)
		if err != nil {
			g.Log().Warning(ctx, "limitSendMailCnt:", "tokenId:", tokenId, "mail", mail, "err", err)
			return mpccode.CodeLimitSendMailCode()
		}
		////
		if len(sendtimes) >= limitSendMailDurationCnt {
			sendtimes = s.delTimeOut(sendtimes, limitSendMailDuration)
		}

		if len(sendtimes) >= limitSendMailDurationCnt {
			g.Log().Info(ctx, "limitSendMailCnt:", "tokenId:", tokenId, "mail", mail)
			return mpccode.CodeLimitSendPhoneCode()
		}
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	} else {
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	}
}
