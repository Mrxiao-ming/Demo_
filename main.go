package main

import (
	"fmt"
	"time"
)

type A struct {
	a int
	b string
}

func main() {
	a1 := A{
		a: 1,
		b:"11",
	}
	a2 := A{
		a: 1,
		b:"11",
	}
	fmt.Println(a1 == a2)
}

func CheckTimeString(dd string) (time.Time, bool) {
	t, err := time.Parse("15:04:05", dd)
	if err != nil {
		return time.Time{}, false
	}
	//当前时间 - 当天0点
	return t, true
}

func CalcPossibilities1(str string) uint64 {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return CalcPossibilities2(str)
			}
		}
	}
	var (
		l   = len(str)
		res = uint64(1)
	)
	for i := 1; i <= l; i++ {
		res = res * uint64(i)
		fmt.Println(res)
	}
	return res
}

func CalcPossibilities2(str string) uint64 {
	var (
		distinctMap = make(map[string]struct{})
		l           = len(str)
	)
	distinctMap[str] = struct{}{}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			for s := range distinctMap {
				t := []byte(s)
				t[i], t[j] = t[j], t[i]
				distinctMap[string(t)] = struct{}{}
			}
		}
	}
	return uint64(len(distinctMap))
}

//func main() {
//	//var (
//	//	port = ":8080"
//	//	//env = flag.String("env", "dev", "请输入运行环境:\n  fat:测试环境\n uat:预上线环境\n product:正式环境 \n")
//	//)
//	//
//	//// mysqlConnect := connections.NewMysqlConnection()
//	//
//	//flag.Parse()
//	//
//	//fmt.Println("start time:", time.Now().Format(string(enum.FormatTime)))
//	//gin.SetMode(gin.DebugMode)
//	//gin.DisableConsoleColor()
//	//engine := gin.New()
//	//engine.Use(route.Cors())
//	//route.InitRoute(engine)
//	//fmt.Println("end time:", time.Now().Format(string(enum.FormatTime)))
//	//_ = engine.Run(port)
//
//}
