package nats

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *NrpcServer) counter(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "counter"
	if v, err := s.cache.Get(ctx, key); err != nil || !v.IsEmpty() {
		return gerror.NewCode(mpccode.CodeApiLimit)
	} else {
		s.cache.Set(ctx, key, 1, apiInterval)
		return nil
	}
}
func (s *NrpcServer) limitSendVerification(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "limitSendVerification"
	if v, err := s.cache.Get(ctx, key); err != nil || !v.IsEmpty() {
		_, err = json.Marshal(func() {})
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key),
				mpccode.ErrDetail("method", method)),
		)
		return err
	} else {
		s.cache.Set(ctx, key, 1, limitSendInterval)
		return nil
	}
}

func (s *NrpcServer) delTimeOut(dts []*gtime.Time, limitDuration time.Duration) []*gtime.Time {
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

func (s *NrpcServer) limitSendPhone(ctx context.Context, tokenId string, phone string) error {
	key := phone + "limitSendPhone"
	sendtimes := []*gtime.Time{}
	if v, err := s.cache.Get(ctx, key); err != nil {
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
		)
		return err
	} else if !v.IsEmpty() {
		err := v.Structs(&sendtimes)
		if err != nil {
			return gerror.Wrap(err,
				mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
			)
		}
		////
		if len(sendtimes) >= limitSendPhoneDurationCnt {
			sendtimes = s.delTimeOut(sendtimes, limitSendPhoneDuration)
		}

		if len(sendtimes) >= limitSendPhoneDurationCnt {
			return mpccode.CodeLimitSendPhoneCode.Error()
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
func (s *NrpcServer) limitSendMail(ctx context.Context, tokenId string, mail string) error {
	key := mail + "limitSendMail"
	sendtimes := []*gtime.Time{}
	if v, err := s.cache.Get(ctx, key); err != nil {
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
		)
		return err
	} else if !v.IsEmpty() {
		err := v.Structs(&sendtimes)
		if err != nil {
			return gerror.Wrap(err,
				mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
			)
		}
		////
		if len(sendtimes) >= limitSendMailDurationCnt {
			sendtimes = s.delTimeOut(sendtimes, limitSendMailDuration)
		}

		if len(sendtimes) >= limitSendMailDurationCnt {
			return err
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
