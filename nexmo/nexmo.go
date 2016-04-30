package nexmo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://rest.nexmo.com"
)

var (
	ApiKey    string
	ApiSecret string
	n         *Nexmo
)

type Nexmo struct {
	queue  chan query
	client *http.Client
}

type query struct {
	url        string
	params     url.Values
	data       interface{}
	responseCh chan response
}

type response struct {
	data interface{}
	err  error
}

func Init() {
	queue := make(chan query)

	n = &Nexmo{
		queue:  queue,
		client: http.DefaultClient,
	}
	go n.start()
}

func (n *Nexmo) start() {
	for q := range n.queue {
		url := q.url
		params := q.params
		data := q.data
		responseCh := q.responseCh

		params.Set("api_key", ApiKey)
		params.Set("api_secret", ApiSecret)

		err := n.execQuery(url, params, data)

		responseCh <- response{data, err}
	}
}

func (n *Nexmo) execQuery(url string, params url.Values, data interface{}) error {
	rsp, err := n.client.Get(fmt.Sprintf("%s?%s", url, params.Encode()))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	return json.NewDecoder(rsp.Body).Decode(data) // BUG: Should be able to marshal keys with dashes e.g. message-count
}

func SendSMS(args url.Values, rsp interface{}) error {
	responseCh := make(chan response)

	n.queue <- query{fmt.Sprintf("%s/sms/json", baseURL), args, &rsp, responseCh}

	err := (<-responseCh).err
	if err != nil {
		return err
	}

	return nil
}
