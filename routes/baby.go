package routes

import (
	"github.com/gin-gonic/gin"
	"redisData/controller/baby"
)

// RegisterWebRoutes 注册路由 baby 游戏路由

var babyController = new(baby.Controller)

func RegisterWebRoutes(router *gin.RouterGroup) {
	//买卖数据筛选
	router.GET("/buy_and_sell", babyController.BuyAndSellHandler)
	//设置baby风控
	router.GET("/setBabyRisk", babyController.SetBabyRiskHandle)
	//设置baby全自动和半自动
	router.GET("/setBabyOnOff", babyController.SetBabyOnOffHandle)
	//设置买入出参数
	router.GET("/setBabyBuyConf", babyController.SetBabyBuyConfHandle)
	//设置卖出参数
	router.GET("/setBabySaleConf", babyController.SetBabySaleConfHandle)
	//获取脚本运行的状态
	router.GET("/getBabyScriptStatus", babyController.GetBabyScriptStatusHandle)
	//设置卖出率
	router.GET("/setBabySaleRate", babyController.SetBabySaleRateHandle)
	//获取监控信息
	router.GET("/getBabyRiskMonitor", babyController.GetBabyRiskMonitorHandle)
	//接收买入卖出参数
	router.Any("/updateBabyOrder", babyController.UpdateBabyOrderHandle)
	//获取市场数据
	router.GET("/getBabyMarketPrice", babyController.GetBabyMarketPriceHandle)
	//设置私钥
	router.Any("/setPrivateKey",babyController.SetPrivateKeyHandle)
	//获取私钥
	router.GET("/getPrivateKey",babyController.GetPrivateKeyHandle)
}
