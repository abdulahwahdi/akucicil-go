package akucicil

import "time"

type InstallmentPlanResp struct {
	Code             string           `json:"code"`
	Message          string           `json:"message"`
	InstallmentPlans InstallmentPlans `json:"installmentPlans"`
}

type InstallmentPlans []InstallmentPlan

type InstallmentPlan struct {
	RepaymentPeriod  int64      `json:"repaymentPeriod"`
	MonthlyPrincipal int64      `json:"monthlyPrincipal"`
	MonthlyRepayment int64      `json:"monthlyRepayment"`
	TotalRepayment   int64      `json:"totalRepayment"`
	MonthlyFee       MonthlyFee `json:"monthlyFee"`
	SummaryFee       SummaryFee `json:"summaryFee"`
}

type MonthlyFee struct {
	MonthlySumFee int64 `json:"monthlySumFee"`
	InterestFee   int64 `json:"interestFee"`
	HandlingFee   int64 `json:"handlingFee"`
}

type SummaryFee struct {
	TotalFee         int64 `json:"totalFee"`
	TotalInterestFee int64 `json:"totalInterestFee"`
	TotalHandlingFee int64 `json:"totalHandlingFee"`
}

type CheckoutOrderResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Order   Order `json:"order"`
	Error   Error  `json:"error"`
}

type Order struct {
	OrderID            string    `json:"orderId"`
	PartnerOrderID     string    `json:"partnerOrderId"`
	PartnerOrderAmount int64    `json:"partnerOrderAmount"`
	OrderStatus        string    `json:"orderStatus"`
	OrderStatusDesc    string    `json:"orderStatusDesc"`
	PaymentEntryURL    string    `json:"paymentEntryUrl"`
	Metadata           string    `json:"metadata"`
	CreateTime         time.Time `json:"createTime"`
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}