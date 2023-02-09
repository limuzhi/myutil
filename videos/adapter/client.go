package adapter

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func NewClient(proxys ...string) *resty.Client {
	client := resty.New()
	if len(proxys) > 0 {
		client.SetProxy(proxys[0])
	}
	client.SetTimeout(time.Second * 5)
	return client
}

func NewMobileHttpRequest(proxys ...string) *resty.Request {
	client := NewClient(proxys...)
	client.SetHeader("User-Agent", GetRandomUserAgent())
	return client.R()
}
