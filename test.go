/**
 @author:way
 @date:2021/12/2
 @note
**/

package main

import (
	"fmt"
	"redisData/dao/mysql"
	"redisData/dao/redis"
	"redisData/pkg/logger"
	"redisData/setting"
	"redisData/utils"
)

func init() {
	// 定义日志目录
	logger.Init("redisData")
	// 初始化 viper 配置
	if err := setting.Init(""); err != nil {
		logger.Info("viper init fail")
		logger.Error(err)
		return
	}
	mysql.InitMysql()
	//初始化redis
	if err := redis.InitClient(); err != nil {
		logger.Info("init redis fail err")
		logger.Error(err)
		return
	}

}
//func Hex2Dec(val string) *big.Int {
//	n, err := strconv.ParseUint(val, 16, 64)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return
//}
//
//func main() {
//	hex := "00000000000000000000000000000000000000000000a2bc77ee287ecf500000"
//	dec := Hex2Dec(hex)
//	fmt.Println(dec)
//}

//00000000000000000000000000000000000000000000131454ae75bda7500000

//func main() {
//	str := "0x467f963d000000000000000000000000d40c03b8680d4b6a4d78fc3c6f6a28c854e94a790000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000012bb890508c125661e03b09ec06e404bc928904000000000000000000000000000000000000000000000131454ae75bda750000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
//	str2 := []byte(str)
//	fmt.Println(string(str2[74+64+64+64:74+64+64+64+64]))
//	str3 := "00000000000000000000000000000000000000000000131454ae75bda7500000"
//	if str3 == string(str2[74+64+64+64:74+64+64+64+64]){
//		println("1111")
//	}
//}
//
func main() {
	str := "0xc37dfc5b38c77acb6b34adda5253c1993e88714d556513b2f4b61a1370cd64a2d6cbae5d000000000000000000000000000000000000000000000015af1d78b58c400000"
	str2 := []byte(str)
	fmt.Println(len(str2))
	//str3 := "0000000000000000000000000000000000000000000013451eb0c55622e00000"
	method := string(str2[:10])
	fmt.Println(method)
	xxx := string(str2[10:74])
	fmt.Println(utils.StringToBigInt(xxx))
	xxx2 := string(str2[74:138])
	fmt.Println(utils.StringToBigInt(xxx2))
}
		//xxx3 := string(str2[138:202])
		//println(xxx3)
		//xxx4 := string(str2[202:266])
		//println(xxx4)
	//	//account :=  string(str2[266:330])
	//	//fmt.Println(account)
	//	//xxx5 :=  string(str2[330:394])
	//	//fmt.Println(xxx5)
	//	//xxx6 :=  string(str2[394:458])
	//	//fmt.Println(xxx6)
	//	str1 := "1.0789442011729374e+48"
	//	float,_ := strconv.ParseFloat(str1,64)
	//	fmt.Println(float)
	//}

	//func main() {
	//
	//
	//		// 倒序：
	//		var kArray = []string{"1000", "1001", "1002", "1003", "1004", "1005"}
	//		sort.Slice(kArray, func(i, j int) bool {
	//			return kArray[i] > kArray[j]
	//		})
	//		fmt.Println("逆序：", kArray)
	//		// 正序：
	//		sort.Strings(kArray)
	//		fmt.Println("正序：", kArray)
	//
	//
	//}

	//func main() {
	//	m := redis.GetHashDataAll("buyAndSale:Metamon Egg")
	//	fmt.Println(m)
	//	data,_ := json.Marshal(m)
	//	fmt.Println(string(data))
	//
	//}

	//func main() {
	//	data1 := "1.05E+05"
	//	v1, err := strconv.ParseFloat(data1, 64)
	//	if err != nil {
	//		logger.Error(err)
	//	}
	//	fmt.Println(v1)
	//	//fmt.Println(fmt.Sprintf("%T", v1))
	//	//fmt.Println(fmt.Sprintf("%T", data))
	//}
	//func main() {
	//	timeStr:=time.Now().Format("2006-01-02 15:04:05")
	//	fmt.Println(timeStr)
	//}

	//func hexToBigInt(hex string) *big.Int {
	//	n := new(big.Int)
	//	n, _ = n.SetString(hex[2:], 16)
	//	return n
	//}
	//func main() {
	//	str := "8db3cfb7ca43b5c7b1090c98eee465b5e8bd0a653c6f98845c6070af2bcdf0cf"
	//	fmt.Println(hexToBigInt(str))
	//}
	//func FloatToString(input_num float64) string {
	//
	//	// to convert a float number to a string
	//	return strconv.FormatFloat(input_num, 'f', 6, 64)
	//}
	//
	//func main() {
	//	fmt.Println(FloatToString(21312421.213123))
	//}
