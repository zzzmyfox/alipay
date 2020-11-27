package alipay

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// OptionFunc 定义客户端配置选项
type OptionFunc func(*Client)

// WithTimeLocation 修改时区，默认为服务器本地时区
func WithTimeLocation(location *time.Location) OptionFunc {
	return func(c *Client) {
		c.location = location
	}
}

// Client 支付宝客户端
type Client struct {
	Trade

	appId string

	gateway  string
	location *time.Location

	appPrivateKey   string
	alipayPublicKey string

	alipayPublicKeyCert string
	alipayRootCert      string
}

var _ ITrade = &Client{}

// New 初始化支付宝客户端
func New(gateway, appId, privateKey string, opts ...OptionFunc) *Client {
	client := &Client{
		Trade:         Trade{},
		appId:         appId,
		gateway:       gateway,
		location:      time.Local,
		appPrivateKey: privateKey,
	}

	client.Trade.client = client

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (client *Client) format(v interface{}, method string) (url.Values, error) {
	values := url.Values{}
	values.Add("app_id", client.appId)
	values.Add("method", method)
	values.Add("charset", "utf-8")
	values.Add("format", "JSON")
	values.Add("sign_type", "RSA2")
	values.Add("timestamp", client.getTimestamp())
	values.Add("version", "1.0")

	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	values.Add("biz_content", string(bytes))

	urlValue := client.getUrlValue(values)

	var sign string
	sign, err = client.getSignature(urlValue)
	if err != nil {
		return nil, err
	}

	values.Add("sign", sign)

	return values, nil
}

func (client *Client) getTimestamp() string {
	return time.Now().In(client.location).Format("2006-01-02 15:04:05")
}

func (client *Client) getSignature(src string) (string, error) {
	privateKey, err := parsePKCS1PrivateKey(client.appPrivateKey)
	if err != nil {
		privateKey, err = parsePKCS8PrivateKey(client.appPrivateKey)
		if err != nil {
			return "", err
		}
	}

	var bytes []byte
	bytes, err = signPKCS1v15([]byte(src), privateKey)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func (client *Client) do(method, str string) ([]byte, error) {
	payload := strings.NewReader(str)
	req, err := http.NewRequest(method, client.gateway, payload)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	c := http.Client{
		Timeout: time.Second * 5,
	}

	resp := &http.Response{}
	resp, err = c.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	return ioutil.ReadAll(resp.Body)
}

func (client *Client) getUrlValue(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}
