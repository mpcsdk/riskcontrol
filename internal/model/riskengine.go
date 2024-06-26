package model

import (
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type RiskContext struct {
	ChainId uint64
	// SceneNo string
	CurTfa *entity.Tfa
}
type RiskExecData struct {
	SignTxs []*analzyer.AnalzyedSignTx
	Context *RiskContext
}
