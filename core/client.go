package core

import (
	"context"
	"crypto/rsa"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"runtime"
	"strings"
	"wechat/utils"
)

// Client 基础信息
type Client struct {
	AppID               string          // 商户appid应该为公众号的appid或者小程序的appid
	MchID               string          // 商户号
	CertificateSerialNo string          // 商户证书序列号
	PrivateKey          *rsa.PrivateKey // 商户私钥
	MchAPIv3Key         string          // 商户APIv3Key
	MchAPIv2Key         string          // 商户APIv2Key
	PublicKey           *rsa.PublicKey  // 平台公钥
	WeChatNo            string          // 平台证书序列号
	Ctx                 context.Context
}

// NewClient 新建一个Client
func NewClient(mchID, certificateSerialNo, privateKey, mchAPIv3Key, mchAPIv2Key, publicKey, weChatNo string) (c *Client, e error) {
	pk, err := utils.LoadPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	public, err := utils.LoadCertificate(publicKey)
	if err != nil {
		return nil, err
	}
	return &Client{
		MchID:               mchID,
		CertificateSerialNo: certificateSerialNo,
		PrivateKey:          pk,
		MchAPIv3Key:         mchAPIv3Key,
		MchAPIv2Key:         mchAPIv2Key,
		PublicKey:           public.PublicKey.(*rsa.PublicKey),
		WeChatNo:            weChatNo,
		Ctx:                 context.Background(),
	}, nil
}

// 设置请求头
func (c *Client) SetHeader(url, method string, requestBody interface{}, header map[string]string) (http.Header, error) {
	var (
		err           error
		authorization string
		request       *http.Request
	)
	request, err = http.NewRequestWithContext(c.Ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := utils.SetBody(requestBody, header[utils.ContentType])
	if err != nil {
		return nil, err
	}
	if authorization, err = c.GenerateAuthorizationHeader(c.Ctx, method, request.URL.RequestURI(), body.String()); err != nil {
		return nil, err
	}
	ua := fmt.Sprintf(UserAgentFormat, Version, runtime.GOOS, runtime.Version())

	request.Header.Add("Accept", "*/*")
	request.Header.Add("Authorization", authorization)
	request.Header.Add("User-Agent", ua)
	for i, h := range header {
		request.Header.Add(i, h)
	}
	return request.Header, nil
}

// Get 向微信支付发送一个 HTTP Get 请求
func (c *Client) Get(url string) (*http.Response, error) {
	headers, err := c.SetHeader(url, http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}
	return utils.Request(c.Ctx, http.MethodGet, url, headers, nil, "")
}

// Post 向微信支付发送一个 HTTP Post 请求
func (c *Client) Post(url string, header map[string]string, requestBody interface{}) (*http.Response, error) {
	headers, err := c.SetHeader(url, http.MethodPost, requestBody, header)
	if err != nil {
		return nil, err
	}
	body, err := utils.SetBody(requestBody, header[utils.ContentType])
	if err != nil {
		return nil, err
	}
	return utils.Request(c.Ctx, http.MethodPost, url, headers, body, body.String())
}

// / Patch 向微信支付发送一个 HTTP Patch 请求
func (c *Client) Patch(url string, requestBody interface{}) (*http.Response, error) {
	headers, err := c.SetHeader(url, http.MethodPatch, requestBody, nil)
	if err != nil {
		return nil, err
	}
	body, err := utils.SetBody(requestBody, headers.Get(utils.ContentType))
	if err != nil {
		return nil, err
	}
	return utils.Request(c.Ctx, http.MethodPatch, url, headers, body, body.String())
}

//	CreateFormField 设置form-data 中的普通属性
//	示例内容
//	Content-Disposition: form-data; name="meta";
//	Content-Type: application/json
//	{ "filename": "file_test.mp4", "sha256": " hjkahkjsjkfsjk78687dhjahdajhk " }
//  Upload 向微信支付发送上传图片或视频文件请求
func (c *Client) Upload(url, meta, requestBody string, header map[string]string) (*http.Response, error) {
	headers, err := c.SetHeader(url, http.MethodPost, meta, header)
	if err != nil {
		return nil, err
	}
	fmt.Println(headers, "headers")
	body, err := utils.SetBody(requestBody, header[utils.ContentType])
	if err != nil {
		return nil, err
	}
	return utils.Request(c.Ctx, http.MethodPost, url, headers, strings.NewReader(requestBody), body.String())
}

//	CreateFormField(w, "meta", "application/json", meta)
func (c *Client) CreateFormField(w *multipart.Writer, fieldName, contentType string, fieldValue []byte) error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s";`, fieldName))
	h.Set(utils.ContentType, contentType)
	part, err := w.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = part.Write(fieldValue)
	return err
}

//	CreateFormFile 设置form-data中的文件
//  示例内容：
//	Content-Disposition: form-data; name="file"; filename="file_test.mp4";
//	Content-Type: video/mp4
//	pic1  //pic1即为媒体视频的二进制内容
//	如果要设置上述内容，则CreateFormFile(w, "file_test.mp4", "video/mp4", pic1)
func (c *Client) CreateFormFile(w *multipart.Writer, filename, contentType string, file []byte) error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set(utils.ContentType, contentType)
	part, err := w.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = part.Write(file)
	return err
}
