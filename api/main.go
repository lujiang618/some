package main

import (
	"some/api/job"
	"some/api/pkg/db"
	"some/api/pkg/web"
)

func init() {

}

// 入口文件，main函数
func main() {
	// gin.SetMode(setting.ServerSetting.RunMode)

	job.NewCommand().Process()

	web.StartWebServer()
	destroy()

}

func destroy() {
	db.Destroy() // 关闭数据库连接
}
