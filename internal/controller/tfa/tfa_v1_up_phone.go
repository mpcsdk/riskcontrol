package tfa

// import (
// 	"context"
// 	v1 "riskcontral/api/tfa/v1"
// 	"riskcontral/internal/consts"
// 	"riskcontral/internal/consts/conrisk"
// 	"riskcontral/internal/model/entity"
// 	"riskcontral/internal/service"

// 	"github.com/gogf/gf/v2/errors/gerror"
// 	"github.com/gogf/gf/v2/frame/g"
// 	"github.com/gogf/gf/v2/net/gtrace"
// )

// // @Summary 验证token，注册用户tfa
// func (c *ControllerV1) UpPhone(ctx context.Context, req *v1.UpPhoneReq) (res *v1.UpPhoneRes, err error) {
// 	//trace
// 	ctx, span := gtrace.NewSpan(ctx, "UpPhone")
// 	defer span.End()
// 	if err := c.counter(ctx, req.Token, "UpPhone"); err != nil {
// 		return nil, err
// 	}
// 	//
// 	///check token
// 	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
// 	if err != nil || userInfo == nil {
// 		g.Log().Warning(ctx, "UpPhone:", req, err)
// 		return nil, gerror.NewCode(consts.CodeTFANotExist)
// 	}
// 	///
// 	tfaInfo, err := service.TFA().TFAInfo(ctx, userInfo.UserId)
// 	if err != nil || tfaInfo == nil {
// 		g.Log().Warning(ctx, "UpMail:", req, err)
// 		return nil, gerror.NewCode(consts.CodeInternalError)
// 	}
// 	///check phone exists
// 	err = service.DB().TfaPhoneNotExists(ctx, req.Phone)
// 	if err != nil {
// 		g.Log().Warning(ctx, "UpPhone:", req, err)
// 		return nil, gerror.NewCode(consts.CodeTFAPhoneExists)
// 	}
// 	///upphoe riskcontrol
// 	//
// 	riskData := &conrisk.RiskTfa{
// 		UserId: tfaInfo.UserId,
// 		Kind:   consts.KEY_TFAKindUpPhone,
// 		Phone:  req.Phone,
// 	}
// 	riskSerial, code := service.Risk().RiskTFA(ctx, tfaInfo.UserId, riskData)
// 	if code == consts.RiskCodeForbidden {
// 		return nil, gerror.NewCode(consts.CodePerformRiskForbidden)
// 	}
// 	if code == consts.RiskCodeError {
// 		return nil, gerror.NewCode(consts.CodePerformRiskError)
// 		///
// 	}
// 	res = &v1.UpPhoneRes{}
// 	res.RiskSerial = riskSerial
// 	err = c.upPhone(ctx, tfaInfo, req.Phone, riskSerial)
// 	if err != nil {
// 		g.Log().Warning(ctx, "UpPhone:", err)
// 		return res, err
// 	}

// 	return res, gerror.NewCode(consts.CodePerformRiskNeedVerification)

// }

// func (c *ControllerV1) upPhone(ctx context.Context, tfainfo *entity.Tfa, phone string, riskSerial string) (err error) {

// 	///
// 	_, err = service.TFA().TFAUpPhone(ctx, tfainfo, phone, riskSerial)
// 	return err
// }
