package utils

import "time"

// HTTP 请求报文 Header 相关常量
const (
	Authorization = "Authorization"  // Header 中的 Authorization 字段
	Accept        = "Accept"         // Header 中的 Accept 字段
	ContentType   = "Content-Type"   // Header 中的 ContentType 字段
	ContentLength = "Content-Length" // Header 中的 ContentLength 字段
	UserAgent     = "User-Agent"     // Header 中的 UserAgent 字段
)

// 常用 ContentType
const (
	ApplicationJSON = "application/json"
	ImageJPG        = "image/jpg"
	ImagePNG        = "image/png"
	VideoMP4        = "video/mp4"
)

const (
	DefaultTimeout = 30 * time.Second // HTTP 请求默认超时时间
)
