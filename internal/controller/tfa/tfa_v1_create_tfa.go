package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// // @Summary 验证token，注册用户tfa
func (c *ControllerV1) CreateTFA(ctx context.Context, req *v1.CreateTFAReq) (res *v1.CreateTFARes, err error) {
	//	//trace
	ctx, span := gtrace.NewSpan(ctx, "CreateTFA:")
	defer span.End()
	g.Log().Debug(ctx, "crateTFA:", req)
	if err := c.counter(ctx, req.Token, "CreateTFA"); err != nil {
		return nil, err
	}
	///
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	///
	tfainfo, err := service.TFA().TFAInfo(ctx, info.UserId)
	if err != nil {
		return nil, err
	}
	if tfainfo != nil {
		return nil, gerror.NewCode(consts.CodeTFAExist)
	}
	///check phone or mail exists
	err = service.DB().TfaMailNotExists(ctx, req.Mail)
	if err != nil {
		g.Log().Warning(ctx, "crateTFA:", req, err)
		return nil, gerror.NewCode(consts.CodeTFAMailExists)
	}
	err = service.DB().TfaPhoneNotExists(ctx, req.Phone)
	if err != nil {
		g.Log().Warning(ctx, "crateTFA:", req, err)
		return nil, gerror.NewCode(consts.CodeTFAPhoneExists)
	}
	///
	// create nft
	riskData := &conrisk.RiskTfa{
		UserId: info.UserId,
		Kind:   consts.KEY_TFAKindCreate,
		Phone:  req.Phone,
		Mail:   req.Mail,
	}
	riskSerial, code := service.Risk().PerformRiskTFA(ctx, info.UserId, riskData)
	if code == consts.RiskCodeError {
		return nil, gerror.NewCode(consts.CodePerformRiskError)
	}
	////
	kind, err := service.TFA().TFACreate(ctx, info.UserId, req.Phone, req.Mail, riskSerial)
	if err != nil {
		g.Log().Error(ctx, "CreateTFA:")
		return nil, err
	}
	res = &v1.CreateTFARes{
		RiskSerial: riskSerial,
		RiskKind:   kind,
	}

	return res, err
}

func (c *ControllerV1) DialCode(ctx context.Context, req *v1.DialCodeReq) (res *v1.DialCodeRes, err error) {
	// 印度尼西亚,Indonesia,+62
	// 菲律宾,Philippines,+63
	// 泰国,Thailand,+66
	// 马来西亚,Malaysia,+60
	// 新加坡,Singapore,+65
	// 越南,Vietnam,+84
	// 缅甸,Myanmar,+95
	// 柬埔寨,Cambodia,+855
	// 老挝,Laos,+856
	// 文莱,Brunei,+673
	// 巴西,Brazil,+55
	// 阿根廷,Argentina,+54
	// 委内瑞拉,Venezuela,+58
	// 哥伦比亚,Colombia,+57
	// 秘鲁,Peru,+51
	// 智利,Chile,+56
	// 厄瓜多尔,Ecuador,+593
	// 玻利维亚,Bolivia,+591
	// 乌拉圭,Uruguay,+598
	// 巴拉圭,Paraguay,+595
	// 圭亚那,Guyana,+592
	// 苏里南,Suriname,+597
	// 美国,United States,+1
	// 印度,India,+91
	// 中国香港,Hong Kong, China,+852
	// 中国澳门,Macao, China,+853
	// 中国台湾,Taiwan, China,+886
	// 中国,China,+86
	return &v1.DialCodeRes{
		DialCodes: []*v1.DialCode{
			{Name: "中国", En: "China", DialCode: "+86"},
			{Name: "印度尼西亚", En: "Indonesia", DialCode: "+62"},
			{Name: "菲律宾", En: "Philippines", DialCode: "+63"},
			{Name: "泰国", En: "Thailand", DialCode: "+66"},
			{Name: "马来西亚", En: "Malaysia", DialCode: "+60"},
			{Name: "新新加坡", En: "Singapore", DialCode: "+65"},
			{Name: "越南", En: "Vietnam", DialCode: "+84"},
			{Name: "缅甸", En: "Myanmar", DialCode: "+95"},
			{Name: "柬埔寨", En: "Cambodia", DialCode: "+855"},
			{Name: "老挝", En: "Laos", DialCode: "+856"},
			{Name: "文莱", En: "Brunei", DialCode: "+673"},
			{Name: "巴西", En: "Brazil", DialCode: "+55"},
			{Name: "阿根廷", En: "Argentina", DialCode: "+54"},
			{Name: "委内瑞拉", En: "Venezuela", DialCode: "+58"},
			{Name: "哥伦比亚", En: "Colombia", DialCode: "+57"},
			{Name: "秘鲁", En: "Peru", DialCode: "+51"},
			{Name: "智利", En: "Chile", DialCode: "+56"},
			{Name: "厄瓜多尔", En: "Ecuador", DialCode: "+593"},
			{Name: "玻利维亚", En: "Bolivia", DialCode: "+591"},
			{Name: "乌拉圭", En: "Uruguay", DialCode: "+598"},
			{Name: "巴拉圭", En: "Paraguay", DialCode: "+595"},
			{Name: "圭亚那", En: "Guyana", DialCode: "+592"},
			{Name: "苏里南", En: "Suriname", DialCode: "+597"},
			{Name: "美国", En: "United States", DialCode: "+1"},
			{Name: "印度", En: "India", DialCode: "+91"},
			{Name: "中国香港", En: "Hong Kong", DialCode: "+852"},
			{Name: "中国澳门", En: "Macao", DialCode: "+853"},
			{Name: "中国台湾", En: "Taiwan", DialCode: "+886"},
		},
	}, nil
}
