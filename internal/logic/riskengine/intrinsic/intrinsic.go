package intrinsic

import (
	"context"
	"fmt"

	"riskcontrol/internal/logic/riskengine/intrinsic/assetdb"
	"riskcontrol/internal/logic/riskengine/intrinsic/contract"
	"riskcontrol/internal/logic/riskengine/intrinsic/slice"
	"riskcontrol/internal/logic/riskengine/intrinsic/time"
	"riskcontrol/internal/logic/riskengine/intrinsic/util"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type riskcode struct {
	Pass   int32
	Verify int32

	Forbidden int32
	Error     int32
	NoCtrl    int32
}
type cfg struct {
	AllChainIds *slice.NumberList
}

func IntrinsicApis(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"print":    fmt.Println,
		"AggDB":    assetdb.NewIntrinsicAssetDB(ctx),
		"Contract": contract.NewIntrinsicContract(ctx),
		"Time":     time.NewIntrinsicTime(),
		"RiskCode": &riskcode{
			Pass:      mpccode.RiskCodePass,
			Verify:    mpccode.RiskCodeNeedVerification,
			Forbidden: mpccode.RiskCodeForbidden,
			Error:     mpccode.RiskCodeError,
		},
		///
		"NumberList": &slice.NumberList{},
		"BigInt":     &analzyer.BigInt{},
		"Util":       &util.Util{},
		"Cfg": &cfg{
			AllChainIds: &slice.NumberList{0},
		},
	}
}
