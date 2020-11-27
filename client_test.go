package alipay

import (
	"testing"
	"time"
)

const (
	privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEApemkXbTCTEFTXg0dwyPP+RKcxF3gTbLdm8O2awgOnIMpYAqh2fQUeSGnWgwD/6Ughj/8e+S+XXJ1dV+SAhj/OL7CGAoGPuN6rqrH7yPXzO3AGVTd+30DEd6XMvw6Ju5xhHl12Lc7tW8mCmF7wuZUQihlulKewfKRHOOEvQWC06n6X5drnN+sMAr0SmN8bc0prIk0nW4fNZwmqiRGazml4zflDJ4Sndz47CXlS0s3mW87M6T0GThmwDiSQUQwQQDf1w2cvQ6GkAYuHzK4GHB6L6BmVNS7MFrYrE6rkH4ceP6QHrL5Buhf+O/nnavmlSUFGhRm3BWe7c9Bt61PrlYWOQIDAQABAoIBAHjzoGjT2wW+ZeldxIG7POWGRRT+nwPlzpq8jeLvR7+f+uzSM3Xx827vMtJ5ify3w8M7KHSlqIX1aF2943J2CLG0l0jxHeaA7bIiPIlA5xS1imKtNPsfArrnO/DmYfp5v/Xkmh34TqYRNnlA4fmO8oQccTTpAGXB0TpvPxiRyPNGfYycUM+t+7jaFh9shxJXgSV5PCujk4pcoPuLj7Yl81jcrOzWazEXgs22L+qBdyXU3sEdqLtIVR4j74NQHSFu+L711GJK9ZeSSRbgxG86q5eYZgYiT2ti9SJ+HAiurjQy/X+2PFSC1PqrJvsrMhMcd9VqslKoBFUe203ot1horgECgYEA2k8aB71m5Q1/opuxnF63hQCgCzbFV+5Wf/1Gee14kkZXLL6y4xOlzJqQ8GhNE/cEVr//qTz9OzQsnoDxvNtAwmJZ9DyrJCKJV+UFduEHb7M5hLhXvj6GOu+JNm0hwBbf4SnvAeaVPoZr1w1WKiL11/ZeUl2F1ORROFM0USO4FfkCgYEAwo61e69XBjUfSWE422XhHooF+C5XH5qPN7oVDSwEKhCVTzSSi13QQLt5/vj2+4Ks4ZRuMdD1A0ArJkyn8bsVw+5UmJegALI2ga3NfwdqhrGON9PhOg/s+MxMU6D/JPgZ11kYJ1BstqvXxHbs3RTTMoJRkQSGt8EK8JBO5LK1EkECgYBCakI/DI4bLSohbEBylBY87l3CS51qDOZf9cvVGDvQNHoc8L83eii8wGFL4k9gvYuiYLME57saodrZNd1VWVawTH+VYEeorKEgDlrFOdyrTNk42WRISnlHwMv7tOPJrqvZsoo2B9JrvTVdrX8DPrOQSjGT2UP36qYS/q0x6i388QKBgQC8AK0sbN36GKE4BmOr1sH4AcYM8bKszmzwm7c1+D+56jZtyE6Hr8rKkp8rnKcFmVu3y/fD2bi5QGux4cc0FuXMZGSI45PwuEVlgG6f/qmYqMDV/7+XnMYQEVL8SQnkTn6iEuz9KIE0789bgNQYOsRu/XEWYjpQHJGWrswdwPaqQQKBgAIddw9MMPIxcDYE0XF6HDT9gm+PXi8ZrgIUkxEt74EemN1Q+3Z7BZerwOsDSL+d32zB954iWVeZCUDachUgevqG9xXMaMF0FdaRrVcLzgkZYgOEcwstN3vjYMnVYpkDxQqe+P9FR3Qj37QqbCKQiqOF99igIVH4d4VXfX6D5azJ
-----END RSA PRIVATE KEY-----
`
	gateway = "https://openapi.alipaydev.com/gateway.do"
	appId   = "2016102500757322"
)

var client *Client

func TestNew(t *testing.T) {
	client = New(gateway, appId, privateKey)
	if client == nil {
		t.Fatal("test new failed")
	}
	t.Log("test new failed")
}

func TestWithTimeLocation(t *testing.T) {
	c := New(gateway, appId, privateKey, WithTimeLocation(time.UTC))
	if c.location != time.UTC {
		t.Fatal("set location failed")
	}

	t.Log("set location success")
}
