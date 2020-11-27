package alipay

type BizContent struct {
	OutTradeNo  string         `json:"out_trade_no"`
	ProductCode string         `json:"product_code"`
	TotalAmount string         `json:"total_amount"`
	Subject     string         `json:"subject"`
	Body        string         `json:"body"`
	GoodsDetail []*GoodsDetail `json:"goods_detail"`
}

type GoodsDetail struct {
	GoodsId        string  `json:"goods_id"`
	AliPayGoodsId  string  `json:"alipay_goods_id,omitempty"`
	GoodsName      string  `json:"goods_name"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
	GoodsCategory  string  `json:"goods_category,omitempty"`
	CategoriesTree string  `json:"categories_tree,omitempty"`
	Body           string  `json:"body,omitempty"`
	ShowURL        string  `json:"show_url,omitempty"`
}
