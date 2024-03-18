package lib

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

func GetCommonHeaders() map[string]string {
	return map[string]string{
		"authority":          "api-one-wscn.awtmt.com",
		"accept":             "*/*",
		"accept-language":    "zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,zh-TW;q=0.6",
		"cache-control":      "no-cache",
		"dnt":                "1",
		"origin":             "https://wallstreetcn.com",
		"pragma":             "no-cache",
		"referer":            "https://wallstreetcn.com/",
		"sec-ch-ua":          `Chromium";v="122":"Not(A:Brand";v="24":"Google Chrome";v="122`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "macOS",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "cross-site",
		"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		"x-client-type":      "pc",
		"x-device-id":        "18e32f68-d43b-53a2-ac6f-4d748aabf314",
		"x-ivanka-app":       "wscn|web|0.38.142|0.0|0",
		"x-ivanka-platform":  "wscn-platform",
		"x-taotie-device-id": "18e32f68-d43b-53a2-ac6f-4d748aabf314",
		"x-track-info":       `{"appId":"com.wallstreetcn.web","appVersion":"0.38.142"}`,
	}
}

func GetHttpClient() *resty.Client {

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(4 * time.Second)

	client.SetHeaders(GetCommonHeaders())
	client.SetContentLength(true)

	client.
		SetRetryCount(3).
		SetRetryWaitTime(3 * time.Second).
		SetDebug(false)

	return client
}

func RequestToStruct[T any](url string, queryParams map[string]string) (t *T, err error) {
	// 发送请求并获取字节
	bytes, err := RequestBytes(url, queryParams)
	if err != nil {
		// 如果有错误发生，返回 nil 和错误
		return nil, err
	}
	fmt.Println(url)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &t)
	fmt.Println(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func RequestBytes(url string, queryParams map[string]string) ([]byte, error) {
	resp, err := GetHttpClient().R().
		SetQueryParams(queryParams).
		//EnableTrace().
		Get(url)

	if err != nil {
		return nil, err
	}
	//fmt.Println("Response Info:")
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())

	if resp.StatusCode() == http.StatusOK {
		return resp.Body(), nil
	} else {
		panic(fmt.Sprintf("错误号码: url is %s ,status code is %d %s", url, resp.StatusCode(), string(resp.Body())))
		return nil, err
	}
}
