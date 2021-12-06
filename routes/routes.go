package routes

import (
	"github.com/gin-gonic/gin"
	"redisData/controller"
	"redisData/middleware"
)

func SetUp() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) //跨域
	r.Use(middleware.TraceLogger())
	//r.Use(middleware.TlsHandler())  // 支持wss
	 // 日志上下文进行绑定追踪
	//查询，查询redis上的数据，返回给前端
	//websocket
	v1 := r.Group("/api/game")
	v1.GET("/getData",controller.GetDataHandle)
	v1.GET("/getMarketPrice",controller.GetMarketPriceHandle)
	v1.GET("/setStartParam",controller.SetStartParamHandler)
	v1.GET("/getBuyData",controller.GetBuyDataHandle)
	v1.GET("/setMngRisk",controller.SetMngRiskHandle)
	v1.GET("/setBuyAndSale",controller.SetBuyAndSaleHandle)
	v1.GET("/setParamOnOff",controller.SetParamOnOffHandle)
	v1.GET("/getScriptStatus",controller.GetScriptStatusHandle)
	return r
}
