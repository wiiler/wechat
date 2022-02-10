package model

// 进件结构体
type ApplymentData struct {
	BusinessCode    string           `json:"business_code"`   // 业务申请编号
	ContactInfo     *ContactInfo     `json:"contact_info"`    // 超级管理员信息
	SubjectInfo     *SubjectInfo     `json:"subject_info"`    // 主体资料
	BusinessInfo    *BusinessInfo    `json:"business_info"`   // 经营资料
	SettlementInfo  *SettlementInfo  `json:"settlement_info"` // 结算规则
	BankAccountInfo *BankAccountInfo `json:"bank_account_info"`
	AdditionInfo    *AdditionInfo    `json:"addition_info"`
}

// 超级管理员信息
type ContactInfo struct {
	ContactName     string `json:"contact_name"`                // 超级管理员姓名 加密
	ContactIDMumber string `json:"contact_id_number,omitempty"` // 超级管理员身份证件号码 加密 ContactIDMumber、OpenID 二选一
	OpenID          string `json:"openid,omitempty"`            // 超级管理员微信openid ContactIDMumber、OpenID 二选一
	MobilePhone     string `json:"mobile_phone"`                // 联系手机 加密
	ContactEmail    string `json:"contact_email"`               // 联系邮箱 加密
}

// 主体资料 无需加密
type SubjectInfo struct {
	SubjectType string `json:"subject_type"` // 主体类型
	// 	SUBJECT_TYPE_INDIVIDUAL（个体户） 营业执照上的主体类型一般为个体户、个体工商户、个体经营；
	// SUBJECT_TYPE_ENTERPRISE（企业） 营业执照上的主体类型一般为有限公司、有限责任公司；
	// SUBJECT_TYPE_INSTITUTIONS（党政、机关及事业单位） 包括国内各级、各类政府机构、事业单位等（如 公安、党团、司法、交通、旅游、工商税务、市政、医疗、教育、学校等机构）；
	// SUBJECT_TYPE_OTHERS（其他组织） 不属于企业、政府/事业单位的组织机构（如社会团体、民办非企业、基金会），要求机构已办理组织机构代码证
	BusinessLicenseInfo *BusinessLicenseInfo `json:"business_license_info"` // 营业执照信息 企业和个体工商户必填
	IdentityInfo        *IdentityInfo        `json:"identity_info"`         // 经营者/法人身份证件
}

// 营业执照信息
type BusinessLicenseInfo struct {
	LicenseCopy   string `json:"license_copy"`   // 营业执照照片
	LicenseMumber string `json:"license_number"` // 注册号/统一社会信用代码
	MerchantMame  string `json:"merchant_name"`  // 商户名称
	LegalPerson   string `json:"legal_person"`   // 个体户经营者/法人姓名
}

// 经营者/法人身份证件
type IdentityInfo struct {
	IDdocType  string      `json:"id_doc_type"`  // 证件类型 // IDENTIFICATION_TYPE_IDCARD 中国大陆居民-身份证 // IDENTIFICATION_TYPE_OVERSEA_PASSPORT 其他国家或地区居民-护照 // IDENTIFICATION_TYPE_HONGKONG_PASSPORT 中国香港居民-来往内地通行证 // IDENTIFICATION_TYPE_MACAO_PASSPORT 中国澳门居民-来往内地通行证 // IDENTIFICATION_TYPE_TAIWAN_PASSPORT 中国台湾居民-来往大陆通行证
	IDCardInfo *IDCardInfo `json:"id_card_info"` // 身份证信息
	Owner      bool        `json:"owner"`        //经营者/法人是否为受益人 则填写：true
}

// 身份证信息
type IDCardInfo struct {
	IDcardCopy      string `json:"id_card_copy"`      // 身份证人像面照片
	IDcardNational  string `json:"id_card_national"`  // 身份证国徽面照片
	IDcardName      string `json:"id_card_name"`      // 身份证姓名 加密
	IDcardNumber    string `json:"id_card_number"`    // 身份证号码 加密
	CardPeriodBegin string `json:"card_period_begin"` // 身份证有效期开始时间
	CardPeriodEnd   string `json:"card_period_end"`   // 身份证有效期结束时间 2036-06-06
}

