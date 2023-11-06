package tfa

import (
	"context"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) TFAInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	if userId == "" {
		return nil, mpccode.ErrArg
	}
	info, err := service.DB().FetchTfaInfo(ctx, userId)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userId", userId),
		))
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	///
	return info, nil
}
