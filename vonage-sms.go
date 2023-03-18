// Package vonagesms Envia sms usando o servico vonage
package vonagesms

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const ApiSmsURL = "https://rest.nexmo.com/sms/json"

type SmsPayload struct {
	ApiSecret string `json:"api_secret"`
	ApiKey    string `json:"api_key"`
	From      string `json:"from"`
	To        string `json:"to"`
	Msg       string `json:"text"`
}

type SmsResponseData struct {
	To               string `json:"to"`
	MessageID        string `json:"message-id"`
	Status           string `json:"status"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
	ErrorText        string `json:"error-text"`
}

type SmsResponse struct {
	Messages     []SmsResponseData `json:"messages"`
	MessageCount string            `json:"message-count"`
}

type VonageSmsRepository interface {
	SendSms(phoneNumber, message string) (response SmsResponse, err error)
}

type VonageSms struct {
	vonageApiKey    string
	vonageApiSecret string
	vonageBrandName string
}

// VonageClient instancia um novo cliente
func VonageClient(apiKey, apiSecret, brandName string) VonageSmsRepository {
	return &VonageSms{
		vonageApiSecret: apiSecret,
		vonageApiKey:    apiKey,
		vonageBrandName: brandName,
	}
}

func (vs VonageSms) SendSms(phoneNumber, message string) (responseData SmsResponse, err error) {

	payload := SmsPayload{
		ApiSecret: vs.vonageApiSecret,
		ApiKey:    vs.vonageApiKey,
		From:      vs.vonageBrandName,
		To:        phoneNumber,
		Msg:       message,
	}

	var body io.Reader

	payloadData, _ := json.Marshal(payload)

	body = bytes.NewBuffer(payloadData)

	request, err := http.NewRequest(http.MethodPost, ApiSmsURL, body)

	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 20,
	}

	response, err := client.Do(request)

	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		err = NewError("")
		return

	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &responseData)

	if responseData.Messages[0].Status == "0" {
		return
	}

	err = NewError(responseData.Messages[0].Status)

	return
}
