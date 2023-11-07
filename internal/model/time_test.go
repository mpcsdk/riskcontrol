package model

import (
	"testing"

	"github.com/gogf/gf/v2/os/gtime"
)

func Test_gtimebefor(t *testing.T) {
	a := gtime.NewFromStr("2023-11-06 16:44:24")
	b := gtime.NewFromStr("2023-11-06 16:45:10")
	if !a.Before(b) {
		t.Error(a, b)
	}
	///
	b = gtime.NewFromStr("2023-11-06 16:41:42")
	if !a.Before(b) {
		t.Error(a, b)
	}
}
