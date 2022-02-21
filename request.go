package akucicil

import "time"

type InstallmentPlanReq struct {
	UserToken          string `json:"userToken,omitempty"`
	PartnerOrderAmount int64  `json:"partnerOrderAmount"`
	Currency           string `json:"currency"`
	Buyer              Buyer  `json:"buyer,omitempty"`
	Items              Items  `json:"items"`
}

type Buyer struct {
	BuyerID          string `json:"buyerId"`
	BuyerPhone       string `json:"buyerPhone,omitempty"`
	BuyerEmail       string `json:"buyerEmail,omitempty"`
	RegistrationDays string `json:"registrationDays,omitempty"`
}

type Seller struct {
	SellerID   string `json:"sellerId"`
	SellerName string `json:"sellerName"`
}

type Items []Item

type Item struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	SkuID        string `json:"skuId"`
	SkuName      string `json:"skuName"`
	UnitPrice    int64  `json:"unitPrice"`
	Qty          int64  `json:"qty"`
	PictureUrl   string `json:"pictureUrl,omitempty"`
	Seller       Seller `json:"seller"`
}

type CheckoutOrderReq struct {
	UserToken          string    `json:"userToken,omitempty"`
	PartnerOrderID     string    `json:"partnerOrderId"`
	PartnerOrderAmount int64     `json:"partnerOrderAmount"`
	Currency           string    `json:"currency"`
	RedirectURL        string    `json:"redirectUrl"`
	CallbackURL        string    `json:"callbackUrl"`
	ExpireTime         time.Time `json:"expireTime"`
	Buyer              Buyer     `json:"buyer"`
	Shipping           Shipping  `json:"shipping"`
	Items              Items     `json:"items"`
	Metadata           string    `json:"metadata"`
	ExtendInfo         string    `json:"extendInfo"`
}

type Shipping struct {
	ReceiverName  string `json:"receiverName"`
	ReceiverPhone string `json:"receiverPhone"`
	Region        string `json:"region"`
	City          string `json:"city"`
	District      string `json:"district"`
	Street        string `json:"street"`
	Postcode      string `json:"postcode"`
	AddressLine   string `json:"addressLine"`
	CountryCode   string `json:"countryCode"`
}

type SignatureReq struct {
	Method  string
	Path    string
	ReqID   string
	ReqTime string
	Body    string
}

type CancelOrderReq struct {
	PartnerOrderId string `json:"partnerOrderId"`
	OrderId        string `json:"orderId"`
}
