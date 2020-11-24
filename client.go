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

// Client 支付客户端结构体
type Client struct {
	appId string

	gateway  string
	location *time.Location

	appPrivateKey   string
	alipayPublicKey string

	alipayPublicKeyCert string
	alipayRootCert      string
}

// OptionFunc 定义配置选项option，option是一个func，入参是*Client实例，在里面我们可以修改实例的值。
type OptionFunc func(*Client)

// WithTimeLocation 设置时区，在初始化客户端的作为入参调用
func WithTimeLocation(location *time.Location) OptionFunc {
	return func(c *Client) {
		c.location = location
	}
}

// New 初始化支付宝客户端
func New(gateway, appId, privateKey string, opts ...OptionFunc) *Client {
	client := &Client{
		appId:         appId,
		gateway:       gateway,
		location:      time.Local,
		appPrivateKey: privateKey,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *Client) Pay(subject, outTradeNo, totalAmount string) (string, error) {
	v := Values{}
	v.Add("app_id", c.appId)
	v.Add("method", "alipay.trade.wap.pay")
	v.Add("charset", "utf-8")
	v.Add("format", "JSON")
	v.Add("sign_type", "RSA2")
	v.Add("timestamp", c.timestamp())
	v.Add("version", "1.0")

	biz := BizContent{
		OutTradeNo:  outTradeNo,
		ProductId:   "FAST_INSTANT_TRADE_PAY",
		TotalAmount: totalAmount,
		Subject:     subject,
		Body:        subject,
	}

	bytes, err := json.Marshal(biz)
	if err != nil {
		return "", err
	}

	v.Add("biz_content", string(bytes))

	sign, err := c.signature(v.Params())
	if err != nil {
		return "", err
	}

	v.Add("sign", url.QueryEscape(sign))

	res, err := url.Parse(c.gateway + "?" + v.Params())
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

// AppPublicCert
func (c *Client) AppPublicCert(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	c.alipayPublicKeyCert = string(bytes)

	return nil
}

// AlipayPublicCert
func (c *Client) AlipayPublicCert() {

}

func (c *Client) timestamp() string {
	return time.Now().In(c.location).Format("2006-01-02 15:04:05")
}

func (c *Client) do(str string) (body []byte, err error) {
	method := "POST"
	payload := strings.NewReader(str)
	var (
		req  *http.Request
		resp *http.Response
	)

	req, err = http.NewRequest(method, c.gateway, payload)
	if err != nil {
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	client := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) signature(values string) (sign string, err error) {
	privateKey, err := parsePKCS1PrivateKey(c.appPrivateKey)
	if err != nil {
		privateKey, err = parsePKCS8PrivateKey(c.appPrivateKey)
		if err != nil {
			return
		}
	}

	var bytes []byte
	bytes, err = signPKCS1v15([]byte(values), privateKey)
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(bytes)
	return
}

type BizContent struct {
	OutTradeNo  string `json:"out_trade_no"`
	ProductId   string `json:"product_id"`
	TotalAmount string `json:"total_amount"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}

// Values
type Values map[string][]string

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

// Params
func (v Values) Params() string {
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
