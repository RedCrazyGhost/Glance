/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package main

import (
	"Glance/core"
	"fmt"
	"runtime"
)

func init() {
	core.InitLog()
	core.InitChannelTrie()
	//配置初始化

	// 初始化后GC 回收
	runtime.GC()
}

func main() {
	core.GloabalLog.Println(core.NewAppMessage("Glance", "Version->"+VERSION))
	core.GloabalLog.Println(core.NewAppMessage("main", "项目启动！"))
	fmt.Println("Hello,world！I am Glance!")

	flow := core.NewDefaultTradeFlow()
	err := core.ReadCSV("test.csv", flow)
	if err != nil {
		core.FailLog.Println(err.Error())
	}
	flow.Show()

	core.GloabalLog.Println(core.NewAppMessage("main", "项目关闭！"))
	closeLog()
}

func closeLog() {
	err := core.SystemLogPool.CloseAll()
	if err != nil {
		return
	}
	return
}
