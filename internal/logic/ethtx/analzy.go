package ethtx

import (
	"context"
	"encoding/hex"
	analyzsigndata "riskcontral/common/ethtx/analyzSignData"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sEthTx) analzyTx(ctx context.Context, tx *analyzsigndata.SignTxData) (*model.AnalzyTxData, error) {

	target := tx.Target
	data := strings.TrimPrefix(tx.Data, "0x")
	///
	contractabi := ""
	if a, err := s.abicache.Get(ctx, target); !a.IsEmpty() {
		contractabi = a.String()
	} else {
		contractabi, err = service.DB().GetAbi(ctx, target)
		if err != nil {
			return nil, err
		}
		s.abicache.Set(ctx, target, contractabi, 0)
	}
	///
	contract, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return nil, err
	}
	//data
	dataByte, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	////
	method, err := contract.MethodById(dataByte[:4])
	if err != nil {
		g.Log().Warning(ctx, "", contract.Methods, err)
		return nil, err
	}
	args := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
	if err != nil {
		return nil, err
	}
	///
	atx := &model.AnalzyTxData{
		Target:     target,
		MethodId:   hex.EncodeToString(method.ID),
		MethodName: method.RawName,
		Sig:        method.Sig,
		Data:       data,
		Args:       args,
	}
	return atx, nil
}

func (s *sEthTx) AnalzyTxs(ctx context.Context, signtxs *analyzsigndata.SignTx) (*model.AnalzyTx, error) {
	// s.tidy(signtxs)
	atx := &model.AnalzyTx{}
	///
	for _, tx := range signtxs.Txs {
		adata, err := s.analzyTx(ctx, tx)
		if err != nil {
			return nil, err
		}
		atx.Txs = append(atx.Txs, adata)
	}
	return atx, nil
}
