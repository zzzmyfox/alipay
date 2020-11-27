package alipay

import "net/url"

// ITrade 定义所有支付方式的接口
type ITrade interface {
	Page(v interface{}) (string, error)
	App(v interface{}) (string, error)
	Wap(v interface{}) (string, error)
}

// Trade 用来配置所有支付方式
type Trade struct {
	client *Client
}

var _ ITrade = &Trade{}

// Page  页面支付
func (trade *Trade) Page(v interface{}) (string, error) {
	method := "alipay.trade.page.pay"

	param, err := trade.client.format(v, method)
	if err != nil {
		return "", nil
	}

	var result *url.URL
	result, err = url.Parse(trade.client.gateway + "?" + param.Encode())
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

// App App 支付
func (trade Trade) App(v interface{}) (string, error) {
	method := "alipay.trade.app.pay"

	param, err := trade.client.format(v, method)
	if err != nil {
		return "", nil
	}
	return param.Encode(), nil
}

// Wap 手机网站支付接口
func (trade Trade) Wap(v interface{}) (string, error) {
	method := "alipay.trade.wap.pay"

	param, err := trade.client.format(v, method)
	if err != nil {
		return "", nil
	}

	var result *url.URL
	result, err = url.Parse(trade.client.gateway + "?" + param.Encode())
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
