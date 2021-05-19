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
	ApiId  string
	ApiURL string
}

// GetSMSInstance return sms service instance
func GetSMSInstance(APIID, APIURL string) *SMSru {
	s := SMSru{ApiId: APIID, ApiURL: APIURL}
	return &s
}

// SendSMS main function for sms send
func (sms *SMSru) SendSMS(number, text string) (SendResponse, error) {
	result := SendResponse{}

	if sms.ApiId == "" && sms.ApiURL == "" {
		return result, errors.New("Not found setting for sms request")
	}
	//https://sms.ru/sms/send?api_id=18B2616D-A8BC-EFE3-0509-9581ABDDEA5F&to=79186313258,74993221627&msg=hello+world&json=1
	params := fmt.Sprintf("%s?api_id=%s&to=%s&msg=%s&json=1", sms.ApiURL, sms.ApiId, number, text)

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
	// var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	log.Printf("SMS send response %s", result.Status)
	return result, nil
}
