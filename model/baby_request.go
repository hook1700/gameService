/**
 @author:way
 @date:2021/12/16
 @note
**/

package model

//  买卖数据

type ParamsBuyAndSellQuery struct {
	Name    string `json:"name" form:"name"`         // 名称
	TokenId string `json:"token_id" form:"token_id"` // token
	Status  string `json:"status" form:"status"`     // 状态 1.买入
}

// 买入卖出后，前端返回交易hash

type ParamsUpdateBuyAndSale struct {
	Id          int     `json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`                 //名称
	FixPrice    float64 `json:"fix_price" form:"fix_price"`       //单价
	SalePrice   float64 `json:"sale_price" form:"sale_price"`     //出售价格
	Profit      float64 `json:"profit" form:"profit"`             //利润
	Status      int     `json:"status" form:"status"`             //状态 1.买入 2.买出
	TokenId     string  `json:"token_id" form:"token_id"`         //token
	MarketPrice float64 `json:"market_price" form:"market_price"` //买入市场价
	TxHash      string  `json:"tx_hash" form:"tx_hash"`           //接收交易hash
}

//ParamSetPrivateKey 设置私钥请求参数
type ParamSetPrivateKey struct {
	PrivateKey string `json:"private_key" form:"private_key"`
}