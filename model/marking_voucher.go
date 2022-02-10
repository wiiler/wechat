package model

// 优惠卷
type Stocks struct {
	StockName          string         `json:"stock_name"`             //批次名称
	Comment            string         `json:"comment"`                // 批次备注
	BelongMerchant     string         `json:"belong_merchant"`        // 归属商户号
	AvailableBeginTime string         `json:"available_begin_time"`   // 可用时间-开始时间 2015-05-20T13:29:35.120+08:00
	AvailableEndTime   string         `json:"available_end_time"`     // 可用时间-结束时间
	StockUseRule       *StockUseRule  `json:"stock_use_rule"`         // 发放规则
	PatternInfo        *PatternInfo   `json:"pattern_info,omitempty"` //样式设置
	CouponUseRule      *CouponUseRule `json:"coupon_use_rule"`        // 核销规则
	NoCash             bool           `json:"no_cash"`                // 营销经费 true 免充值 false 预充值
	StockType          string         `json:"stock_type"`             // 批次类型 仅支持 NORMAL 固定面额满减券批次
	OutRequestNo       string         `json:"out_request_no"`         // 商户创建批次凭据号
	ExtInfo            string         `json:"ext_info,omitempty"`     //扩展属性 {'exinfo1':'1234','exinfo2':'3456'}
}

// 发放规则
type StockUseRule struct {
	MaxCoupons         uint64 `json:"max_coupons"`                 // 发放总上限
	MaxAmount          uint64 `json:"max_amount"`                  // 总预算 单位 分 max_amount需要等于coupon_amount（面额） * max_coupons（发放总上限）
	MaxAmountByDay     uint64 `json:"max_amount_by_day,omitempty"` // 单天预算发放上限  不能大于总预算 max_amount_by_day不可以为0
	MaxCouponsPerUser  uint32 `json:"max_coupons_per_user"`        // 单个用户可领个数 不能大于发放总个数 最少为1个，最多为60个
	NaturalPersonLimit bool   `json:"natural_person_limit"`        // 是否开启自然人限制 false 否  true 是
	PreventApiAbuse    bool   `json:"prevent_api_abuse"`           // 是否开启防刷拦截

}

// 样式设置
type PatternInfo struct {
	Description     string `json:"description"`                // 使用说明 最多1000个UTF8字符
	MerchantLogo    string `json:"merchant_logo,omitempty"`    // 商户logo 商户logo大小需为120像素*120像素。2、支持JPG/JPEG/PNG格式，且图片小于1M。3、最多128个UTF8字符
	MerchantName    string `json:"merchant_name,omitempty"`    // 品牌名称 最多12个中文汉字 最多36个英文字符
	BackgroundColor string `json:"background_color,omitempty"` // 背景颜色 可设置10种颜色
	CouponImage     string `json:"coupon_image,omitempty"`     // 券详情图片
}

// 核销规则
type CouponUseRule struct {
	CouponAvailableTime *CouponAvailableTime `json:"coupon_available_time,omitempty"` // 券生效时间 该字段暂未开放
	FixedNormalCoupon   *FixedNormalCoupon   `json:"fixed_normal_coupon,omitempty"`   // 固定面额满减券使用规则
	GoodsTag            *[]string            `json:"goods_tag,omitempty"`             // 订单优惠标记
	LimitPay            string               `json:"limit_pay,omitempty"`             // 指定付款方式 指定付款方式的交易可核销/使用代金券，可指定零钱付款、指定银行卡付款，需填入支付方式编码， 不在此列表中的银行卡，暂不支持此功能。黄色标记部分为不可使用的银行。校验规则：条目个数限制为【1，1】。零钱：CFT示例值：ICBC_CREDIT
	LimitCard           *LimitCard           `json:"limit_card,omitempty"`            // 指定银行卡BIN
	TradeType           *[]string            `json:"trade_type,omitempty"`            // 支付方式 	MICROAPP 小程序支付 APPPAY APP支付 PPAY 免密支付 CARD 刷卡支付 FACE 人脸支付 OTHER 其他支付
	CombineUse          bool                 `json:"combine_use,omitempty"`           //是否可叠加其他优惠
	AvailableItems      *[]string            `json:"available_items,omitempty"`       // 可核销商品编码包含指定SKU商品编码的交易才可核销/使用代金券：活动商户在交易下单时，需传入用户购买的所有SKU商品编码，当命中代金券中设置的商品编码时可享受优惠。
	UnavailableItems    *[]string            `json:"unavailable_items,omitempty"`     // 该字段暂未开放 不可核销商品编码
	AvailableMerchants  *[]string            `json:"available_merchants"`             //可用商户号
}

type CouponAvailableTime struct {
	FixAvailableTime          *FixAvailableTime `json:"fix_available_time,omitempty"`           // 该字段暂未开放 固定时间段可用 允许指定券在特殊时间段生效。当设置固定时间段可用时不可设置领取后N天有效
	SecondDayAvailable        bool              `json:"second_day_available,omitempty"`         // 领取后N天有效
	AvailableTimeAfterReceive uint32            `json:"available_time_after_receive,omitempty"` // 领取后有效时间 领取后，券的结束时间为领取N天后，如设置领取后7天有效，那么7月1日领券，在7月7日23:59:59失效（在可用时间内计算失效时间，若券还未到领取后N天，但是已经到了可用结束时间，那么也会过期）
}
type FixAvailableTime struct {
	AvailableWeekDay uint32 `json:"available_week_day,omitempty"` // 可用星期，0代表周日生效，1代表周一生效，以此类推；不填则代表在可用时间内周一至周日都生效。
	BeginTime        uint32 `json:"begin_time,omitempty"`         // 允许指定特殊生效星期数中的具体生效的时间段。当天开始时间，单位：秒
	EndTime          uint32 `json:"end_time,omitempty"`           // 允许指定特殊生效星期数中的具体生效的时间段。当天结束时间，单位：秒，默认为23点59分59秒。
}

// 固定面额满减券使用规则
type FixedNormalCoupon struct {
	CouponAmount       uint32 `json:"coupon_amount"`       // 面额 分
	TransactionMinimum uint32 `json:"transaction_minimum"` // 使用券金额门槛，单位：分
}
type LimitCard struct {
	Name string    `json:"name,omitempty"` // 银行卡名称 将在微信支付收银台向用户展示，最多4个中文汉字示例值：精粹白金
	Bin  *[]string `json:"bin,omitempty"`  // 指定卡BIN 使用指定卡BIN的银行卡支付方可享受优惠，按json格式特殊规则：单个卡BIN的字符长度为【6,9】,条目个数限制为【1,10】。示例值：['62123456','62123457']
}

type CreateCouponRsp struct {
	CreateTime string `json:"create_time"`
	StockID    string `json:"stock_id"`
}
