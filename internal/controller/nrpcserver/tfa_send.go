package nats

import (
	"context"
	"riskcontral/api/risk/nrpc"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *NrpcServer) RpcSendMailCode(ctx context.Context, req *nrpc.SendMailCodeReq) (res *nrpc.SendMailCodeRes, err error) {
	//limit
	if err := s.counter(ctx, req.UserId, "SendMailCode"); err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		return nil, err
	}
	if err := s.limitSendVerification(ctx, req.UserId, "SendMailCode"); err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
	}
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	///
	//
	// info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	// if err != nil || info == nil {
	// 	g.Log().Warningf(ctx, "%+v", err)
	// 	return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	// }
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "SendMailCode:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	////
	risk := service.TFA().GetRiskVerify(ctx, req.UserId, req.RiskSerial)
	if risk == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	///
	////
	switch risk.RiskKind {
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
		if req.Mail == "" {
			return nil, gerror.NewCode(mpccode.CodeParamInvalid)
		}
		err = service.DB().TfaMailNotExists(ctx, req.Mail)
		if err != nil {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeTFAMailExists)
		}
		////
		service.TFA().TfaSetMail(ctx, tfaInfo, req.Mail, req.RiskSerial, risk.RiskKind)
		///
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
	default:
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	/// limit send cnt
	// risk := service.TFA().GetRiskVerify(ctx, info.UserId, req.RiskSerial)
	// if risk == nil {
	// 	return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	// }
	v := risk.Verifier(model.VerifierKind_Mail)
	if v == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	if err := s.limitSendMail(ctx, req.UserId, v.Destination()); err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
	}
	///
	_, err = service.TFA().SendMailCode(ctx, req.UserId, req.RiskSerial)

	if err != nil {
		if gerror.Cause(err).Error() == mpccode.CodeLimitSendMailCode.Error().Error() {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
		}
		g.Log().Warningf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFASendMailFailed)
	}
	return nil, nil
}

func (s *NrpcServer) RpcSendPhoneCode(ctx context.Context, req *nrpc.SendPhoneCodeReq) (res *nrpc.SendPhoneCodeRes, err error) {

	//limit
	if err := s.counter(ctx, req.UserId, "SendSmsCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	if err := s.limitSendVerification(ctx, req.UserId, "SendSmsCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
	}

	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	///
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "SendSmsCode:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	///
	risk := service.TFA().GetRiskVerify(ctx, req.UserId, req.RiskSerial)
	if risk == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	///
	////
	switch risk.RiskKind {
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
		if req.Phone == "" {
			return nil, gerror.NewCode(mpccode.CodeParamInvalid)
		}
		err = service.DB().TfaPhoneNotExists(ctx, req.Phone)
		if err != nil {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeTFAPhoneExists)
		}
		////
		service.TFA().TfaSetPhone(ctx, tfaInfo, req.Phone, req.RiskSerial, risk.RiskKind)
		///
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
	default:
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}

	v := risk.Verifier(model.VerifierKind_Phone)
	if v == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	if err := s.limitSendPhone(ctx, req.UserId, v.Destination()); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendPhoneCode)
	}
	///
	_, err = service.TFA().SendPhoneCode(ctx, req.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFASendSmsFailed)
	}
	if err != nil {
		if gerror.Cause(err).Error() == mpccode.CodeLimitSendPhoneCode.Error().Error() {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeLimitSendPhoneCode)
		}
		g.Log().Warningf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFASendSmsFailed)
	}
	return nil, nil
}
