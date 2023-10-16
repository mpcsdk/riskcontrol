package config

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type Cache struct {
	ApiInterval     int `json:"apiInterval" v:"required|min:1"`
	SessionDuration int `json:"sessionDuration" v:"required|min:1"`
}
type Etcd struct {
	Address string `json:"address" v:"required"`
}
type UserRisk struct {
	UserControl              bool   `json:"userControl" v:"required"`
	TxControl                bool   `json:"txControl" v:"required"`
	VerificationCodeDuration int    `json:"verificationCodeDuration" v:"required|min:1"`
	ForbiddenTime            string `json:"forbiddenTime" v:"required"`
}
type SmsHuawei struct {
	APIAddress        string `json:"apiAddress" v:"required"`
	ApplicationKey    string `json:"applicationKey" v:"required"`
	ApplicationSecret string `json:"applicationSecret" v:"required"`
	Sender            string `json:"sender" v:"required"`
	TemplateID        string `json:"templateID" v:"required"`
	Signature         string `json:"signature" v:"required"`
}
type SmsTenc struct {
	SecretId                    string `json:"secretId" v:"required"`
	SecretKey                   string `json:"secretKey" v:"required"`
	Endpoint                    string `json:"endpoint" v:"required"`
	SignMethod                  string `json:"signMethod" v:"required"`
	Region                      string `json:"region" v:"required"`
	SmsSdkAppId                 string `json:"smsSdkAppId" v:"required"`
	SignName                    string `json:"signName"`
	VerificationTemplateId      string `json:"verificationTemplateId" v:"required"`
	BindingCompletionTemplateId string `json:"bindingCompletionTemplateId" v:"required"`
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
	From                        string `json:"from" v:"required"`
	SecretId                    string `json:"secretId" v:"required"`
	SecretKey                   string `json:"secretKey" v:"required"`
	VerificationTemplateId      int    `json:"verificationTemplateId" v:"required"`
	BindingCompletionTemplateId int    `json:"bindingCompletionTemplateId" v:"required"`
	Subject                     string `json:"subject" v:"required"`
}

type Server struct {
	Address string `json:"address" v:"required"`
	WorkId  int    `json:"workId" v:"required|min:1"`
	Name    string `json:"name" v:"required"`
}

// //
type Cfg struct {
	Server       *Server   `json:"server" v:"required"`
	Cache        *Cache    `json:"cache" v:"required"`
	UserRisk     *UserRisk `json:"userRisk" v:"required"`
	Etcd         *Etcd     `json:"etcd" v:"required"`
	Sms          *Sms      `json:"sms" v:"required"`
	ExEmail      *ExEmail  `json:"exEmail" v:"required"`
	UserTokenUrl string    `json:"userTokenUrl" v:"required"`
	JaegerUrl    string    `json:"jaegerUrl" v:"required"`
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
