/**
 @author:way
 @date:2021/11/27
 @note
**/

package main

import (
	"fmt"
	"redisData/dao/mysql"
	"redisData/dao/redis"
	"redisData/logic"
	"redisData/pkg/logger"
	"redisData/setting"
	"strconv"
	"time"
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

func startBuy(listKey string,marketPrice string)  {
	var buy, _ = redis.GetData("buy")
	f, _ := strconv.ParseFloat(buy, 64)
	fmt.Printf("买入设置百分比为%f\n", f)
	//市场价直接从redis中取
	float64MarketPrice, _ := strconv.ParseFloat(marketPrice, 64)
	//执行买入脚本
	logic.SetBuyALG(listKey,float64MarketPrice, f)
	fmt.Println("本轮购买完毕")
	time.Sleep(time.Second * 1)
}
func main() {
	defer redis.Close()
	//开始缓存
	for {
		//key := "Metamon Egg.List"
		//marketPrice, _ := redis.GetData("Metamon Egg.MarketPrice")
		key := "Potion.List"
		marketPrice, _ := redis.GetData("Potion.MarketPrice")
		startBuy(key,marketPrice)
	}
}
