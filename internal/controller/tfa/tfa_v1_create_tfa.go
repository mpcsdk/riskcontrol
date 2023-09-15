package tfa

// // @Summary 验证token，注册用户tfa
// func (c *ControllerV1) CreateTFA(ctx context.Context, req *v1.CreateTFAReq) (res *v1.CreateTFARes, err error) {
// 	///
// 	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
// 	if err != nil {
// 		return nil, gerror.NewCode(consts.CodeAuthFailed)
// 	}
// 	///
// 	_, err = dao.Tfa.Ctx(ctx).Insert(entity.Tfa{
// 		UserId:         userInfo.UserId,
// 		CreatedAt:      gtime.Now(),
// 		Mail:           req.Mail,
// 		Phone:          req.Phone,
// 		PhoneUpdatedAt: gtime.Now(),
// 		MailUpdatedAt:  gtime.Now(),
// 	})
// 	return nil, err
// }
