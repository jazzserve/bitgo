package bitgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type BitGo struct {
	Host  string
	Token string

	coin string
}

type ListParams struct {
	Limit     int    `url:"limit,omitempty"`
	PrevId    string `url:"prevId,omitempty"`
	AllTokens bool   `url:"allTokens,omitempty"`
}

func (b *BitGo) clone() *BitGo {
	return &BitGo{
		Host:  b.Host,
		Token: b.Token,
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

	req, err := http.NewRequest(http.MethodGet, b.Host+"/"+url, nil)
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

	req, err := http.NewRequest(method, b.Host+"/"+url, body)
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
	req.Header.Add("Authorization", "Bearer "+b.Token)

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
		err = errors.New(string(resp))
		return
	}

	err = json.Unmarshal(resp, &responce)
	return
}
