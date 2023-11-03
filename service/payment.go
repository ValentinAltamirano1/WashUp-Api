package service

import (
	"github.com/fabianMendez/mercadopago"
	"fmt"
	"bytes"
    "encoding/json"
    "net/http"
)

type PaymentParams struct {
	Title string `json:"title"`
	Price float64 `json:"price"`
}

type PreferenceHandler struct {
	mpClient *mercadopago.Client
}

func NewPreferenceHandler() *PreferenceHandler {
	client := mercadopago.NewClient("https://api.mercadopago.com/v1", "TEST-76053192-d831-4b11-8281-3270e7a283f1", "TEST-3705400319827255-110219-cc509cb730e4aa70fb23f32cae71acee-1532920685")
	return &PreferenceHandler{mpClient: &client}
}

func (h *PreferenceHandler) CreatePreference(params PaymentParams) (string, error) {
    preferenceData := map[string]interface{}{
        "items": []map[string]interface{}{
            {
                "title":       params.Title,
                "quantity":    1,
                "currency_id": "USD",
                "unit_price":  params.Price,
            },
        },
    }

    data, err := json.Marshal(preferenceData)
    if err != nil {
        return "", err
    }

    url := "https://api.mercadopago.com/checkout/preferences?access_token=TEST-3705400319827255-110219-cc509cb730e4aa70fb23f32cae71acee-1532920685"
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
	fmt.Println(resp)
	fmt.Println(resp.StatusCode)

    if resp.StatusCode != http.StatusCreated {
        return "", fmt.Errorf("failed to create preference, status code: %d", resp.StatusCode)
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    initPoint, ok := result["init_point"].(string)
    if !ok {
        return "", fmt.Errorf("init_point not found in MercadoPago response")
    }
	fmt.Println(initPoint)
    return initPoint, nil
}