/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package main

import (
	"Glance/core"
)

func init() {
	core.InitLog()
	core.InitChannelTrie()
	core.InitConfig()

}

func main() {
	core.GloabalLog.Println(core.NewAppMessage("Glance", "Version->"+VERSION))
	core.GloabalLog.Println(core.NewAppMessage("main", "项目启动！"))
	core.GloabalLog.Println(core.NewAppMessage("Glance", "Hello,world！I am Glance!"))

	flow := core.NewDefaultTradeFlow()
	err := core.ReadCSV("test.csv", "AlipayParser", flow)
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
