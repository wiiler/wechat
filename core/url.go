package core

// 微信支付 API 地址
const (
	BaseUrl       = "https://api.mch.weixin.qq.com"  // 微信支付 API 地址
	BaseUrlBackup = "https://api2.mch.weixin.qq.com" // 微信支付 API 备份地址
)

// 微信支付 API url
// 	获取平台证书列表
const (
	GetCertsUrl = BaseUrl + "/v3/certificates" // 获取平台证书列表
)

// 进件
const (
	ApplymentUrl      = BaseUrl + "/v3/applyment4sub/applyment/"                     // 特约商户进件提交申请
	QueryApplymentUrl = BaseUrl + "/v3/applyment4sub/applyment/applyment_id/%s"      // GET 特约商户进件查询 s 业务申请编号
	UpdateMchBankUrl  = BaseUrl + "/v3/apply4sub/sub_merchants/%s/modify-settlement" // 修改结算账号 s 商户号
	QueryMchBankUrl   = BaseUrl + "/v3/apply4sub/sub_merchants/%s/settlement"        // GET 查询结算账号 s 商户号
)

// 其他能力
const (
	ImgUploadUrl   = BaseUrl + "/v3/merchant/media/upload"       // 图片上传ImgUpload
	VideoUploadUrl = BaseUrl + "/v3/merchant/media/video_upload" // 视频上传
)

// 优惠卷
const (
	VoucherStocksUrl        = BaseUrl + "/v3/marketing/favor/coupon-stocks"     // 创建优惠卷
	VoucherStartStocksUrl   = BaseUrl + "/v3/marketing/favor/stocks/%s/start"   // 激活优惠卷
	VoucherPauseStocksUrl   = BaseUrl + "/v3/marketing/favor/stocks/%s/pause"   // 暂停优惠卷
	VoucherRestartStocksUrl = BaseUrl + "/v3/marketing/favor/stocks/%s/restart" // 重启代金券
	VoucherGetStocksUrl     = BaseUrl + "/v3/marketing/favor/users/%s/coupons"  // 发放代金劵
)

// 支付
const (
	JsapiPayUrl           = BaseUrl + "/v3/pay/partner/transactions/jsapi"                                    // jsapi 下单
	AppPayUrl             = BaseUrl + "/v3/pay/partner/transactions/app"                                      // App 下单
	NativePayUrl          = BaseUrl + "/v3/pay/partner/transactions/native"                                   // Native下单
	QueryTransactionIDUrl = BaseUrl + "/v3/pay/partner/transactions/id/%s?sp_mchid=%s&sub_mchid=%s"           // 微信支付订单号查询
	QueryOutTradeNoUrl    = BaseUrl + "/v3/pay/partner/transactions/out-trade-no/%s?sp_mchid=%s&sub_mchid=%s" // 商户订单号查询
	CloseOutTradeNoUrl    = BaseUrl + "/v3/pay/partner/transactions/out-trade-no/%s/close"                    // 通过商户订单号关闭订单
	RefundUrl             = BaseUrl + "/v3/refund/domestic/refunds"                                           // 订单退款
	QueryRefundUrl        = BaseUrl + "/v3/refund/domestic/refunds/%s?sub_mchid=%s"                           // 查询退款订单详情 by OutTradeNo
)
