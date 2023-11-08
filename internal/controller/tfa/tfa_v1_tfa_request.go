package tfa

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"

	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"
)

func (c *ControllerV1) TfaRequest(ctx context.Context, req *v1.TfaRequestReq) (res *v1.TfaRequestRes, err error) {
	// limit
	if err := c.counter(ctx, req.Token, "SendMailCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	/////
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	/////
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	}
	///
	tfaInfo, err := service.TFA().TFAInfo(ctx, info.UserId)
	if err != nil {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	///
	///
	var riskKind model.RiskKind = model.RiskKind_Nil
	//
	switch req.CodeType {
	case model.Type_TfaBindPhone:
		riskKind = model.RiskKind_BindPhone
		if tfaInfo != nil && tfaInfo.Phone != "" {
			return nil, gerror.NewCode(mpccode.CodeTFAExist)
		}
		if tfaInfo == nil {
			err = service.DB().InsertTfaInfo(ctx, info.UserId, &do.Tfa{
				UserId:    info.UserId,
				TokenData: info,
				CreatedAt: gtime.Now(),
			})
			if err != nil {
				err = gerror.Wrap(err, mpccode.ErrDetails(
					mpccode.ErrDetail("userId", info.UserId),
				))
				return nil, err
			}
		}
		///
	case model.Type_TfaBindMail:
		riskKind = model.RiskKind_BindMail
		if tfaInfo != nil && tfaInfo.Mail != "" {
			return nil, gerror.NewCode(mpccode.CodeTFAExist)
		}
		if tfaInfo == nil {
			err = service.DB().InsertTfaInfo(ctx, info.UserId, &do.Tfa{
				UserId:    info.UserId,
				TokenData: info,
				CreatedAt: gtime.Now(),
			})
			if err != nil {
				err = gerror.Wrap(err, mpccode.ErrDetails(
					mpccode.ErrDetail("userId", info.UserId),
				))
				return nil, err
			}
		}
		////
	case model.Type_TfaUpdatePhone:
		riskKind = model.RiskKind_UpPhone
		if tfaInfo == nil || tfaInfo.Phone == "" {
			return nil, gerror.NewCode(mpccode.CodeTFANotExist)
		}
	case model.Type_TfaUpdateMail:
		riskKind = model.RiskKind_UpMail
		if tfaInfo == nil || tfaInfo.Mail == "" {
			return nil, gerror.NewCode(mpccode.CodeTFANotExist)
		}
	default:
		return nil, gerror.NewCode(mpccode.CodeParamInvalid)
	}
	///
	tfaInfo, err = service.TFA().TFAInfo(ctx, info.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	}
	///
	riskSerial, code := service.Risk().RiskTFA(ctx, tfaInfo, &model.RiskTfa{
		UserId: tfaInfo.UserId,
		Type:   req.CodeType,
	})
	if code == mpccode.RiskCodeForbidden {
		return nil, gerror.NewCode(mpccode.CodePerformRiskForbidden)
	}
	if code == mpccode.RiskCodeError {
		return nil, gerror.NewCode(mpccode.CodePerformRiskError)
	}
	///
	vl, _ := service.TFA().TfaRiskTidy(ctx, tfaInfo, riskSerial, riskKind)
	res = &v1.TfaRequestRes{
		RiskSerial: riskSerial,
		VList:      vl,
	}
	return res, nil
	///
}
