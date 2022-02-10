package core

import "wechat/utils"

// SDK 相关信息
const (
	Version         = "0.2.8"                      // SDK 版本
	UserAgentFormat = "WechatPay-Go/%s (%s) GO/%s" // UserAgent中的信息
)

// 请求报文签名相关常量
const (
	SignatureMessageFormat    = "%s\n%s\n%d\n%s\n%s\n"                                                                // 数字签名原文格式
	HeaderAuthorizationFormat = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\"" // HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
)

// HTTP 应答报文 Header 相关常量
const (
	WechatPayTimestamp = "Wechatpay-Timestamp" // 微信支付回包时间戳
	WechatPayNonce     = "Wechatpay-Nonce"     // 微信支付回包随机字符串
	WechatPaySignature = "Wechatpay-Signature" // 微信支付回包签名信息
	WechatPaySerial    = "Wechatpay-Serial"    // 微信支付回包平台序列号
	RequestID          = "Request-Id"          // 微信支付回包请求ID
)

// 时间相关常量
const (
	FiveMinute = 5 * 60 // 回包校验最长时间（秒）
)

// 默认请求头
var (
	DefaultHeader = map[string]string{
		utils.ContentType: utils.ApplicationJSON,
	}
)
