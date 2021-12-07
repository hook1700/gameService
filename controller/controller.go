/**
 @author:way
 @date:2021/12/3
 @note
**/

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"redisData/dao/mysql"
	"redisData/dao/redis"
	"redisData/logic"
	"redisData/model"
	"redisData/pkg/logger"
	"strconv"
)

// GetDataHandle 返回id详情列表前端用
func GetDataHandle(c *gin.Context)  {
	//获取参数
	dataType := c.Query("dataType")
	if len(dataType) <= 0 {
		c.JSON(500,gin.H{
			"code" : 500,
			"msg" : "dataType为必填参数",
			"data":"",
		})
	}
	//逻辑处理 1.根据ID找到前缀
	gid,_ := strconv.Atoi(dataType)
	asset := mysql.GetAssetName(gid)
	data, err := logic.GetKeysByPfx(asset.TypeName)
	if err != nil {
		logger.Info(err)
		return 
	}

	//返回数据
	c.JSON(200,gin.H{
		"msg" : "ok",
		"code":200,
		"data" : data,
	})
}

//GetMarketPriceHandle 获取市场价格
func GetMarketPriceHandle(c *gin.Context)  {
	//获取参数
	//var p model.ParamTypeId
	//Berr := c.Bind(&p)
	//if Berr != nil {
	//	logger.Info(Berr)
	//	return
	//}
	//逻辑处理

	sliceInt := []int{17,15}
	marketPriceMap := make(map[string]interface{})
	for _,v := range sliceInt{
		d := mysql.GetAssetName(v)
		marketKey := fmt.Sprintf("%s.MarketPrice",d.TypeName)
		data, err := redis.GetData(marketKey)
		if err != nil {
			logger.Info(data)
			return
		}
		//返回数据
		marketPriceMap[strconv.Itoa(d.TypeId)] = data
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg" : "ok",
		"data":marketPriceMap,
	})

}

//SetStartParamHandler 设置启动参数
func SetStartParamHandler(c *gin.Context)  {
	//获取参数
	var p model.ParamStart
	err := c.Bind(&p)
	if err != nil {
		logger.Info(err)
		return 
	}
	if p.Buy==0||p.Sale==0||p.Safe ==0{
		c.JSON(500,gin.H{
			"code":500,
			"msg":"缺少相关参数",
			"data":"",
		})
		return
	}
	//逻辑处理
	err1 := redis.CreateDurableKey("buy", p.Buy)
	if err != nil {
		logger.Info(err1)
		return 
	}
	err2 := redis.CreateDurableKey("sale", p.Sale)
	if err != nil {
		logger.Info(err2)
		return
	}
	err3 := redis.CreateDurableKey("safe", p.Safe)
	if err != nil {
		logger.Info(err3)
		return 
	}
	//返回参数
	c.JSON(200,gin.H{
		"msg" : "ok",
		"code":200,
		"data":"",
	})
}

// GetBuyDataHandle 返回买入卖出的数据
func GetBuyDataHandle(c *gin.Context)  {

	//通过查询最新10条买入数据
		data1 := mysql.GetBuyData(1)
		result1 := make([]model.RespBuy,len(data1))
		for i,v := range data1{
			result1[i].Gid = v.Gid
			result1[i].Name = v.Name
			result1[i].Count = v.Count
			result1[i].TokenId = v.TokenId
			result1[i].MarketPrice = v.MarketPrice
			result1[i].SaleAddress = v.SaleAddress
			result1[i].FixedPrice = v.FixedPrice
			result1[i].TotalPrice = v.FixedPrice *float64(v.Count)
			result1[i].CreateTime = v.CreatedAt
			result1[i].Type = v.Type
			result1 = append(result1,result1[i])
		}
	//通过查询卖出的最新10条数据
		data2 := mysql.GetBuyData(2)
		result2 := make([]model.RespBuy,len(data2))
		for i,v := range data2{
			result2[i].Gid = v.Gid
			result2[i].Name = v.Name
			result2[i].Count = v.Count
			result2[i].MarketPrice = v.MarketPrice
			result2[i].SaleAddress = v.SaleAddress
			result2[i].Profit = v.Profit
			result2[i].TokenId = v.TokenId
			result2[i].FixedPrice = v.FixedPrice
			result2[i].TotalPrice = v.FixedPrice *float64(v.Count)
			result2[i].CreateTime = v.CreatedAt
			result2[i].Type = v.Type
			result2 = append(result2,result2[i])
		}


		c.JSON(200,gin.H{
			"buy_data": result1,
			"sale_data":result2,
			"msg" : "ok",
			"code":200,
		})
}

