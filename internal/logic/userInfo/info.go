package userInfo

import (
	"context"
	"riskcontral/common"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcfg"
)

type sUserInfo struct {
	cache *gcache.Cache
	url   string
	///
	userGeter *common.UserTokenInfoGeter
}

func (s *sUserInfo) GetUserInfo(ctx context.Context, userToken string) (userInfo *common.UserInfo, err error) {

	g.Log().Debug(ctx, "GetUserInfo:", userToken)
	if userToken == "" {
		g.Log().Error(ctx, "GetUserInfo:", userToken)
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	//todo: check userToekn
	if err != nil {
		g.Log().Error(ctx, "GetUserInfo:", userToken)
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	///
	// 用户信息示例
	// "id": 10,
	// "appPubKey": "038c90b87d77f2cc3d26132e1ea26e14646d663e3f43f17180345df3d54b8b5c70",
	// "email": "sunwenhao0421@163.com",
	// "loginType": "tkey-auth0-twitter-cyan",
	// "address": "0xe73E35d8Ecc3972481138D01799ED3934cc57853",
	// "keyHash": "U2FsdGVkX1/O6j9czaWzdjjDo/XPjk1hI8pIoaxSuS52zIxVuStK/nS07ucgiM5si8NjN97rAux3aH7Ld2i5oO8UuL6tpNZmLMG9ZpwVTxvGkCa3H14vTxWNz+yBoWG8",
	// "create_time": 1691118876
	info, err := s.userGeter.GetUserInfo(ctx, userToken)
	return info, err
	// return &model.UserInfo{
	// 	Id: 10,
	// 	// AppPubKey:  "038c90b87d77f2cc3d26132e1ea26e14646d663e3f43f17180345df3d54b8b5c70",
	// 	AppPubKey:  utility.GenNewSid(),
	// 	Email:      "sunwenhao0421@163.com",
	// 	LoginType:  "tkey-auth0-twitter-cyan",
	// 	Address:    "0xe73E35d8Ecc3972481138D01799ED3934cc57853",
	// 	KeyHash:    "U2FsdGVkX1/O6j9czaWzdjjDo/XPjk1hI8pIoaxSuS52zIxVuStK/nS07ucgiM5si8NjN97rAux3aH7Ld2i5oO8UuL6tpNZmLMG9ZpwVTxvGkCa3H14vTxWNz+yBoWG8",
	// 	CreateTime: 1691118876,
	// }, nil
}

func new() *sUserInfo {
	///
	url, err := gcfg.Instance().Get(context.Background(), "userTokenUrl")
	if err != nil {
		panic(err)
	}
	///
	userGeter, err := common.NewUserInfoGeter(url.String())
	if err != nil {
		panic(err)
	}
	//
	s := &sUserInfo{
		userGeter: userGeter,
	}
	///

	return s
}

func init() {
	service.RegisterUserInfo(new())
}
