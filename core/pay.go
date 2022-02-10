package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"wechat/model"
	"wechat/utils"
)

type Response struct {
	PrepayID string `json:"prepay_id"`
}

type ResponseUrl struct {
	CodeUrl string `json:"code_url"`
}

//	JsApiPay、小程序下单
func (c *Client) JsApiPay(d *model.JsApiPay) (*model.JsApiPayRes, error) {
	response, err := c.Post(JsapiPayUrl, DefaultHeader, d)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := new(Response)
	json.Unmarshal(body, &res)
	if len(res.PrepayID) == 0 {
		return nil, errors.New("失败")
	}
	nonce, _ := utils.GenerateNonce()
	jsRes := &model.JsApiPayRes{
		AppId:     c.AppID,
		TimeStamp: utils.TimeStamp(),
		NonceStr:  nonce,
		Package:   res.PrepayID,
		SignType:  "RSA",
		PaySign:   "",
	}
	msg := fmt.Sprintf("%s\n%s\n%s\n%s\n", jsRes.AppId, jsRes.TimeStamp, jsRes.NonceStr, jsRes.Package)
	jsRes.PaySign, _ = utils.SignSHA256WithRSA(msg, c.PrivateKey)
	return jsRes, nil
}

// App 下单
func (c *Client) AppPay(d *model.AppNativePay) (*model.AppPayRes, error) {
	response, err := c.Post(AppPayUrl, DefaultHeader, d)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := new(Response)
	json.Unmarshal(body, res)
	if len(res.PrepayID) == 0 {
		return nil, errors.New("失败")
	}
	nonce, _ := utils.GenerateNonce()
	jsRes := &model.AppPayRes{
		AppId:     c.AppID,
		PartnerID: c.MchID,
		PrepayID:  res.PrepayID,
		Package:   "Sign=WXPay",
		NonceStr:  nonce,
		TimeStamp: utils.TimeStamp(),
		PaySign:   "",
	}
	msg := fmt.Sprintf("%s\n%s\n%s\n%s\n", jsRes.AppId, jsRes.TimeStamp, jsRes.NonceStr, jsRes.PrepayID)
	jsRes.PaySign, _ = utils.SignSHA256WithRSA(msg, c.PrivateKey)
	return jsRes, nil
}

// Native 下单
func (c *Client) NativePay(d *model.AppNativePay) (string, error) {
	response, err := c.Post(NativePayUrl, DefaultHeader, d)
	if err != nil {
		return "", err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	res := new(ResponseUrl)
	json.Unmarshal(body, res)
	if len(res.CodeUrl) == 0 {
		return "", errors.New("失败")
	}

	return res.CodeUrl, nil
}

// 微信支付订单号查询
func (c *Client) QueryTransactionID(transactionid string, submchid string) (*model.AppNativePay, error) {
	url := fmt.Sprintf(QueryTransactionIDUrl, transactionid, c.MchID, submchid)

	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, _ := ioutil.ReadAll(response.Body)
	res := new(model.AppNativePay)
	json.Unmarshal(body, res)
	return res, nil
}

// 商户订单号查询
func (c *Client) QueryOutTradeNo(outtradeno string, submchid string) (*model.AppNativePay, error) {
	url := fmt.Sprintf(QueryOutTradeNoUrl, outtradeno, c.MchID, submchid)

	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, _ := ioutil.ReadAll(response.Body)
	res := new(model.AppNativePay)
	json.Unmarshal(body, res)
	return res, nil
}

func (c *Client) CloseOutTradeNo(outtradeno, submchid string) (bool, error) {
	url := fmt.Sprintf(CloseOutTradeNoUrl, outtradeno)
	d := map[string]string{
		"sp_mchid":  c.MchID,
		"sub_mchid": submchid,
	}
	response, err := c.Post(url, DefaultHeader, d)
	if err != nil {
		return false, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	if len(string(body)) == 0 && response.StatusCode == 204 {
		return true, nil
	}
	return false, errors.New("关闭失败")
}

//  退款申请
func (c *Client) Refund(d *model.Refund) (*model.RefundRes, error) {
	response, err := c.Post(RefundUrl, DefaultHeader, d)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	res := new(model.RefundRes)
	json.Unmarshal(body, res)
	return res, nil
}

//  查询退款
func (c *Client) QueryRefund(out_refund_no, sub_mchid string) (*model.RefundRes, error) {

	url := fmt.Sprintf(QueryRefundUrl, out_refund_no, sub_mchid)

	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	res := new(model.RefundRes)
	json.Unmarshal(body, res)
	return res, nil
}