//前端系统监控使用

//SetMngRiskHandle 设置风控
func SetMngRiskHandle(c *gin.Context)  {
	//获取参数
	var p model.ParamRiskMng
	err := c.Bind(&p)
	if err != nil {
		logger.Info(err)
		return
	}
	if p.Situation==""||p.TimeLevel==0||p.Percentage ==0||p.OperationType==0||p.Status==0{
		c.JSON(500,gin.H{
			"msg":"缺少相关参数",
			"code":500,
			"data": "",
		})
		return
	}
	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["Situation"]= p.Situation
	m["TimeLevel"]=p.TimeLevel
	m["Percentage"]=p.Percentage
	m["OperationType"]=p.OperationType
	m["Status"]=p.Status
	logger.Info(m)
	redis.CreatHashKey(fmt.Sprintf("risk:%s",p.Situation),m)
	//返回参数
	c.JSON(200,gin.H{
		"msg" : "ok",
		"code":200,
		"data":"",
	})

}

//SetBuyAndSaleHandle 设置买入卖出百分比
func SetBuyAndSaleHandle(c *gin.Context){
	//获取参数
	var p model.ParamBuyAndSale
	err := c.Bind(&p)
	if err != nil {
		logger.Info(err)
		return
	}
	if p.ProductID==0||p.Status==0{
		c.JSON(500,gin.H{
			"code" : 500,
			"msg":"缺少相关参数",
			"data":"",
		})
		return
	}

	//逻辑处理 1.根据ID找到前缀
	asset := mysql.GetAssetName(p.ProductID)

	productName :=asset.TypeName
	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["ProductName"]= productName
	m["RisePercentage"]=p.RisePercentage
	m["FallPercentage"]=p.FallPercentage
	m["Status"]=p.Status
	logger.Info(m)
	redis.CreatHashKey(fmt.Sprintf("buyAndSale:%s",productName),m)
	//返回参数
	c.JSON(200,gin.H{
		"msg" : "ok",
		"code":200,
		"data":"",
	})
}

//SetParamOnOffHandle 设置买入卖出总开关
func SetParamOnOffHandle(c *gin.Context)  {
	//获取参数
	var p model.ParamOnOff
	err := c.Bind(&p)
	if err != nil {
		logger.Info(err)
		return
	}
	//if p.CrlName==""||p.Super==0{
	//	c.JSON(500,gin.H{
	//		"code":500,
	//		"msg":"缺少相关参数",
	//		"data":"",
	//	})
	//	return
	//}
	//把数据存进redis 中的哈希表
	m := make(map[string]interface{})
	m["CrlName"]= p.CrlName
	m["Super"]=p.Super
	logger.Info(m)
	redis.CreatHashKey(fmt.Sprintf("buyAndSale:%s", p.CrlName),m)
	//返回参数
	c.JSON(200,gin.H{
		"msg" : "ok",
		"code":200,
		"data":"",
	})
}

