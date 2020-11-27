package alipay

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTrade_Page(t *testing.T) {
	var biz = BizContent{}
	biz.Subject = "测试PC网页支付"
	biz.OutTradeNo = fmt.Sprintf("test%d", rand.Int63n(time.Now().UnixNano()))
	biz.TotalAmount = "8.88"
	biz.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.Page(biz)
	if err != nil {
		t.Fatalf("test wap failed, err:%v\n", err)
	}

	fmt.Println(url)
	t.Log(url)
}

func TestTrade_Wap(t *testing.T) {
	var biz = BizContent{}
	biz.Subject = "测试手机网页支付"
	biz.OutTradeNo = fmt.Sprintf("test%d", rand.Int63n(time.Now().UnixNano()))
	biz.TotalAmount = "8.88"
	biz.ProductCode = "QUICK_WAP_WAY"

	url, err := client.Wap(biz)
	if err != nil {
		t.Fatalf("test wap failed, err:%v\n", err)
	}

	fmt.Println(url)
	t.Log(url)
}

func TestTrade_App(t *testing.T) {

}
