package common

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func RequestJsonBufByPost(fullUrl string, jsonPayload []byte, proxyes ...string) ([]byte, error) {
	hc := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	if len(proxyes) > 0 {
		proxyUrl0 := proxyes[0]
		if proxyUrl0 != "" {
			tr := &http.Transport{TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}}
			proxyUrl, err := url.Parse(proxyUrl0)
			if err == nil { // 使用传入代理
				tr.Proxy = http.ProxyURL(proxyUrl)
			}
			hc.Transport = tr
		}
	}
	resp, err := hc.Post(fullUrl, "application/json", bytes.NewReader(jsonPayload))
	return handleHttpResp(resp, err)
}

func handleHttpResp(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status=%d,body=%s", resp.StatusCode, payload)
	}
	return payload, nil
}

func RequestJsonBufByGet(fullUrl string, proxyes ...string) ([]byte, error) {
	hc := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	if len(proxyes) > 0 {
		proxyUrl0 := proxyes[0]
		if proxyUrl0 != "" {
			tr := &http.Transport{TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}}
			proxyUrl, err := url.Parse(proxyUrl0)
			if err == nil {
				tr.Proxy = http.ProxyURL(proxyUrl)
			}
			hc.Transport = tr
		}
	}
	resp, err := hc.Get(fullUrl)
	return handleHttpResp(resp, err)
}
