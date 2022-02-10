package model

// 平台证书返回结果
type CertsRes struct {
	Data []*ItemCerts `json:"data"`
}

//  平台证书item
type ItemCerts struct {
	EffectiveTime      string       `json:"effective_time"`
	ExpireTime         string       `json:"expire_time"`
	PublicKey          string       `json:"public_key"`
	SerialNo           string       `json:"serial_no"`
	EncryptCertificate *encryptCert `json:"encrypt_certificate"`
}

//  平台加密数据
type encryptCert struct {
	Algorithm      string `json:"algorithm"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
}
