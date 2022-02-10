package model

// 申请结果
type ApplymentRes struct {
	ApplymentID int `json:"applyment_id"`
}

// 查询申请结果
type QueryApplymentRes struct {
	ApplymentID       int            `json:"applyment_id"`           // 微信支付申请单号
	BusinessCode      string         `json:"business_code"`          // 业务申请编号
	SubMchid          string         `json:"sub_mchid,omitempty"`    // 特约商户号
	ApplymentState    string         `json:"applyment_state"`        // 申请单状态
	ApplymentStateMsg string         `json:"applyment_state_msg"`    // 申请状态描述
	AuditDetail       *[]AuditDetail `json:"audit_detail,omitempty"` // 驳回原因详情
	SignUrl           string         `json:"sign_url,omitempty"`     // 超级管理员签约链接
}

type AuditDetail struct {
	Field        string `json:"field"`         // 提交申请单的资料项字段名。
	FieldName    string `json:"field_name"`    // 提交申请单的资料项字段名称。
	RejectReason string `json:"reject_reason"` // 提交资料项被驳回的原因。
}

// 查询结算账户结果
type ApplymentBankRes struct {
	AccountType      string `json:"account_type"`                 // 账户类型
	AccountBank      string `json:"account_bank"`                 // 开户银行
	BankName         string `json:"bank_name,omitempty"`          // 开户银行全称（含支行）
	BankBranchID     string `json:"bank_branch_id,omitempty"`     // 开户银行联行号
	AccountNumber    string `json:"account_number"`               // 银行账号
	VerifyResult     string `json:"verify_result,omitempty"`      // 汇款验证结果
	VerifyFailReason string `json:"verify_fail_reason,omitempty"` //汇款验证失败原因
}
