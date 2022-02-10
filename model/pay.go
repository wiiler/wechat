package model

type JsApiPay struct {
	SpAppID     string      `json:"sp_appid"`              // 服务商应用ID
	SpMchID     string      `json:"sp_mchid"`              // 服务商户号
	SubAppID    string      `json:"sub_appid,omitempty"`   // 子商户应用ID
	SubMchID    string      `json:"sub_mchid"`             // 子商户号
	Description string      `json:"description"`           // 商品描述
	OutTradeNo  string      `json:"out_trade_no"`          // 商户订单号
	TimeExpire  string      `json:"time_expire,omitempty"` // 交易结束时间
	Attach      string      `json:"attach,omitempty"`      // 附加数据
	NotifyUrl   string      `json:"notify_url"`            // 通知地址
	GoodsTag    string      `json:"goods_tag"`             // 订单优惠标记
	SettleInfo  *SettleInfo `json:"settle_info,omitempty"` // 结算信息
	Amount      *Amount     `json:"amount"`                // 订单金额
	Payer       *Payer      `json:"payer"`                 // 支付者
	Detail      *Detail     `json:"detail"`                // 优惠功能
	SceneInfo   *SceneInfo  `json:"scene_info,omitempty"`  // 场景信息
}

//  app or native
type AppNativePay struct {
	SpAppID     string      `json:"sp_appid"`              // 服务商应用ID
	SpMchID     string      `json:"sp_mchid"`              // 服务商户号
	SubAppID    string      `json:"sub_appid,omitempty"`   // 子商户应用ID
	SubMchID    string      `json:"sub_mchid"`             // 子商户号
	Description string      `json:"description"`           // 商品描述
	OutTradeNo  string      `json:"out_trade_no"`          // 商户订单号
	TimeExpire  string      `json:"time_expire,omitempty"` // 交易结束时间
	Attach      string      `json:"attach,omitempty"`      // 附加数据
	NotifyUrl   string      `json:"notify_url"`            // 通知地址
	GoodsTag    string      `json:"goods_tag"`             // 订单优惠标记
	SettleInfo  *SettleInfo `json:"settle_info,omitempty"` // 结算信息
	Amount      *Amount     `json:"amount"`                // 订单金额
	Detail      *Detail     `json:"detail"`                // 优惠功能
	SceneInfo   *SceneInfo  `json:"scene_info,omitempty"`  // 场景信息
}

//  结算信息
type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing"` // 是否指定分账
}

// 订单金额
type Amount struct {
	Total    int    `json:"total"`              // 订单总金额，单位为分
	Currency string `json:"currency,omitempty"` // CNY：人民币，境内商户号仅支持人民币
}

// 支付者
type Payer struct {
	SpOpenID  string `json:"sp_openid,omitempty"`  // 用户服务标识
	SubOpenID string `json:"sub_openid,omitempty"` // 用户子标识
}

//  优惠功能
type Detail struct {
	CostPrice   int            `json:"cost_price"`   // 订单原价
	InvoiceID   string         `json:"invoice_id"`   // 商品小票ID
	GoodsDetail *[]GoodsDetail `json:"goods_detail"` // 单品列表
}

// 单品列表
type GoodsDetail struct {
	MerchantGoodsID  string `json:"merchant_goods_id"`            // 商户侧商品编码
	WechatpayGoodsID string `json:"wechatpay_goods_id,omitempty"` // 微信侧商品编码
	GoodsName        string `json:"goods_name"`                   // 商品名称
	Quantity         int    `json:"quantity"`                     // 商品数量
	UnitPrice        int    `json:"unit_price"`                   // 商品单价
}

// 场景信息
type SceneInfo struct {
	PayerClientIP string     `json:"payer_client_ip"`      // 用户终端IP
	DeviceID      string     `json:"device_id"`            // 商户端设备号
	StoreInfo     *StoreInfo `json:"store_info,omitempty"` // 商户门店信息
}

// 商户门店信息
type StoreInfo struct {
	ID       string `json:"id"`                  // 门店编号
	Name     string `json:"name,omitempty"`      // 门店名称
	AreaCode string `json:"area_code,omitempty"` // 地区编码
	Address  string `json:"address,omitempty"`   // 详细地址
}

// 退款
type Refund struct {
	SubMchID      string         `json:"sub_mchid"`                // 子商户号
	TransactionID string         `json:"transaction_id,omitempty"` // 微信支付订单号
	OutTradeNo    string         `json:"out_trade_no,omitempty"`   // 商户支付订单号
	OutRefundNo   string         `json:"out_refund_no"`            // 商户退款单号
	Reason        string         `json:"reason,omitempty"`         // 退款原因
	NotifyUrl     string         `json:"notify_url,omitempty"`     // 通知地址
	FundsAccount  string         `json:"funds_account,omitempty"`  // 退款资金来源 AVAILABLE：可用余额账户
	Amount        *RefundAmount  `json:"amount"`                   // 金额信息
	GoodsDetail   *[]GoodsDetail `json:"goods_detail,omitempty"`   // 退款商品
}

// 金额信息
type RefundAmount struct {
	Refund   int     `json:"refund"`   // 退款金额
	Total    int     `json:"total"`    // 原订单金额
	Currency string  `json:"currency"` // 退款币种 CNY
	From     *[]From `json:"from"`     // 退款出资账户及金额
}

// 退款出资账户及金额
type From struct {
	Account string `json:"account"` //下面枚举值多选一。 AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额
	Amount  int    `json:"amount"`  // 出资金额
}
