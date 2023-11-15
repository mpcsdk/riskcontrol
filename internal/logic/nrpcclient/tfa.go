package nrpcclient

import (
	"context"
	v1 "riskcontral/api/tfa/nrpc/v1"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *sNrpcClient) RpcTfaTx(ctx context.Context, userId string, riskSerial string) ([]string, error) {
	rst, err := s.cli.RpcTfaTx(&v1.TfaTxReq{
		UserId:     userId,
		RiskSerial: riskSerial,
	})

	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcRiskTFA TimeOut:")
			s.Flush()
			return nil, nil
		}
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("useid", userId),
		))
		return nil, err
	}
	///
	return rst.Kinds, nil
}
func (s *sNrpcClient) RpcTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error) {

	rst, err := s.cli.RpcTfaInfo(&v1.TFAReq{
		UserId: userId,
	})

	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcRiskTFA TimeOut:")
			s.Flush()
			return nil, nil
		}
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("useid", userId),
		))
		return nil, err
	}
	///
	return &entity.Tfa{
		UserId: rst.UserId,
		Mail:   rst.Mail,
		Phone:  rst.Phone,
		PhoneUpdatedAt: func() *gtime.Time {
			if rst.UpPhoneTime == "" {
				return nil
			}
			return gtime.NewFromStr(rst.UpPhoneTime)
		}(),
		MailUpdatedAt: func() *gtime.Time {
			if rst.UpMailTime == "" {
				return nil
			}
			return gtime.NewFromStr(rst.UpMailTime)
		}(),
	}, nil
}
func (s *sNrpcClient) RpcAlive(ctx context.Context) error {

	_, err := s.cli.RpcAlive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcAlive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
