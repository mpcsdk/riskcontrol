package contract

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mpcsdk/mpcCommon/mpcdao/dao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func NewIntrinsicContract(ctx context.Context) *IntrinsicContract {
	return &IntrinsicContract{ctx: ctx}
}

type IntrinsicContract struct {
	ctx context.Context
}

func (s *IntrinsicContract) IsNft(contractAddress string, chainId int64) bool {
	g.Log().Debug(s.ctx, "IsNFt:", contractAddress, chainId)
	if contractAddress == "" {
		return false
	}
	where := dao.Contractabi.Ctx(s.ctx).
		Where(dao.Contractabi.Columns().ContractAddress, contractAddress).
		Where(dao.Contractabi.Columns().ChainId, gconv.String(chainId)).
		Where(dao.Contractabi.Columns().ContractKind, []string{"erc721", "erc1155"})
	cnt, err := where.Count()
	if err != nil {
		g.Log().Warning(s.ctx, "IsNft", contractAddress, err)
		return false
	}
	g.Log().Debug(s.ctx, "IsNFt:", contractAddress, chainId, cnt)
	if cnt > 0 {
		return true
	}
	return false
}
func (s *IntrinsicContract) IsFt(contractAddress string, chainId int64) bool {
	g.Log().Debug(s.ctx, "IsFt:", contractAddress, gconv.String(chainId))
	if contractAddress == "" {
		return true
	}
	cnt, err := dao.Contractabi.Ctx(s.ctx).
		Where(dao.Contractabi.Columns().ContractAddress, contractAddress).
		Where(dao.Contractabi.Columns().ChainId, gconv.String(chainId)).
		Where(dao.Contractabi.Columns().ContractKind, "erc20").
		Count()
	if err != nil {
		g.Log().Warning(s.ctx, "IsFt", contractAddress, chainId, err)
		return false
	}
	g.Log().Debug(s.ctx, "IsFt:", contractAddress, chainId, cnt)
	if cnt > 0 {
		return true
	}
	return false
}

func (s *IntrinsicContract) ContractRule(contract string) *entity.Contractrule {
	rst, err := dao.Contractrule.Ctx(s.ctx).Where(do.Contractrule{
		ContractAddress: contract,
	}).One()
	if err != nil {
		return nil
	}
	///
	var data *entity.Contractrule
	err = rst.Struct(&data)
	if err != nil {
		return nil
	}
	if data == nil {
		return nil
	}
	return data
}

func (s *IntrinsicContract) RiskRule(contract string) *entity.Contractrule {
	return s.ContractRule(contract)

}
func (s *IntrinsicContract) ContractName(contract string) string {
	rule := s.RiskRule(contract)
	if rule == nil {
		return "RPG"
	}
	return rule.ContractName
}
