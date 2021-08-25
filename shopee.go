package cekresi

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
)

type statusList struct {
	Timestamp int    `json:"timestamp"`
	Code      int    `json:"code"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
}
type trackingList struct {
	Status    string `json:"status"`
	Timestamp int    `json:"timestamp"`
	Message   string `json:"message"`
}
type responseData struct {
	DeliveryType      string         `json:"delivery_type"`
	CurrentStatus     string         `json:"current_status"`
	TrackingList      []trackingList `json:"tracking_list"`
	Phone             string         `json:"phone"`
	RecipientName     string         `json:"recipient_name"`
	SlsTrackingNumber string         `json:"sls_tracking_number"`
	StatusList        []statusList   `json:"status_list"`
}

type response struct {
	Message string       `json:"message"`
	Data    responseData `json:"data"`
	Retcode int          `json:"retcode"`
}

func trackingNumber(resi string) string {
	k := "MGViZmZmZTYzZDJhNDgxY2Y1N2ZlN2Q1ZWJkYzlmZDY="
	r := math.Floor(float64(time.Now().UnixNano() / int64(time.Millisecond) / 1e3))
	h := sha256.New()
	rs := fmt.Sprintf("%d", int64(r))
	h.Write([]byte(resi + rs + k))
	return fmt.Sprintf(resi+"|"+rs+"%x", h.Sum(nil))
}

func request(resi string, trackingNumber string) (response, error) {
	var data response
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://spx.co.id/api/v2/fleet_order/tracking/search?sls_tracking_number="+trackingNumber, nil)
	if err != nil {
		return data, err
	}
	req.Header.Set("Authority", "spx.co.id")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://spx.co.id/detail/"+resi)
	req.Header.Set("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Cookie", "fms_language=id; _ga=GA1.3.767969781.1629894593; _gid=GA1.3.339992191.1629894593")

	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func Shopee(resi string) (response, error) {
	tn := trackingNumber(resi)
	return request(resi, tn)
}
