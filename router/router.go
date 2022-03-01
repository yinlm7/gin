package router

import (
	"github.com/gin-gonic/gin"
	"goAdmin/app/controller"
	"goAdmin/app/middleware/logger"
)

func InitRouter(port string) {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(logger.LoggerToFile())
	//router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	//router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	v1 := router.Group("user")
	{
		v1.GET("/test", controller.User)

		v1.GET("/ttt", controller.Test)
		v1.GET("/mgo", controller.MgoTest)

		////Info级别的日志
		//logger.Logger().WithFields(logrus.Fields{
		//	"name": "hanyun",
		//}).Info("记录一下日志", "Info")
		////Error级别的日志
		//logger.Logger().WithFields(logrus.Fields{
		//	"name": "hanyun",
		//}).Error("记录一下日志", "Error")
		////Warn级别的日志
		//logger.Logger().WithFields(logrus.Fields{
		//	"name": "hanyun",
		//}).Warn("记录一下日志", "Warn")
		////Debug级别的日志
		//logger.Logger().WithFields(logrus.Fields{
		//	"name": "hanyun",
		//}).Debug("记录一下日志", "Debug")
	}

	router.Run(":" + port)
}
