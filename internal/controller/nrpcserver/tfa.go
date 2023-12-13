package nats

import (
	"context"
	"riskcontral/api/risk/nrpc"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// func (s *NrpcServer) RpcTfaTx(ctx context.Context, req *tfav1.TfaTxReq) (*tfav1.TfaTxRes, error) {

// 	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
// 	if err != nil || tfaInfo == nil {
// 		g.Log().Warning(ctx, "TFATx:", "req:", req)
// 		g.Log().Errorf(ctx, "%+v", err)
// 		return nil, gerror.NewCode(mpccode.CodeInternalError)
// 	}
// 	kinds, err := service.TFA().TFATx(ctx, tfaInfo, req.RiskSerial)
// 	if err != nil {
// 		g.Log().Errorf(ctx, "%+v", err)
// 		return nil, gerror.NewCode(mpccode.CodePerformRiskError)
// 	}
// 	//

// 	return &tfav1.TfaTxRes{
// 		Kinds: kinds,
// 	}, nil
// }

func (*NrpcServer) RpcTfaInfo(ctx context.Context, req *nrpc.TfaInfoReq) (res *nrpc.TfaInfoRes, err error) {

	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendSmsCode")
	defer span.End()
	g.Log().Info(ctx, "RpcTfaInfo:", req)
	// info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	// if err != nil {
	// 	g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", info)
	// 	return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	// }
	if req.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", req)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil {
		g.Log().Warning(ctx, "TFAInfo no info:", "req:", req, "tfaInfo:", tfaInfo)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	if tfaInfo == nil {
		return nil, nil
	}
	res = &nrpc.TfaInfoRes{
		UserId: tfaInfo.UserId,
		Phone:  tfaInfo.Phone,
		UpPhoneTime: func() string {
			if tfaInfo.PhoneUpdatedAt == nil {
				return ""
			}

			return tfaInfo.PhoneUpdatedAt.String()
		}(),
		Mail: tfaInfo.Mail,
		UpMailTime: func() string {
			if tfaInfo.MailUpdatedAt == nil {
				return ""
			}
			return tfaInfo.MailUpdatedAt.String()
		}(),
	}
	g.Log().Info(ctx, "RpcTfaInfo:", res)
	return res, nil
}

// func (s *NrpcServer) RpcSendSmsCode(ctx context.Context, req *nrpc.SendPhoneCodeReq) (res *nrpc.SendMailCodeRes, err error) {
// 	//trace
// 	ctx, span := gtrace.NewSpan(ctx, "RpcSendSmsCode")
// 	defer span.End()
// 	//
// 	// info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	///
// 	risk := service.TFA().GetRiskVerify(ctx, req.UserId, req.RiskSerial)
// 	if risk == nil {
// 		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
// 	}
// 	v := risk.Verifier(model.VerifierKind_Phone)
// 	if v == nil {
// 		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
// 	}
// 	if err := s.limitSendPhone(ctx, req.UserId, v.Destination()); err != nil {
// 		g.Log().Errorf(ctx, "%+v", err)
// 		return nil, gerror.NewCode(mpccode.CodeLimitSendPhoneCode)
// 	}
// 	///
// 	_, err = service.TFA().SendPhoneCode(ctx, req.UserId, req.RiskSerial)
// 	if err != nil {
// 		g.Log().Errorf(ctx, "%+v", err)
// 	}
// 	return nil, err
// }

// func (s *NrpcServer) RpcSendMailCode(ctx context.Context, req *nrpc.SendMailCodeReq) (res *nrpc.SendMailCodeRes, err error) {
// 	//limit
// 	if err := s.counter(ctx, req.UserId, "SendMailCode"); err != nil {
// 		g.Log().Warningf(ctx, "%+v", err)
// 		return nil, err
// 	}
// 	if err := s.limitSendVerification(ctx, req.UserId, "SendMailCode"); err != nil {
// 		g.Log().Warningf(ctx, "%+v", err)
// 		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
// 	}
// 	//trace
// 	//trace
// 	ctx, span := gtrace.NewSpan(ctx, "RpcSendMailCode")
// 	defer span.End()
// 	//
// 	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
// 	if err != nil || tfaInfo == nil {
// 		g.Log().Warning(ctx, "SendMailCode:", req, err)
// 		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
// 	}
// 	risk := service.TFA().GetRiskVerify(ctx, req.UserId, req.RiskSerial)
// 	if risk == nil {
// 		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
// 	}
// 	v := risk.Verifier(model.VerifierKind_Mail)
// 	if v == nil {
// 		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
// 	}
// 	if err := s.limitSendMail(ctx, req.UserId, v.Destination()); err != nil {
// 		g.Log().Errorf(ctx, "%+v", err)
// 		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
// 	}
// 	///
// 	////
// 	switch risk.RiskKind {
// 	case model.RiskKind_BindMail, model.RiskKind_UpMail:
// 		if req.Mail == "" {
// 			return nil, gerror.NewCode(mpccode.CodeParamInvalid)
// 		}
// 		err = service.DB().TfaMailNotExists(ctx, req.Mail)
// 		if err != nil {
// 			g.Log().Warningf(ctx, "%+v", err)
// 			return nil, gerror.NewCode(mpccode.CodeTFAMailExists)
// 		}
// 		////
// 		service.TFA().TfaSetMail(ctx, tfaInfo, req.Mail, req.RiskSerial, risk.RiskKind)
// 		///
// 	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
// 	default:
// 		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
// 	}
// 	///
// 	_, err = service.TFA().SendMailCode(ctx, req.UserId, req.RiskSerial)
// 	if err != nil {
// 		if gerror.Cause(err).Error() == mpccode.CodeLimitSendMailCode.Error().Error() {
// 			g.Log().Warningf(ctx, "%+v", err)
// 			return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
// 		}
// 		g.Log().Warningf(ctx, "%+v", err)
// 		return nil, gerror.NewCode(mpccode.CodeTFASendMailFailed)
// 	}
// 	return nil, nil
// }

// func (*NrpcServer) RpcSendVerifyCode(ctx context.Context, req *tfav1.VerifyCodekReq) (res *tfav1.VerifyCodeRes, err error) {
// 	//trace
// 	ctx, span := gtrace.NewSpan(ctx, "RpcSendVerifyCode")
// 	defer span.End()
// 	//
// 	//
// 	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
// 	if err != nil {
// 		g.Log().Warningf(ctx, "%+v", err)
// 		return nil, mpccode.CodeTokenInvalid.Error()
// 	}
// 	code := &model.VerifyCode{
// 		MailCode:  req.MailCode,
// 		PhoneCode: req.PhoneCode,
// 	}
// 	err = service.TFA().VerifyCode(ctx, info.UserId, req.RiskSerial, code)
// 	if err != nil {
// 		g.Log().Errorf(ctx, "%+v", err)
// 		return nil, err
// 	}
// 	return nil, err
// }
