/**
 @author:way
 @date:2021/12/16
 @note
**/

package baby

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	cmap "github.com/orcaman/concurrent-map"
	"redisData/dao/redis"
	"redisData/model"
	"redisData/pkg/logger"
	"redisData/pkg/mysql"
)


type Controller struct {
}

//BuyAndSellHandler  买卖数据
func (h *Controller) BuyAndSellHandler(c *gin.Context) {
	var (
		params     model.ParamsBuyAndSellQuery // 接收请求参数
		BuyAndSell []model.RespBabyOrder       // 查询数据
	)
	// 绑定参数
	_ = c.Bind(&params)
	where := cmap.New().Items()
	if len(params.Name) > 0 {
		where["name"] = params.Name
	}
	if len(params.TokenId) > 0 {
		where["token_id"] = params.TokenId
	}
	if len(params.Status) > 0 {
		where["status"] = params.Status
	}
	mysql.DB.Debug().Model(model.BabyOrder{}).Where(where).Find(&BuyAndSell)
	ResponseSuccess(c, BuyAndSell)
}

//SetBabyRiskHandle 设置baby风控
func SetBabyRiskHandle(c *gin.Context) {
	//获取参数
	var p model.ParamRiskMng
	err := c.Bind(&p)
	if err != nil {
		logger.Info(err)
		return
	}
	if p.Situation == "" || p.TimeLevel == 0 || p.Percentage == 0 || p.OperationType == 0 || p.Status == 0 {
		ResponseError(c,500)
		return
	}
	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["Situation"] = p.Situation
	m["TimeLevel"] = p.TimeLevel
	m["Percentage"] = p.Percentage
	m["OperationType"] = p.OperationType
	m["Status"] = p.Status
	logger.Info(m)
	redis.CreatHashKey(fmt.Sprintf("baby:ConfigRisk:%s", p.Situation), m)
	//返回参数
	ResponseSuccess(c,"")

}

//SetBabyOnOffHandle 设置baby全自动和半自动
func SetBabyOnOffHandle(c *gin.Context) {
	//获取参数
	var p model.ParamOnOff
	err := c.Bind(&p)
	if err != nil {
		logger.Info(err)
		ResponseError(c,501)
		return
	}

	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["CrlName"] = p.CrlName
	m["Super"] = p.Super
	logger.Info(m)
	redis.CreatHashKey(fmt.Sprintf("baby:ConfigStopAuto:%s", p.CrlName), m)
	//返回参数
	ResponseSuccess(c,"")
}

//SetBabyBuyConfHandle 设置买入出参数
func SetBabyBuyConfHandle(c *gin.Context) {
	var p model.ParamBuyAndSaleSet
	//var reps model.RespBuyAndSaleSet
	Eerr := c.Bind(&p)
	if Eerr != nil {
		logger.Error(Eerr)
		return
	}
	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["percent"] = p.Percent
	m["status"] = p.Status
	m["types"] = p.Types
	m["market_price"] = p.MarketPrice
	logger.Info(m)
	redis.CreatHashKey("baby:ConfigBuy", m)
	//返回参数
	ResponseSuccess(c,200)
}

//SetBabySaleConfHandle 设置卖出参数
func SetBabySaleConfHandle(c *gin.Context) {
	var p model.ParamBuyAndSaleSet
	Eerr := c.Bind(&p)
	if Eerr != nil {
		logger.Error(Eerr)
		return
	}
	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["percent"] = p.Percent
	m["status"] = p.Status
	m["types"] = p.Types
	m["market_price"] = p.MarketPrice
	redis.CreatHashKey("baby:ConfigSale", m)
	//返回参数
	ResponseSuccess(c,200)
}

//GetBabyScriptStatusHandle 获取脚本运行的状态
func GetBabyScriptStatusHandle(c *gin.Context) {
	//获取买入卖出总开关
	var buyStruct model.RespAllOnOff
	buy := redis.GetHashDataAll("baby:ConfigStopAuto:buy")
	mapstructure.Decode(buy, &buyStruct)

	var saleStruct model.RespAllOnOff
	sale := redis.GetHashDataAll("baby:ConfigStopAuto:sale")
	mapstructure.Decode(sale, &saleStruct)

	//通过reids获取市场价格
	babyMarket,_ := redis.GetData("baby:marketPrice")
	//获取baby买入数据
	var babyBuy model.RespBuyAndSaleSet
	babyBuy2 := redis.GetHashDataAll("baby:ConfigBuy")
	err := mapstructure.Decode(babyBuy2, &babyBuy)
	//全自动提供市场运算的数据
	babyBuy.AotuMarketprice = babyMarket
	if err != nil {
		logger.Error(err)
		return
	}
	//获取baby卖出入数据
	var babySale model.RespBuyAndSaleSet
	baby_sale := redis.GetHashDataAll("baby:ConfigSale")
	mapstructure.Decode(baby_sale, &babySale)
	babySale.AotuMarketprice = babyMarket

	var all model.RespAllSwitch
	allOnOffSlice := make([]model.RespAllOnOff, 2)
	buyAndSaleSetSlice := make([]model.RespBuyAndSaleSet, 4)

	buyAndSaleSetSlice[0] = babyBuy
	buyAndSaleSetSlice[1] = babySale

	allOnOffSlice[0] = buyStruct
	allOnOffSlice[1] = saleStruct

	all.AllOnOff = allOnOffSlice
	all.BuyAndSale = buyAndSaleSetSlice

	//获取买出率
	var resp model.RespSellingRate
	sellRate := redis.GetHashDataAll("baby:ConfigSaleRate")
	mapstructure.Decode(sellRate, &resp)
	ResponseSuccess(c,all)

}

//SetBabySaleRateHandle 设置卖出率
func SetBabySaleRateHandle(c *gin.Context) {
	var p model.ParamSellingRate
	BErr := c.Bind(&p)
	if BErr != nil {
		logger.Error(BErr)
		return
	}
	m := make(map[string]interface{})
	m["time_level"] = p.TimeLevel
	m["percent"] = p.Percent
	m["status"] = p.Status
	m["operation_type"] = p.OperationType
	logger.Info(m)
	redis.CreatHashKey("baby:ConfigSaleRate", m)
	ResponseSuccess(c,"")
}