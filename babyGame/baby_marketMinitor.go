/**
 @author:way
 @date:2021/12/16
 @note
**/
package main

import (
	"redisData/dao/mysql"
	"redisData/dao/redis"
	"redisData/pkg/email"
	"redisData/pkg/logger"
	"redisData/setting"
	"redisData/utils"
)

func init() {
	// 定义日志目录
	logger.Init("buy")
	// 初始化 viper 配置
	if err := setting.Init(""); err != nil {
		logger.Info("viper init fail")
		logger.Error(err)
		return
	}
	// 初始化MySQL
	mysql.InitMysql()
	//初始化redis
	if err := redis.InitClient(); err != nil {
		logger.Info("init redis fail err")
		logger.Error(err)
		return
	}
}

func main() {
	//一直循环监控
	for {
		//获取配置文件上涨下跌的全部参数
		rise := redis.GetHashDataAll("baby:ConfigRisk:rise")
		rise_operationType := rise["OperationType"]
		rise_percentage := rise["Percentage"]
		rise_percentage_float := utils.StringToFloat64(rise_percentage)
		rise_status := rise["Status"]
		rise_timeLevel := rise["TimeLevel"]
		rise_timeLevel_int := utils.StringToInt64(rise_timeLevel)


		fall := redis.GetHashDataAll("baby:ConfigRisk:fall")
		fall_operationType :=  fall["OperationType"]
		fall_percentage := fall["Percentage"]
		fall_percentage_float := utils.StringToFloat64(fall_percentage)
		fall_status := fall["Status"]
		fall_timeLevel := fall["TimeLevel"]
		fall_timeLevel_int := utils.StringToInt64(fall_timeLevel)

		//获取当前市场价格
		marketPrice, GErr := redis.GetData("baby:marketPrice")
		if GErr != nil {
			logger.Error(GErr)
			return
		}
		marketPrice_float := utils.StringToFloat64(marketPrice)
		//根据输入参数拿对应时间段的市场价
		//获取当前时间
		now := utils.GetNowTimeS()

		//获取设置时间
		timeOfRise := now - rise_timeLevel_int
		timeOfFall := now - fall_timeLevel_int
		//转化成str
		strTimeRise := utils.TimestampToDatetime(timeOfRise)
		strTimeFall := utils.TimestampToDatetime(timeOfFall)
		//根据设置的时间查询mysql的数据
		riseMarketData := mysql.GetMarketPriceByTime(strTimeRise)
		fallMarketData := mysql.GetMarketPriceByTime(strTimeFall)

		//判断脚本状态是否打开
		if rise_status == "1"{
			//对比市场价计算风控值  新市场价的除以旧的
			riseRate := marketPrice_float/riseMarketData.MarketData
			if riseRate > 1{   //大于1上涨
				//计算涨幅 ex. 11/10 -1
				if riseRate -1 > rise_percentage_float{
					//上涨警报
					//选择对应的操作
					switch rise_operationType {
					case "1":
						logger.Info("停止脚本")
					case "2":
						logger.Info("发送钉钉")
						email.SendDingMsg("市场价在上涨","市场价上涨率过高")
					case "3":
						logger.Info("停止脚本，且发送钉钉")
						//获取买卖脚本的配置文件，修改其中的状态
						//发送钉钉
						email.SendDingMsg("市场价在上涨","市场价上涨率过高")
					}
				}
			}
		}
		if fall_status == "1"{
			fallRate := marketPrice_float/fallMarketData.MarketData
			if marketPrice_float/fallMarketData.MarketData<1{   //小于1下跌
				//计算跌幅 ex. 1 - 9/10
				if 1 - fallRate  > fall_percentage_float{
					//下跌警报
					switch fall_operationType {
					case "1":
						logger.Info("停止脚本")
					case "2":
						logger.Info("发送钉钉")
						email.SendDingMsg("市场价在上涨","市场价上涨率过高")
					case "3":
						logger.Info("停止脚本，且发送钉钉")
						//获取买卖脚本的配置文件，修改其中的状态
						//发送钉钉
						email.SendDingMsg("市场价在上涨","市场价上涨率过高")
					}
				}
			}
		}
		logger.Info("上涨风控和下跌风控未打开")
	}



}