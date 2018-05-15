package bitgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type BitGo struct {
	host  string
	token string
	coin  string
}

type ListParams struct {
	Limit     int    `url:"limit,omitempty"`
	PrevId    string `url:"prevId,omitempty"`
	AllTokens bool   `url:"allTokens,omitempty"`
}

func New(env string, token string) *BitGo {
	switch env {
	case "test":
		env = "https://test.bitgo.com/api/v2"
	case "prod":
		env = "https://www.bitgo.com/api/v2"
	}
	return &BitGo{
		host:  env,
		token: token,
	}
}

func (b *BitGo) clone() *BitGo {
	return &BitGo{
		host:  b.host,
		token: b.token,
		coin:  b.coin,
	}
}

func (b *BitGo) Coin(coin string) *BitGo {
	c := b.clone()
	c.coin = coin
	return c
}

func (b *BitGo) get(url string, params interface{}, responce interface{}) (err error) {
	if params != nil {
		v, _ := query.Values(params)
		url = url + "?" + v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, b.host+"/"+url, nil)
	if err != nil {
		return
	}

	return b.request(req, responce)
}

func (b *BitGo) modify(method string, url string, params interface{}, responce interface{}) (err error) {
	var body *bytes.Buffer

	if params != nil {
		body = new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(params)
		if err != nil {
			return
		}
	}

	req, err := http.NewRequest(method, b.host+"/"+url, body)
	if err != nil {
		return
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return b.request(req, responce)
}

func (b *BitGo) post(url string, params interface{}, responce interface{}) (err error) {
	return b.modify(http.MethodPost, url, params, responce)
}

func (b *BitGo) put(url string, params interface{}, responce interface{}) (err error) {
	return b.modify(http.MethodPut, url, params, responce)
}

func (b *BitGo) request(req *http.Request, responce interface{}) (err error) {
	req.Header.Add("Authorization", "Bearer "+b.token)

	client := &http.Client{
		Timeout: time.Minute,
	}

	r, err := client.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	if r.StatusCode != http.StatusOK {
		berr := Error{}

		err = json.Unmarshal(resp, &berr)
		if err != nil {
			return err
		}

		return berr
	}

	err = json.Unmarshal(resp, &responce)
	return
}
