package routes

import (
	"github.com/gin-gonic/gin"
	apiV1 "jiangchen.tech/demo_app/api/v1/route"
	"jiangchen.tech/demo_app/middleware"
)

type Option func(engine *gin.Engine)

var Options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	Options = append(Options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	// 加载多路由
	Include(apiV1.Routers)
	// 初始化日志
	//gin.DisableConsoleColor()
	//// 创建日志文件
	//f, _ := os.Create("go_demo_app.log")
	//// 写入日志
	//gin.DefaultWriter = io.MultiWriter(f)
	// 创建一个默认路由
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	// 加载注册的app路由
	for _, opt := range Options {
		opt(r)
	}
	// 抛出指针
	return r
}
