package time

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type IntrinsicTime struct {
	Hour int64
	Day  int64
}

func (t *IntrinsicTime) TimeStampBefore(x int64) int64 {
	return gtime.Now().Add(-time.Duration(x)).Timestamp()
}
func (t *IntrinsicTime) Now() *gtime.Time {
	return gtime.Now()
}
func (t *IntrinsicTime) NowAdd(d int64) *gtime.Time {
	return gtime.Now().Add(time.Duration(d))
}
func (t *IntrinsicTime) After(a, b *gtime.Time) bool {
	return a.After(b)
}

func NewIntrinsicTime() *IntrinsicTime {
	return &IntrinsicTime{
		Hour: int64(gtime.H),
		Day:  int64(gtime.D),
	}
}
