package conf

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type Cache struct {
	Duration int `json:"duration" v:"required|min:1000"`

	ApiInterval     int `json:"apiInterval" v:"required|min:1"`
	SessionDuration int `json:"sessionDuration" v:"required|min:1"`
	////
	VerificationCodeDuration int `json:"VerificationCodeDuration" v:"required|min:1"`
	LimitSendPhoneCount      int `json:"limitSendPhoneCount" v:"required|min:1"`
	LimitSendPhoneDuration   int `json:"limitSendPhoneDuration" v:"required|min:1"`
	LimitSendMailCount       int `json:"limitSendMailCount" v:"required|min:1"`
	LimitSendMailDuration    int `json:"limitSendMailDuration" v:"required|min:1"`
}

type UserRisk struct {
	UserControl              bool   `json:"userControl" v:"required"`
	TxControl                bool   `json:"txControl" v:"required"`
	PenddingRiskeDuration    int    `json:"penddingRiskeDuration" v:"required|min:1"`
	VerificationCodeInterval int    `json:"verificationCodeInterval" v:"required|min:1"`
	ForbiddenTime            string `json:"forbiddenTime" v:"required"`
}
type SmsHuawei struct {
	APIAddress        string `json:"apiAddress" v:"required"`
	ApplicationKey    string `json:"applicationKey" v:"required"`
	ApplicationSecret string `json:"applicationSecret" v:"required"`
	Sender            string `json:"sender" v:"required"`
	SenderCompletion  string `json:"senderCompletion" v:"required"`
	Signature         string `json:"signature" v:"required"`

	VerificationTemplateId        string `json:"verificationTemplateId" v:"required"`
	BindingVerificationTemplateId string `json:"bindingVerificationTemplateId" v:"required"`
	BindingCompletionTemplateId   string `json:"bindingCompletionTemplateId" v:"required"`
	UpVerificationTemplateId      string `json:"upVerificationTemplateId" v:"required"`
	UpCompletionTemplateId        string `json:"upCompletionTemplateId" v:"required"`
}
type SmsTenc struct {
	SecretId                      string `json:"secretId" v:"required"`
	SecretKey                     string `json:"secretKey" v:"required"`
	Endpoint                      string `json:"endpoint" v:"required"`
	SignMethod                    string `json:"signMethod" v:"required"`
	Region                        string `json:"region" v:"required"`
	SmsSdkAppId                   string `json:"smsSdkAppId" v:"required"`
	SignName                      string `json:"signName"`
	VerificationTemplateId        string `json:"verificationTemplateId" v:"required"`
	BindingVerificationTemplateId string `json:"bindingVerificationTemplateId" v:"required"`
	BindingCompletionTemplateId   string `json:"bindingCompletionTemplateId" v:"required"`
	UpVerificationTemplateId      string `json:"upVerificationTemplateId" v:"required"`
	UpCompletionTemplateId        string `json:"upCompletionTemplateId" v:"required"`
}
type SmsDomestic struct {
	Huawei *SmsHuawei `json:"huawei" v:"required"`
}
type SmsForeign struct {
	Tenc *SmsTenc `json:"tenc" v:"required"`
}
type Sms struct {
	Domestic *SmsDomestic `json:"domestic" v:"required"`
	Foreign  *SmsForeign  `json:"foreign" v:"required"`
}

type ExEmail struct {
	From                          string `json:"from" v:"required"`
	SecretId                      string `json:"secretId" v:"required"`
	SecretKey                     string `json:"secretKey" v:"required"`
	Subject                       string `json:"subject" v:"required"`
	VerificationTemplateId        int    `json:"verificationTemplateId" v:"required"`
	BindingVerificationTemplateId int    `json:"bindingVerificationTemplateId" v:"required"`
	BindingCompletionTemplateId   int    `json:"bindingCompletionTemplateId" v:"required"`
	UpVerificationTemplateId      int    `json:"upVerificationTemplateId" v:"required"`
	UpCompletionTemplateId        int    `json:"upCompletionTemplateId" v:"required"`
}

type Server struct {
	Address string `json:"address" v:"required"`
	WorkId  int    `json:"workId" v:"required|min:1"`
	Name    string `json:"name" v:"required"`
}

type NatsCfg struct {
	NatsUrl string `json:"natsUrl" v:"required"`
}
type Swap struct {
	ChainId  int64  `json:"chainId"`
	Contract string `json:"contract"`
}
type NftMarket struct {
	ChainId  int64  `json:"chainId"`
	Contract string `json:"contract"`
}

// //
type Cfg struct {
	Server       *Server   `json:"server" v:"required"`
	Cache        *Cache    `json:"cache" v:"required"`
	UserRisk     *UserRisk `json:"userRisk" v:"required"`
	Sms          *Sms      `json:"sms" v:"required"`
	ExEmail      *ExEmail  `json:"exEmail" v:"required"`
	UserTokenUrl string    `json:"userTokenUrl" v:"required"`
	JaegerUrl    string    `json:"jaegerUrl" v:"required"`
	Nats         *NatsCfg  `json:"nats" v:"required"`

	Swap      []*Swap      `json:"swap" v:"required"`
	NftMarket []*NftMarket `json:"nftMarket v:"required"`
}

var Config = &Cfg{}

func init() {
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()
	v, err := cfg.Data(ctx)
	if err != nil {
		panic(err)
	}
	val := gvar.New(v)
	err = val.Structs(Config)
	if err != nil {
		panic(err)
	}
	if err := g.Validator().Data(Config).Run(ctx); err != nil {
		panic(err)
	}
}
