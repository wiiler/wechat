package core

import (
	"encoding/json"
	"io/ioutil"
	"wechat/model"
	"wechat/utils"
)

// GetCert 获取微信平台证书
func (c *Client) GetCert() ([]*model.ItemCerts, error) {

	res, err := c.Get(GetCertsUrl)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)
	certs := new(model.CertsRes)
	json.Unmarshal(body, certs)
	for i, v := range certs.Data {
		c, e := utils.DecryptAES256GCM(c.MchAPIv3Key, v.EncryptCertificate.AssociatedData, v.EncryptCertificate.Nonce, v.EncryptCertificate.Ciphertext)
		if e != nil {
			return nil, e
		}
		certs.Data[i].PublicKey = c
	}
	return certs.Data, nil
}