//GetScriptStatusHandle 获取脚本运行的状态 没写
func GetScriptStatusHandle(c *gin.Context)  {
	//获取买入卖出总开关
	var buyStruct model.RespAllOnOff
	buy:= redis.GetHashDataAll("buyAndSale:buy")
	mapstructure.Decode(buy,&buyStruct)

	var saleStruct model.RespAllOnOff
	sale:= redis.GetHashDataAll("buyAndSale:sale")
	mapstructure.Decode(sale,&saleStruct)

	//获取买入卖出元兽蛋开关
	var Egg model.RespBuyAndSale
	egg := redis.GetHashDataAll("buyAndSale:Metamon Egg")
	mapstructure.Decode(egg,&Egg)

	//获取买入卖出药水开关
	var Potion model.RespBuyAndSale
	potion := redis.GetHashDataAll("buyAndSale:Potion")
	mapstructure.Decode(potion,&Potion)

	var all model.RespAllSwitch
	allOnOffSlice := make([]model.RespAllOnOff,2)
	buyAndSaleSlice := make([]model.RespBuyAndSale,2)

	buyAndSaleSlice[0] = Egg
	buyAndSaleSlice[1] = Potion
	allOnOffSlice[0] =  buyStruct
	allOnOffSlice[1] =  saleStruct

	all.AllOnOff = allOnOffSlice
	all.BuyAndSale = buyAndSaleSlice

	c.JSON(200,gin.H{
		"msg" :"ok",
		"code":200,
		"data":all,
	})

//	//获取下跌风险开关
//	risk_fall, RFerr := redis.GetData("risk:fall")
//	if RFerr != nil {
//		logger.Error(RFerr)
//		return
//	}
//	//获取上涨风险开关
//	risk_rise,RRerr := redis.GetData("risk:rise")
//	if RRerr != nil{
//		logger.Error(RRerr)
//		return
//	}
//
//
}

//GetRiskMonitorHandle  返回监控信息状态
func GetRiskMonitorHandle(c *gin.Context)  {
	var fall  model.RespRiskMonitor
	var rise  model.RespRiskMonitor
	fallMap:= redis.GetHashDataAll("risk:fall")
	mapstructure.Decode(fallMap,&fall)
	riseMap := redis.GetHashDataAll("risk:rise")
	mapstructure.Decode(riseMap,&rise)

	all := make([]model.RespRiskMonitor,2)
	all[0] = fall
	all[1] = rise

	c.JSON(200,gin.H{
		"msg":"ok",
		"code":"200",
		"data":all,
	})
}

//GetMarketPriceLineHandle 获取对应的市场数据
func GetMarketPriceLineHandle(c *gin.Context)  {
	//参数
	var p model.ParamTypeId
	Berr := c.Bind(&p)
	if Berr != nil {
		logger.Error(Berr)
		return
	}

	//获取通过id获取类型的名称
	d := mysql.GetAssetName(p.TypeId)
	marketPriceKey := fmt.Sprintf("%s.MarketPrice",d.TypeName)

	//一小时前数据
	time1 := mysql.GetHistoryMarketData(3600,marketPriceKey)
	//两小时前数据
	time2 := mysql.GetHistoryMarketData(7200,marketPriceKey)
	//三小时前数据
	time3 := mysql.GetHistoryMarketData(10800,marketPriceKey)
	//四小时前数据
	time4 := mysql.GetHistoryMarketData(14400,marketPriceKey)
	//五小时前数据
	time5 := mysql.GetHistoryMarketData(18000,marketPriceKey)
	//六小时前数据
	time6 := mysql.GetHistoryMarketData(21600,marketPriceKey)


	//strtime1 := strconv.FormatFloat(time1.MarketData, 'E', -1, 64)
	//strtime2 := strconv.FormatFloat(time2.MarketData, 'E', -1, 64)
	//strtime3 := strconv.FormatFloat(time3.MarketData, 'E', -1, 64)
	//strtime4 := strconv.FormatFloat(time4.MarketData, 'E', -1, 64)
	//strtime5 := strconv.FormatFloat(time5.MarketData, 'E', -1, 64)
	//strtime6 := strconv.FormatFloat(time6.MarketData, 'E', -1, 64)

	var timeSlice  []float64
	timeSlice = append(timeSlice,time1.MarketData,time2.MarketData,time3.MarketData,time4.MarketData,time5.MarketData,time6.MarketData)

	//返回数据
	c.JSON(200,gin.H{
		"code":200,
		"msg":"ok",
		"data": timeSlice,
	})

}

// GetIncomeHandle 查询当前利润
func GetIncomeHandle(c *gin.Context)  {
	//获取利润
	data, Gerr := redis.GetData("income")
	if Gerr != nil {
		logger.Error(Gerr)
		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"msg":"ok",
		"data":data,
	})
}