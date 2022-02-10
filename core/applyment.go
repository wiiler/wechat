package core

// 商户进件
// Applyment 提交申请
// QueryApplyment 查询申请单状态 id 为微信返回id
// UpdateMchBank 修改结算账号
// QueryMchBank 查询结算账户

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"wechat/model"
	"wechat/utils"
)

// 提交申请
func (c *Client) Applyment(d *model.ApplymentData) (int, error) {
	header := make(map[string]string)
	header[utils.ContentType] = utils.ApplicationJSON
	header["Wechatpay-Serial"] = c.WeChatNo
	response, err := c.Post(ApplymentUrl, header, d)
	if err != nil {
		return 0, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	if response.Body != nil {
		defer response.Body.Close()
	}
	res := new(model.ApplymentRes)
	json.Unmarshal(body, res)
	return res.ApplymentID, nil
}

// 查询申请单状态 id 为微信返回id
func (c *Client) QueryApplyment(id string) (*model.ApplymentRes, error) {
	url := fmt.Sprintf(QueryApplymentUrl, id)

	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	if response.Body != nil {
		defer response.Body.Close()
	}
	res := new(model.ApplymentRes)
	json.Unmarshal(body, res)
	return res, nil
}

// 修改结算账号
func (c *Client) UpdateMchBank(mchId string, d *model.UpdateMchBankData) (bool, error) {
	url := fmt.Sprintf(UpdateMchBankUrl, mchId)
	response, err := c.Post(url, DefaultHeader, d)
	if err != nil {
		return false, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	if response.StatusCode == 204 {
		return true, nil
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		return false, errors.New(string(body))
	}

}

// 查询结算账户 mchId 为商户号
func (c *Client) QueryMchBank(mchId string) (*model.ApplymentBankRes, error) {
	url := fmt.Sprintf(QueryMchBankUrl, mchId)

	response, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	if response.Body != nil {
		defer response.Body.Close()
	}
	res := new(model.ApplymentBankRes)
	json.Unmarshal(body, res)
	return res, nil
}
