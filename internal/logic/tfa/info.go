package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sTFA) TFAInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	if userId == "" {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	return service.DB().FetchTfaInfo(ctx, userId)

}

// func (s *sTFA) getTfaCache(ctx context.Context, userId string) (*entity.Tfa, error) {
// 	if v, ok := service.Cache().Get(ctx, userId+consts.KEY_TFAInfoCache); ok == nil {
// 		if v.IsEmpty() {
// 			return nil, gerror.NewCode(consts.CodeTFANotExist)
// 		}
// 		info := &entity.Tfa{}
// 		err := v.Struct(info)
// 		if err != nil {
// 			return nil, gerror.NewCode(consts.CodeInternalError)
// 		}
// 		return info, nil
// 	}
// 	return nil, gerror.NewCode(consts.CodeTFANotExist)
// }

// func (s *sTFA) setTfaCache(ctx context.Context, userId string, info *entity.Tfa) error {
// 	if info == nil || info.UserId == "" {
// 		return nil
// 	}
// 	return service.Cache().Set(ctx, userId+consts.KEY_TFAInfoCache, info, consts.SessionDur)
// }
