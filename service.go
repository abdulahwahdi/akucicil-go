package akucicil

type AkucicilService interface {
	GetInstallmentPlan(req InstallmentPlanReq) (resp InstallmentPlanResp, err error)
	CheckoutOrder(req CheckoutOrderReq) (resp CheckoutOrderResp, err error)
	Cancel(req CancelOrderReq)(resp CancelResp, err error)
}
