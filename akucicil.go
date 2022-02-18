package akucicil

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hashicorp/go-uuid"
	"io"
	"time"
)

type Akucicil struct {
	BaseUrl    string
	KeyVersion string
	ApiKey     string
	client     *akucicilHttp
	PrivateKey []byte
	*logger
}

func New(baseUrl string, keyVersion string, apiKey string, privateKey []byte, timeout time.Duration) *Akucicil {
	httpRequest := newRequest(timeout)
	return &Akucicil{
		BaseUrl:    baseUrl,
		KeyVersion: keyVersion,
		ApiKey:     apiKey,
		client:     httpRequest,
		PrivateKey: privateKey,
		logger:     newLogger(),
	}
}

func (ac *Akucicil) call(method string, path string, body io.Reader, v interface{}, headers map[string]string) error {
	ac.info().Println("Starting http call..")
	ac.info().Println(path)
	return ac.client.exec(method, path, body, v, headers)
}

func (ac *Akucicil) GetInstallmentPlan(req InstallmentPlanReq) (resp InstallmentPlanResp, err error) {
	ac.info().Println("Starting Get Installment Plan URL Akucicil")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			ac.error().Println(err.Error())
		}
	}()

	var response InstallmentPlanResp

	// set url
	uri := fmt.Sprintf("/openpay/v1/%s/checkout/installment-plan", ac.ApiKey)
	url := fmt.Sprintf("%s%s", ac.BaseUrl, uri)

	method := "POST"
	uuidData, _ := uuid.GenerateUUID()
	reqTime := time.Now().Format(time.RFC3339)
	//Marshal Payload
	payload, errPayload := json.Marshal(req)
	if errPayload != nil {
		return response, errPayload
	}

	payloadSignature := SignatureReq{
		Method:  method,
		Path:    uri,
		ReqID:   uuidData,
		ReqTime: reqTime,
		Body:    string(payload),
	}

	dataSignature, errSignature := ac.GetSignature(payloadSignature)
	if errSignature != nil {
		return response, errSignature
	}

	signature := fmt.Sprintf("Signature: alg=RSA256, keyVersion=%s, signature=%s", ac.KeyVersion, dataSignature)

	// set header
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["X-OP-ApiKey"] = ac.ApiKey
	headers["X-OP-ReqId"] = uuidData
	headers["X-OP-ReqTime"] = reqTime
	headers["Signature"] = signature

	err = ac.call(method, url, bytes.NewBuffer(payload), &response, headers)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ac *Akucicil) CheckoutOrder(req CheckoutOrderReq) (resp CheckoutOrderResp, err error) {
	ac.info().Println("Starting Get Installment Plan URL Akucicil")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			ac.error().Println(err.Error())
		}
	}()

	var response CheckoutOrderResp

	// set url
	uri := fmt.Sprintf("/openpay/v1/%s/checkout/orders/%s", ac.ApiKey, req.PartnerOrderID)
	url := fmt.Sprintf("%s%s", ac.BaseUrl, uri)

	method := "POST"
	uuidData, _ := uuid.GenerateUUID()
	reqTime := time.Now().Format(time.RFC3339)
	//Marshal Payload
	payload, errPayload := json.Marshal(req)
	if errPayload != nil {
		return response, errPayload
	}

	payloadSignature := SignatureReq{
		Method:  method,
		Path:    uri,
		ReqID:   uuidData,
		ReqTime: reqTime,
		Body:    string(payload),
	}

	dataSignature, errSignature := ac.GetSignature(payloadSignature)
	if errSignature != nil {
		return response, errSignature
	}

	signature := fmt.Sprintf("Signature: alg=RSA256, keyVersion=%s, signature=%s", ac.KeyVersion, dataSignature)

	// set header
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["X-OP-ApiKey"] = ac.ApiKey
	headers["X-OP-ReqId"] = uuidData
	headers["X-OP-ReqTime"] = reqTime
	headers["Signature"] = signature

	err = ac.call(method, url, bytes.NewBuffer(payload), &response, headers)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ac *Akucicil) Cancel(req CancelOrderReq) (resp CancelResp, err error) {
	ac.info().Println("Starting Cancel Order Akucicil")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			ac.error().Println(err.Error())
		}
	}()

	var response CancelResp

	// set url
	uri := fmt.Sprintf("/openpay/v1/%s/orders/%s/cancel", ac.ApiKey, req.PartnerOrderId)
	url := fmt.Sprintf("%s%s", ac.BaseUrl, uri)

	method := "POST"
	uuidData, _ := uuid.GenerateUUID()
	reqTime := time.Now().Format(time.RFC3339)
	//Marshal Payload
	payload, errPayload := json.Marshal(req)
	if errPayload != nil {
		return response, errPayload
	}

	payloadSignature := SignatureReq{
		Method:  method,
		Path:    uri,
		ReqID:   uuidData,
		ReqTime: reqTime,
		Body:    string(payload),
	}

	dataSignature, errSignature := ac.GetSignature(payloadSignature)
	if errSignature != nil {
		return response, errSignature
	}

	signature := fmt.Sprintf("Signature: alg=RSA256, keyVersion=%s, signature=%s", ac.KeyVersion, dataSignature)

	// set header
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["X-OP-ApiKey"] = ac.ApiKey
	headers["X-OP-ReqId"] = uuidData
	headers["X-OP-ReqTime"] = reqTime
	headers["Signature"] = signature

	err = ac.call(method, url, bytes.NewBuffer(payload), &response, headers)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ac *Akucicil) getPrivateKey() (key *rsa.PrivateKey) {
	var err error
	verifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(ac.PrivateKey)
	if err != nil {
		return nil
	}
	return verifyKey
}

func (ac *Akucicil) GetSignature(req SignatureReq) (resp string, err error) {
	contentSignature := fmt.Sprintf("%s|%s|%s|%s|%s|%s", req.Method, req.Path, ac.ApiKey, req.ReqID, req.ReqTime, req.Body)
	h := sha256.New()
	h.Write([]byte(contentSignature))
	d := h.Sum(nil)

	key := ac.getPrivateKey()

	signature, errSig := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, d)
	if errSig != nil {
		return "", errSig
	}

	signatureString := base64.StdEncoding.EncodeToString(signature)

	return signatureString, nil
}
