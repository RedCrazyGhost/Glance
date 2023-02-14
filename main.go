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
	initLog()
	//配置初始化

}

func main() {
	GloabalLog.Println(core.NewAppMessage("main", "项目启动！"))
	fmt.Println("Hello,world！I am Glance!")
	GloabalLog.Println(core.NewAppMessage("main", "项目关闭！"))
	close()
}

func close() {
	err := logPool.closeAll()
	if err != nil {
		return
	}
	return
}
