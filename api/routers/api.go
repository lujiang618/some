package routers

import (
	"some/api/application/controllers"
	"some/api/application/middleware"
)

func CreateApiRouter() {

	// // 会员API路由配置
	// v1Member := Router.Group("/v1/user").Use(middleware.VerifyAuth())
	// {
	// 	v1Member.GET("/info/:id", controller2.NewMember().Info)           // 会员详情
	// 	v1Member.PUT("/password", controller2.NewMember().UpdatePassword) // 修改密码
	// 	v1Member.POST("/token", controller2.NewMember().CreateToken)      // 生成TOKEN
	// }

	// // 登录认证API路由配置
	// v1Auth := Router.Group("/v1/auth")
	// {
	// 	v1Auth.POST("/login", controller2.NewAuth().Login)       // 登录
	// 	v1Auth.POST("/logout", controller2.NewAuth().Logout)     // 登出
	// 	v1Auth.POST("/register", controller2.NewAuth().Register) // 注册
	// 	v1Auth.GET("/forget", controller2.NewAuth().Forget)      // 忘记密码
	// }

	// 登录认证API路由配置
	v1Cost := Router.Group("/v1/w/c")
	{
		v1Cost.GET("/cost", controllers.NewCostController().GetList)                      // 支出列表
		v1Cost.POST("/cost", controllers.NewCostController().Create)                      // 添加支持记录
		v1Cost.PUT("/cost", controllers.NewCostController().Update)                       // 修改支持记录
		v1Cost.GET("/cost/categories", controllers.NewCostController().GetCostCategories) // 支出分类

	}

	v1Earning := Router.Group("/v1/w/e")
	{
		v1Earning.GET("/earning", controllers.NewEarningController().GetList) // 支出列表
		v1Earning.POST("/earning", controllers.NewEarningController().Create) // 支出列表
		v1Earning.PUT("/earning", controllers.NewEarningController().Update)  // 支出列表
	}

	v1Analyse := Router.Group("/v1/w/a")
	{
		v1Analyse.POST("/analyse", controllers.NewAnalyseController().GetAll) // 支出列表
	}

	// 系统API路由配置
	system := Router.Group("/system").Use(middleware.VerifyAuth())
	{
		system.GET("/health", controllers.NewSystem().Health)    // 监控检查
		system.POST("/setting", controllers.NewSystem().Setting) // 监控检查
	}

}
