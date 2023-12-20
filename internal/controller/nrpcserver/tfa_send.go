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
	g.Log().Notice(ctx, "RpcSendMailCode:", "req:", req)
	//limit
	if err := s.apiLimit(ctx, req.UserId, "SendMailCode"); err != nil {
		return nil, err
	}
	// if err := s.limitSendMailDuration(ctx, req.UserId, "SendMailCode"); err != nil {
	// 	return nil, mpccode.CodeLimitSendMailCode()
	// }
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	///
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		return nil, mpccode.CodeTFANotExist()
	}
	////
	risk := service.TFA().GetRiskVerify(ctx, req.UserId, req.RiskSerial)
	if risk == nil {
		return nil, mpccode.CodeRiskSerialNotExist()
	}
	///
	////
	switch risk.RiskKind {
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
		if req.Mail == "" {
			return nil, mpccode.CodeParamInvalid()
		}
		notexists, err := service.DB().TfaMailNotExists(ctx, req.Mail)
		if err != nil || !notexists {
			return nil, mpccode.CodeTFAMailExists()
		}
		////
		service.TFA().TfaSetMail(ctx, tfaInfo, req.Mail, req.RiskSerial, risk.RiskKind)
		///
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
	default:
		return nil, mpccode.CodeRiskSerialNotExist()
	}

	v := risk.Verifier(model.VerifierKind_Mail)
	if v == nil {
		return nil, mpccode.CodeRiskSerialNotExist()
	}
	if err := s.limitSendMailCnt(ctx, req.UserId, v.Destination()); err != nil {
		return nil, mpccode.CodeLimitSendMailCode()
	}
	///
	_, err = service.TFA().SendMailCode(ctx, req.UserId, req.RiskSerial)

	if err != nil {
		if gerror.Cause(err).Error() == mpccode.CodeLimitSendMailCode().Error() {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, mpccode.CodeLimitSendMailCode()
		}
		g.Log().Warningf(ctx, "%+v", err)
		return nil, mpccode.CodeTFASendMailFailed()
	}
	return nil, nil
}

func (s *NrpcServer) RpcSendPhoneCode(ctx context.Context, req *nrpc.SendPhoneCodeReq) (res *nrpc.SendPhoneCodeRes, err error) {
	g.Log().Notice(ctx, "RpcSendPhoneCode:", "req:", req)

	//limit
	if err := s.apiLimit(ctx, req.UserId, "SendSmsCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	// if err := s.limitSendVerification(ctx, req.UserId, "SendSmsCode"); err != nil {
	// 	g.Log().Errorf(ctx, "%+v", err)
	// 	return nil, mpccode.CodeLimitSendMailCode()
	// }

	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	///
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		return nil, mpccode.CodeTFANotExist()
	}
	///
	risk := service.TFA().GetRiskVerify(ctx, req.UserId, req.RiskSerial)
	if risk == nil {
		return nil, mpccode.CodeRiskSerialNotExist()
	}
	///
	////
	switch risk.RiskKind {
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
		if req.Phone == "" {
			return nil, mpccode.CodeParamInvalid()
		}
		notexists, err := service.DB().TfaPhoneNotExists(ctx, req.Phone)
		if err != nil || !notexists {
			return nil, mpccode.CodeTFAPhoneExists()
		}
		////
		service.TFA().TfaSetPhone(ctx, tfaInfo, req.Phone, req.RiskSerial, risk.RiskKind)
		///
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
	default:
		return nil, mpccode.CodeRiskSerialNotExist()
	}

	v := risk.Verifier(model.VerifierKind_Phone)
	if v == nil {
		return nil, mpccode.CodeRiskSerialNotExist()
	}
	if err := s.limitSendPhoneCnt(ctx, req.UserId, v.Destination()); err != nil {
		return nil, mpccode.CodeLimitSendPhoneCode()
	}
	///
	_, err = service.TFA().SendPhoneCode(ctx, req.UserId, req.RiskSerial)
	if err != nil {
		return nil, mpccode.CodeTFASendSmsFailed()
	}
	if err != nil {
		if gerror.Cause(err).Error() == mpccode.CodeLimitSendPhoneCode().Error() {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, mpccode.CodeLimitSendPhoneCode()
		}
		g.Log().Warningf(ctx, "%+v", err)
		return nil, mpccode.CodeTFASendSmsFailed()
	}
	return nil, nil
}
