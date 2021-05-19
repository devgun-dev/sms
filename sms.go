package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// SendResponse reply after sms send
type SendResponse struct {
	Status     string   `json:"status"`
	StatusCode int      `json:"status_code"`
	SMS        []Number `json:"sms"`
	Balance    float32  `json:"balance"`
}

// Number operation status for each number in sendlist
type Number struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	SmsID      string `json:"sms_id"`
}

// SMSru main instance of service sms.ru
type SMSru struct {
	API_ID  string
	API_URL string
}

// GetSMSInstance return sms service instance
func GetSMSInstance(APIID, APIURL string) *SMSru {
	s := SMSru{API_ID: APIID, API_URL: APIURL}
	return &s
}

// SendSMS main function for sms send
func (sms *SMSru) SendSMS(number, text string) (SendResponse, error) {
	result := SendResponse{}

	if sms.API_ID == "" && sms.API_URL == "" {
		return result, errors.New("Not found setting for sms request")
	}

	params := fmt.Sprintf("%s?api_id=%s&to=%s&msg=%s&json=1", sms.API_URL, sms.API_ID, number, text)

	client := http.Client{}
	request, err := http.NewRequest("GET", params, nil)
	if err != nil {
		return result, err
	}

	r, err := client.Do(request)
	if err != nil {
		return result, err
	}
	log.Printf("SMS send response to number %s\n", number)

	json.NewDecoder(r.Body).Decode(&result)
	log.Printf("SMS send response %s", result.Status)
	return result, nil
}
