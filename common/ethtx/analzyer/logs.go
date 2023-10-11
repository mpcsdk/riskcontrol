package analzyer

import (
	"context"
	"encoding/json"
	"math/big"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *Analzyer) AnalzyLogNFT(contract string, log *types.Log, nftrule *model.NftRule) (*entity.EthTx, error) {
	contract = strings.ToLower(contract)
	///
	if abicontract, ok := s.abiStructs[contract]; !ok {
		return nil, nil
	} else {
		if log.Topics[0].Hex() != nftrule.EventTopic {
			g.Log().Debug(context.Background(), "AnalzyLogNFT:", abicontract.Events, abicontract.Methods, log)
			return nil, nil
		}
		event, err := abicontract.EventByID(log.Topics[0])
		if err != nil {
			return nil, err
		}

		///
		///
		from := ""
		to := ""
		val := ""

		//
		for i, arg := range event.Inputs {
			if arg.Indexed {
				if arg.Name == nftrule.EventFromField {
					from = common.HexToAddress(log.Topics[i+1].Hex()).String()
				} else if arg.Name == nftrule.EventToField {
					to = common.HexToAddress(log.Topics[i+1].Hex()).String()
				} else if arg.Name == nftrule.EventTokenIdField {
					val = log.Topics[i+1].String()
				}
			}
		}

		datastr := common.Bytes2Hex(log.Data)
		topicstr, _ := json.Marshal(log.Topics)
		entity := &entity.EthTx{
			//todo: blocktime
			Name:        nftrule.Name,
			Kind:        "nft",
			BlockNumber: int64(log.BlockNumber),
			BlockHash:   log.BlockHash.String(),
			TxHash:      log.TxHash.String(),
			TxIndex:     int64(log.TxIndex),
			LogIndex:    int64(log.Index),

			Address:  from,
			Contract: contract,

			MethodName: nftrule.MethodName,
			MethodSig:  nftrule.MethodSig,
			EventName:  nftrule.EventName,
			EventSig:   nftrule.EventSig,

			Data:   datastr,
			Topics: string(topicstr),
			From:   from,
			To:     to,
			Value:  val,
		}

		return entity, nil
	}
}
func (s *Analzyer) AnalzyLogFT(contract string, log *types.Log, ftrule *model.FtRule) (*entity.EthTx, error) {
	contract = strings.ToLower(contract)
	///
	if abicontract, ok := s.abiStructs[contract]; !ok {
		return nil, nil
	} else {
		if log.Topics[0].Hex() != ftrule.EventTopic {
			return nil, nil
		}
		event, err := abicontract.EventByID(log.Topics[0])
		if err != nil {
			return nil, err
		}

		///
		///
		from := ""
		to := ""
		val := ""

		//
		for i, arg := range event.Inputs {
			if arg.Indexed {
				if arg.Name == ftrule.EventFromField {
					from = common.HexToAddress(log.Topics[i+1].Hex()).String()
				} else if arg.Name == ftrule.EventToField {
					to = common.HexToAddress(log.Topics[i+1].Hex()).String()
				}
			} else {
				args := make(map[string]interface{})
				event.Inputs.UnpackIntoMap(args, log.Data)
				if v, ok := args[ftrule.EventValueField]; ok {
					val = v.(*big.Int).String()
				}
			}
		}

		datastr := common.Bytes2Hex(log.Data)
		topicstr, _ := json.Marshal(log.Topics)
		entity := &entity.EthTx{
			//todo: blocktime
			Name:        ftrule.Name,
			Kind:        "ft",
			BlockNumber: int64(log.BlockNumber),
			BlockHash:   log.BlockHash.String(),
			TxHash:      log.TxHash.String(),
			TxIndex:     int64(log.TxIndex),
			LogIndex:    int64(log.Index),

			Address:  from,
			Contract: contract,

			MethodName: ftrule.MethodName,
			MethodSig:  ftrule.MethodSig,
			EventName:  ftrule.EventName,
			EventSig:   ftrule.EventSig,

			Data:   datastr,
			Topics: string(topicstr),
			From:   from,
			To:     to,
			Value:  val,
		}

		return entity, nil
	}
}
