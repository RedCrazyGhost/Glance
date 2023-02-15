/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package main

import (
	"Glance/core"
	"fmt"
)

func init() {
	core.InitLog()
	//配置初始化

}

func main() {
	core.GloabalLog.Println(core.NewAppMessage("Glance", "Version->"+VERSION))
	core.GloabalLog.Println(core.NewAppMessage("main", "项目启动！"))
	fmt.Println("Hello,world！I am Glance!")
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
