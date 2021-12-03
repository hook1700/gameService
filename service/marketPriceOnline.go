/**
 @author:way
 @date:2021/12/1
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
	"time"
)

func init() {
	// 定义日志目录
	logger.Init("marketPriceOnline")
	// 初始化 viper 配置
	if err := setting.Init(""); err != nil {
		logger.Info("viper init fail")
		logger.Error(err)
		return
	}
	mysql.InitMysql()
	//初始化redis
	if err := redis.InitClient(); err != nil {
		logger.Error(err)
		return
	}

}

func main() {
	defer redis.Close()
	for {
		//1.请求数据
		pageSize := 10
		category := 17
		data := logic.GetAssertsData(pageSize, category)
		fmt.Println("拿到数据")
		//计算市场价然后返回
		logic.SetMarketPriceOnline(data)
		fmt.Println("算出市场价")
		time.Sleep(1 * time.Second)
	}
}