// 经营资料
type BusinessInfo struct {
	MerchantShortname string     `json:"merchant_shortname"` // 商户简称
	ServicePhone      string     `json:"service_phone"`      // 客服电话
	SalesInfo         *SalesInfo `json:"sales_info"`
}

// 经营场景
type SalesInfo struct {
	SalesScenesType []string      `json:"sales_scenes_type"` // 经营场景类型
	BizStoreInfo    *BizStoreInfo `json:"biz_store_info"`    // 线下门店场景
	// SALES_SCENES_STORE 线下门店
	// SALES_SCENES_MP 公众号
	// SALES_SCENES_MINI_PROGRAM 小程序
	// SALES_SCENES_WEB 互联网
	// SALES_SCENES_APP APP
	// SALES_SCENES_WEWORK 企业微信
}

// 线下门店场景
type BizStoreInfo struct {
	BizStoreName     string   `json:"biz_store_name"`     // 门店名称
	BizAddressCode   string   `json:"biz_address_code"`   // 门店省市编码
	BizStoreAddress  string   `json:"biz_store_address"`  // 门店地址
	StoreEntrancePic []string `json:"store_entrance_pic"` // 门店门头照片
	IndoorPic        []string `json:"indoor_pic"`         // 店内环境照片
}

// 结算规则
type SettlementInfo struct {
	SettlementID        string   `json:"settlement_id"`                  // 入驻结算规则ID
	QualificationType   string   `json:"qualification_type"`             // 所属行业
	Qualifications      []string `json:"qualifications,omitempty"`       // 特殊资质图片
	ActivitiesID        string   `json:"activities_id,omitempty"`        // 优惠费率活动ID "20191030111cff5b5e"
	ActivitiesRate      string   `json:"activities_rate,omitempty"`      // 优惠费率活动值 默认0.38 填了汇率后需要填写 ActivitiesID
	ActivitiesAdditions []string `json:"activities_additions,omitempty"` // 优惠费率活动补充材料
}

// 结算银行账户
type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"`        // 账户类型 // BANK_ACCOUNT_TYPE_CORPORATE 对公银行账户 // BANK_ACCOUNT_TYPE_PERSONAL 经营者个人银行卡
	AccountName     string `json:"account_name"`             // 开户名称 加密
	AccountBank     string `json:"account_bank"`             // 开户银行
	BankAddressCode string `json:"bank_address_code"`        // 开户银行省市编码
	BankBranchID    string `json:"bank_branch_id,omitempty"` // 开户银行联行号
	BankName        string `json:"bank_name,omitempty"`      // 开户银行全称（含支行）
	AccountNumber   string `json:"account_number"`           // 银行账号 加密
}

// 补充材料
type AdditionInfo struct {
	LegalPersonCommitment string   `json:"legal_person_commitment"` // 法人开户承诺函
	LegalPersonVideo      string   `json:"legal_person_video"`      // 法人开户意愿视频
	BusinessAdditionPics  []string `json:"business_addition_pics"`  // 补充材料
	BusinessAdditionMsg   string   `json:"business_addition_msg"`   // 补充说明
}

//  修改结算账户
type UpdateMchBankData struct {
	AccountType     string `json:"account_type"`             //  账户类型
	AccountBank     string `json:"account_bank"`             // 开户银行
	BankAddressCode string `json:"bank_address_code"`        //  开户银行省市编码 需至少精确到市
	BankName        string `json:"bank_name,omitempty"`      // 开户银行全称（含支行）
	BankBranchID    string `json:"bank_branch_id,omitempty"` // 开户银行联行号
	AccountNumber   string `json:"account_number"`           // 银行账号 加密
}
