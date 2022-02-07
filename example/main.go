package main

import (
	"encoding/json"
	"fmt"
	"github.com/abdulahwahdi/akucicil-go"
	"time"
)

func main() {
	baseUrl := "https://test-pay.akulaku.com"
	privateKey := []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCig8aw5KFGilNm/ksd1sqUqLYpzKcPgX2v8aVe0ChL2YLbQk0ZW9IPxarTWo90HoAthoWD5llRtB8wvlvsMcUztIgasssURDq6byMUQjfrHxa5nZvLFW/ETnsb0UrkjKBR5n5lItAXM/2Zh/QtUWHZiYeCXjny9rRpWnbbBp6561P3QOPo4oC/uTrWCtAqjHT9v4fTvHnJab6Zb8oklSvRzS6xkfnvry0XAXIdfR5lVzcna0k3xIaLMMXJBozb6nkaOfwLb3LFHeWwjP+T36QDpI4U20XKtn7nmlIoU4oiVvVZcAwzHgcY4zXeOWqC/FLr3DL3JyB7cquobgzEZltTAgMBAAECggEALyb61EJRxMzHrB9mHg8/Eejtia02X5pbpFwBdkQBHRjMDd/rzMrGZQseKFqjOrbKy+q2/HEuQzV2w0bIFvzN0dW0k8KCh6vQOl2amPspjE+l4U8ob8EVf6ihsY2FJFDlBz5QN/nk1IaiMxSuNSPityJfLeU2/Ra1e8fiqmP+U8TMZeSPcjIl7PwESTdjN+q1E6vyD8huMkuKCIe/OFmapCOr94Hhsug8dMheJiHCs+XUfd1LxD70rZogKQL8Gl1Kaqk6laRZl7ZmnfW5mCyQh0SaHZIy4vJhQEOcKQotLv9e7rNQ7tVa2g1dVRzqTQuzaHfEKyBBJdRmOVDgQEAsAQKBgQDOiH4FXRVZdKnNc6MU10+TYFcVS2tz5J0IdeWoOEC6bxTQswfgrXXzgDjry/dHcGXndLiJLpxQ3B1SkrhqPsY5hQ7Qti14NVQ8uKxwmGdjOqMLDcf3vBrwKYq/KCP9OZhszYpH0JgRJgme5gNQ0jWdeSLHgs+nppilPYx20v69UwKBgQDJcE9dfhaGh3FeC/2Am9FyOtuuaAaEzNo5VdqeNnkoIvfpTr2sQEyX6TFqkHou0kEusNRpqbihAPzjoHWe7U8M5RqDHSh2jjniNVTYTa+9AtP2DbxClabnfuTgjXnZmxhp51Jg54Hpn9xboadtQZkoO1yjI5yMGJA/LnwNnUgqAQKBgQDMYkiL2ETjZM8Q8QmG+1fVWXZ0LVGlu6wnDjz6DskHdhvcN/9ouV/LW614szz/ZlyCA4EGyKiKYHloGKgFBCA4FdPGeTJgDi5RdNO5DSjrCsEElFu6bl0eI44Zg4ix9EhCC9l3HhcOiVzl8WuAlLXYI2Nmmq4Sk4vlohYaFM1g3wKBgGOfo+C9Jy7JdYvpqioc5ez0pwkOcYrkNVj92O0+S5VKFgnQ44V5F2hZ8BKf+Y2Gdq48zhCj26fk8S0ygnljLCFYfVAy73wSwxXZAmknoq5745BhIqZblPwQiel0jrrMNbrKqLc8R8ffRAAdZsxmEPUQxrP4PMgcrdpRhxDEpJgBAoGADG7gg91PZZxONNWzOGhX0AFUUGBkm9KVLQgvzaubYNpeFbU4noobplb0bPBaC0Iena5mPFRnTCYTOZJ4wt/k4raEsm8LI+1L4A1eog0lfSFGc0BYtg/rq21OKN00jRSjLRlkILCKw6fSM/WbE6b2Sfyr61VcDDI2QznGm8GXusg=
-----END PRIVATE KEY-----`)
	apiKey := "74967498"
	newAkucicil := akucicil.New(baseUrl, "1.0", apiKey, privateKey, 5*time.Minute)
	//getInstallmentPlant(newAkucicil)
	getCheckoutOrder(newAkucicil)
}

func getInstallmentPlant(newAc akucicil.AkucicilService) {
	jsonData := `{
    "partnerOrderAmount": 72000000,
    "currency": "IDR",
    "buyer": {
        "buyerId": "Buyer-0001"
    },
    "items": [
        {
            "categoryId": "Category-0001",
            "categoryName": "Mobile",
            "skuId": "Mobile-0001",
            "skuName": "xiaomi",
            "unitPrice": 59000000,
            "qty": 1,
            "pictureUrl": "https://d12x8ezp3au6m3.cloudfront.net/item/oh0ymJAba3Ftg2NiGa1-pLD0MXo-KYUeHujlFwu79-4.png?w=526&h=526&fmt=webp&q=85&fit=0",
            "seller": {
                "sellerId": "Seller-0001",
                "sellerName": "Xiaomi-Seller"
            }
        },
        {
            "categoryId": "Category-0002",
            "categoryName": "PhoneCase",
            "skuId": "PhoneCase-0001",
            "skuName": "iphone13-case",
            "unitPrice": 6500000,
            "qty": 2,
            "pictureUrl": "https://d12x8ezp3au6m3.cloudfront.net/item/oh0ymJAba3Ftg2NiGa1-pLD0MXo-KYUeHujlFwu79-4.png?w=526&h=526&fmt=webp&q=85&fit=0",
            "seller": {
                "sellerId": "Seller-0002",
                "sellerName": "PhoneCase-Seller"
            }
        }
    ]
}`
	var req akucicil.InstallmentPlanReq
	json.Unmarshal([]byte(jsonData), &req)

	resp, errResp := newAc.GetInstallmentPlan(req)
	if errResp != nil{
		println(errResp)
	}

	fmt.Printf("%+v\n", resp)

}

func getCheckoutOrder(newAc akucicil.AkucicilService) {
	jsonData := `{
    "partnerOrderId": "REF-8812324",
    "partnerOrderAmount": 72000000,
    "currency": "IDR",
    "redirectUrl": "https://www.your-redirect-url.com",
    "callbackUrl": "https://www.callback-url.com",
    "expireTime": "2022-02-07T11:09:00+00:00",
    "buyer": {
        "buyerId": "Buyer-0001",
        "buyerPhone": "+6281212345678",
        "buyerEmail": "buyer@email.com",
        "registrationDays": "180"
    },
    "shipping": {
        "receiverName": "Test",
        "receiverPhone": "+6281212345678",
        "region": "Bandung",
        "city": "Jawa Barat",
        "district": "Bandung",
        "street": "Jl Jend Sudirman 297, Jawa Barat",
        "postcode": "40241",
        "addressLine": "Akulaku Office",
        "countryCode": "IDN"
    },
    "items": [
        {
            "categoryId": "Category-0001",
            "categoryName": "Mobile",
            "skuId": "Mobile-0001",
            "skuName": "xiaomi",
            "unitPrice": 59000000,
            "qty": 1,
            "pictureUrl": "",
            "seller": {
                "sellerId": "Seller-0001",
                "sellerName": "Xiaomi-Seller"
            }
        },
        {
            "categoryId": "Category-0002",
            "categoryName": "PhoneCase",
            "skuId": "PhoneCase-0001",
            "skuName": "iphone13-case",
            "unitPrice": 6500000,
            "qty": 2,
            "pictureUrl": "",
            "seller": {
                "sellerId": "Seller-0002",
                "sellerName": "PhoneCase-Seller"
            }
        }
    ],
    "metadata": "",
    "extendInfo": ""
}`
	var req akucicil.CheckoutOrderReq
	json.Unmarshal([]byte(jsonData), &req)

	resp, errResp := newAc.CheckoutOrder(req)
	if errResp != nil{
		println(errResp)
	}

	fmt.Printf("%+v\n", resp)

}