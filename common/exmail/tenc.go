package exmail

import (
	"bytes"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
)

type TencMailClient struct {
	client                      *ses.Client
	verificationTemplateID      uint64
	bindingCompletionTemplateID uint64
	from                        string
	subject                     string
}

func NewTencMailClient(SecretId, SecretKey string, VerificationTemplateID, BindingCompletionTemplateID uint64, From string, Subject string) *TencMailClient {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取

	credential := common.NewCredential(SecretId, SecretKey)

	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ses.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ses.NewClient(credential, "ap-guangzhou", cpf)

	///
	return &TencMailClient{client: client,
		verificationTemplateID:      VerificationTemplateID,
		bindingCompletionTemplateID: BindingCompletionTemplateID,
		from:                        From, subject: Subject}
}
func (t *TencMailClient) SendBindingMail(destination string) (string, error) {
	return t.sendMail(destination, t.bindingCompletionTemplateID, "")
}
func (t *TencMailClient) SendMail(destination string, code string) (string, error) {
	return t.sendMail(destination, t.verificationTemplateID, code)
}

func (t *TencMailClient) sendMail(destination string, templateId uint64, code string) (string, error) {
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := ses.NewSendEmailRequest()

	request.FromEmailAddress = common.StringPtr(t.from)
	request.Destination = common.StringPtrs([]string{destination})
	//
	if code != "" {
		buf := bytes.Buffer{}
		buf.WriteString(`{"code":"`)
		buf.WriteString(code)
		buf.WriteString(`"}`)
		code = buf.String()
	}
	//
	request.Template = &ses.Template{
		TemplateID: common.Uint64Ptr(templateId),
		// TemplateData: common.StringPtr(`{"code":"123456"}`),
		TemplateData: common.StringPtr(code),
	}
	request.Subject = common.StringPtr(t.subject)

	// 返回的resp是一个SendEmailResponse的实例，与请求对象对应
	response, err := t.client.SendEmail(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", err
	}
	if err != nil {
		return "", err
	}
	// fmt.Printf("%s", response.ToJsonString())
	return response.ToJsonString(), nil
}
