package core

import (
	"context"
	"fmt"
	"strings"
	"wechat/utils"
)

// Sign 对信息使用 SHA256WithRSA 算法进行签名
func (s *Client) Sign(_ context.Context, message string) (*string, error) {
	if s.PrivateKey == nil {
		return nil, fmt.Errorf("you must set privatekey to use SHA256WithRSASigner")
	}
	if strings.TrimSpace(s.CertificateSerialNo) == "" {
		return nil, fmt.Errorf("you must set mch certificate serial no to use SHA256WithRSASigner")
	}
	signature, err := utils.SignSHA256WithRSA(message, s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &signature, nil
}

// Algorithm 返回使用的签名算法：SHA256-RSA2048
func (s *Client) Algorithm() string {
	return "SHA256-RSA2048"
}
