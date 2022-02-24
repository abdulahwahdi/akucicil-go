package akucicil

type AkucicilService interface {
	GetInstallmentPlan(req InstallmentPlanReq) (resp InstallmentPlanResp, err error)
	CheckoutOrder(req CheckoutOrderReq) (resp CheckoutOrderResp, err error)
	Cancel(req CancelOrderReq)(resp CancelResp, err error)
	GetDetailOrder(patnerOrderId string)(resp DetailOrderResp, err error)
	GetSignature(req SignatureReq) (resp string, err error)
}
