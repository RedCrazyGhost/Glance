/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package main

import (
	"Glance/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	core.InitLog()
	core.InitChannelTrie()
	core.InitConfig()

}

func main() {
	core.GloabalLog.Println(core.NewAppMessage("Glance", "Version->"+VERSION))
	core.GloabalLog.Println(core.NewAppMessage("main", "项目启动！"))
	switch core.SystemConfig.Run {
	case "Shell":
		RunShell()
		break
	case "Server":
		RunServer()
		break
	default:
		core.SystemLogPool.Println("请配置正确的启动模式！Shell or Server")
		break
	}

	core.GloabalLog.Println(core.NewAppMessage("main", "项目关闭！"))
	closeLog()
}

// RunShell 以命令行的形式启动
func RunShell() {
	flow := core.NewDefaultTradeFlow()
	err := core.ReadCSV(core.SystemConfig.Shell.FilePath, core.SystemConfig.Shell.ParserName, flow)
	if err != nil {
		core.FailLog.Println(err.Error())
	}
	flow.Show()
	flow.CreateFile()
}

// RunServer 以服务器化的形式启动
func RunServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":" + core.SystemConfig.Server.Port)
	if err != nil {
		return
	}
}

func closeLog() {
	err := core.SystemLogPool.CloseAll()
	if err != nil {
		return
	}
	return
}
