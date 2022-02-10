package core

import (
	"context"
	"fmt"
	"time"
	"wechat/utils"
)

// 生成请求报文头中的 Authorization 信息
// ctx 上下文
// method  请求方法
// request.URL.RequestURI()  请求url
// signBody 签名的字符串
func (c *Client) GenerateAuthorizationHeader(ctx context.Context, method, canonicalURL, signBody string) (string, error) {
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf(SignatureMessageFormat, method, canonicalURL, timestamp, nonce, signBody)
	signatureResult, err := c.Sign(ctx, message)
	if err != nil {
		return "", err
	}
	authorization := fmt.Sprintf(
		HeaderAuthorizationFormat, c.getAuthorizationType(),
		c.MchID, nonce, timestamp, c.CertificateSerialNo, *signatureResult,
	)
	return authorization, nil
}

func (c *Client) getAuthorizationType() string {
	return "WECHATPAY2-" + c.Algorithm()
}
