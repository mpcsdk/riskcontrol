package riskserver

import (
	"context"
	"fmt"
	"riskcontrol/internal/service"
	"testing"
)

func Test_ctrl(t *testing.T) {
	a, b := service.RiskCtrl().RiskCtrlTx(context.Background(), "123", nil, "txstr", "scene")
	fmt.Println(a, b)
}
