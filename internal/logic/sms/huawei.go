package sms

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	uuid "github.com/satori/go.uuid"
)

type huawei struct {
	APIAddress        string
	ApplicationKey    string
	ApplicationSecret string
	Sender            string
	TemplateID        string
	Signature         string
}

func newhuawei() *huawei {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	return &huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.huawei.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.huawei.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.huawei.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.huawei.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.huawei.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.huawei.Signature").String(),
	}
}

// 无需修改,用于格式化鉴权头域,给"X-WSSE"参数赋值
const WSSE_HEADER_FORMAT = "UsernameToken Username=\"%s\",PasswordDigest=\"%s\",Nonce=\"%s\",Created=\"%s\""

// 无需修改,用于格式化鉴权头域,给"Authorization"参数赋值
const AUTH_HEADER_VALUE = "WSSE realm=\"SDP\",profile=\"UsernameToken\",type=\"Appkey\""

func (s *huawei) sendSms(receiver, code string) (resp string, statusCallBack string, err error) {
	apiAddress := s.APIAddress
	appKey := s.ApplicationKey
	appSecret := s.ApplicationSecret
	sender := s.Sender
	templateId := s.TemplateID
	signature := s.Signature

	//选填,短信状态报告接收地址,推荐使用域名,为空或者不填表示不接收状态报告
	// statusCallBack := ""

	/*
	 * 选填,使用无变量模板时请赋空值 string templateParas = "";
	 * 单变量模板示例:模板内容为"您的验证码是${1}"时,templateParas可填写为"[\"369751\"]"
	 * 双变量模板示例:模板内容为"您有${1}件快递请到${2}领取"时,templateParas可填写为"[\"3\",\"人民公园正门\"]"
	 * 模板中的每个变量都必须赋值，且取值不能为空
	 * 查看更多模板和变量规范:产品介绍>模板和变量规范
	 */
	templateParas := "[\"" + code + "\"]" //模板变量，此处以单变量验证码短信为例，请客户自行生成6位验证码，并定义为字符串类型，以杜绝首位0丢失的问题（例如：002569变成了2569）。

	body := buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Authorization"] = AUTH_HEADER_VALUE
	headers["X-WSSE"] = buildWsseHeader(appKey, appSecret)
	resp, err = post(apiAddress, []byte(body), headers)

	return resp, statusCallBack, err
}

/**
 * sender,receiver,templateId不能为空
 */
func buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature string) string {
	param := "from=" + url.QueryEscape(sender) + "&to=" + url.QueryEscape(receiver) + "&templateId=" + url.QueryEscape(templateId)
	if templateParas != "" {
		param += "&templateParas=" + url.QueryEscape(templateParas)
	}
	if statusCallBack != "" {
		param += "&statusCallback=" + url.QueryEscape(statusCallBack)
	}
	if signature != "" {
		param += "&signature=" + url.QueryEscape(signature)
	}
	return param
}

func post(url string, param []byte, headers map[string]string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func buildWsseHeader(appKey, appSecret string) string {
	var cTime = time.Now().Format("2006-01-02T15:04:05Z")
	var nonce = uuid.NewV4().String()
	nonce = strings.ReplaceAll(nonce, "-", "")

	h := sha256.New()
	h.Write([]byte(nonce + cTime + appSecret))
	passwordDigestBase64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf(WSSE_HEADER_FORMAT, appKey, passwordDigestBase64Str, nonce, cTime)
}
