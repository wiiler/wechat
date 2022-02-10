package core

import (
	"fmt"
	"io/ioutil"
	"wechat/model"
	"wechat/utils"
)

//  代金劵
//  CreateVoucher 创建优惠卷
func (c *Client) CreateVoucher(data *model.Stocks) (string, error) {

	header := make(map[string]string)
	// ContentType
	header[utils.ContentType] = utils.ApplicationJSON
	response, err := c.Post(VoucherStocksUrl, header, data)
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
	return string(body), nil
}

// StartVoucher 激活优惠卷
// 创建批次的商户号 mchid
// 批次号 stockid
func (c *Client) StartVoucher(stockid string) (string, error) {
	param := map[string]string{
		"stock_creator_mchid": c.MchID,
	}
	url := fmt.Sprintf(VoucherStartStocksUrl, stockid)
	response, err := c.Post(url, DefaultHeader, param)

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
	return string(body), nil
}

// PauseVoucher 暂停优惠卷
func (c *Client) PauseVoucher(stockid string) (string, error) {
	param := map[string]string{
		"stock_creator_mchid": c.MchID,
	}

	url := fmt.Sprintf(VoucherPauseStocksUrl, stockid)
	response, err := c.Post(url, DefaultHeader, param)

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
	return string(body), nil
}

// PauseVoucher 暂停优惠卷
func (c *Client) RestartVoucher(stockid string) (string, error) {
	param := map[string]string{
		"stock_creator_mchid": c.MchID,
	}

	url := fmt.Sprintf(VoucherRestartStocksUrl, stockid)

	response, err := c.Post(url, DefaultHeader, param)

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
	return string(body), nil
}

// GetVoucher发放优惠卷
func (c *Client) GetVoucher(openid string, merchantClient *Client, stockid string) (string, error) {

	url := fmt.Sprintf(VoucherGetStocksUrl, openid)
	header := make(map[string]string)
	header[utils.ContentType] = utils.ApplicationJSON

	param := map[string]string{
		"stock_creator_mchid": merchantClient.MchID,
		"stock_id":            stockid,
		"out_request_no":      utils.OrderNo(),
		"appid":               merchantClient.AppID,
	}
	response, err := merchantClient.Post(url, header, param)
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
	return string(body), nil
}
