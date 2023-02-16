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

	node := core.NewTradeNode(core.AlipayParser{}, []string{"1**********1", "2**************9", "4**********5", "收钱码收款", "2022/11/1 21:05", "**峰(194***@qq.com)", "300", "0", "481.61", "支付宝", "在线支付", ""}...)
	fmt.Println(node)
	core.GloabalLog.Println(core.NewAppMessage("main", "项目关闭！"))
	close()
}

func close() {
	err := core.SystemLogPool.CloseAll()
	if err != nil {
		return
	}
	return
}
