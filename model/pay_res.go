package model

type JsApiPayRes struct {
	AppId     string `json:"appId"`     // 商户申请的公众号对应的appid, 或者 小程序下单
	TimeStamp string `json:"timeStamp"` // 时间戳
	NonceStr  string `json:"nonceStr"`  // 随机字符串,不长于32位。
	Package   string `json:"package"`   // JSAPI下单接口返回的prepay_id参数值，提交格式如：prepay_id=***示例值：prepay_id=wx201410272009395522657a690389285100
	SignType  string `json:"signType"`  // 签名类型，默认为RSA，仅支持RSA
	PaySign   string `json:"paySign"`   // 签名
}

type AppPayRes struct {
	AppId     string `json:"appId"`     // 微信开放平台审核通过的移动应用appid ，为二级商户申请的应用appid
	PartnerID string `json:"partnerid"` // 商户号
	PrepayID  string `json:"prepayid"`  // 微信返回的支付交易会话id。示例值： WX1217752501201407033233368018
	Package   string `json:"package"`   // 暂填写固定值Sign=WXPay
	NonceStr  string `json:"nonceStr"`  // 随机字符串,不长于32位。
	TimeStamp string `json:"timeStamp"` // 时间戳
	PaySign   string `json:"paySign"`   // 签名
}

type RefundRes struct {
	RefundID            string             `json:"refund_id"`                  // 微信支付退款单号
	OutRefundNo         string             `json:"out_refund_no"`              // 商户退款单号
	TransactionID       string             `json:"transaction_id"`             // 微信支付订单号
	OutTradeNo          string             `json:"out_trade_no,omitempty"`     // 商户支付订单号
	Channel             string             `json:"channel"`                    // 退款渠道// ORIGINAL：原路退款 // BALANCE：退回到余额 // OTHER_BALANCE：原账户异常退到其他余额账户 // OTHER_BANKCARD：原银行卡异常退到其他银行卡
	UserReceivedAccount string             `json:"user_received_account"`      // 退款入账账户 下几种情况：1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通
	SuccessTime         string             `json:"success_time,omitempty"`     // 退款成功时间
	CreateTime          string             `json:"create_time"`                // 退款创建时间
	Status              string             `json:"status"`                     // 退款状态 SUCCESS：退款成功 CLOSED：退款关闭 PROCESSING：退款处理中 ABNORMAL：退款异常
	FundsAccount        string             `json:"funds_account,omitempty"`    // 资金账户 UNSETTLED : 未结算资金 AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额  OPERATION : 运营户 BASIC : 基本账户（含可用余额和不可用余额）
	Amount              *RefundAmountRes   `json:"amount"`                     // 金额信息
	PromotionDetail     *[]PromotionDetail `json:"promotion_detail,omitempty"` // 优惠退款信息
}

// 金额信息
type RefundAmountRes struct {
	Total            int     `json:"total"`             // 订单总金额，单位为分
	Refund           int     `json:"refund"`            // 退款金额
	From             *[]From `json:"from"`              // 退款出资账户及金额
	PayerTotal       int     `json:"payer_total"`       // 用户支付金额
	PayerRefund      int     `json:"payer_refund"`      // 退款给用户的金额，不包含所有优惠券金额
	SettlementRefund int     `json:"settlement_refund"` // 应结退款金额 去掉非充值代金券退款金额后的退款金额，单位为分，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementTotal  int     `json:"settlement_total"`  // 应结订单金额 应结订单金额=订单金额-免充值代金券金额，应结订单金额<=订单金额，单位为分
	DiscountRefund   int     `json:"discount_refund"`   // 优惠退款金额 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠，单位为分
	Currency         string  `json:"currency"`          // 退款币种 目前只支持人民币：CNY
}

// 优惠退款信息
type PromotionDetail struct {
	PromotionID  string         `json:"promotion_id"`           // 券ID
	Scope        string         `json:"scope"`                  // 优惠范围	GLOBAL：全场代金券 SINGLE：单品优惠
	Type         string         `json:"type"`                   // 优惠类型  COUPON：代金券，需要走结算资金的充值型代金券 DISCOUNT：优惠券，不走结算资金的免充值型优惠券
	Amount       int            `json:"amount"`                 // 优惠券面额
	RefundAmount int            `json:"refund_amount"`          // 优惠退款金额
	GoodsDetail  *[]GoodsDetail `json:"goods_detail,omitempty"` // 商品列表
}
