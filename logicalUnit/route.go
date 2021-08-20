package main

import (
	"./controllers"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var r = gin.Default()

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func initRouter() {
	r.Use(CORSMiddleware())
	r.GET("/toTheEnd", controllers.ToTheEnd)
	r.POST("/setGlobals", controllers.SetGlobals)
	r.POST("/initData", controllers.SetInitData)
	r.GET("/nextEpoch", controllers.NextEpoch)
	r.GET("/currentData", controllers.CurrentData)
	r.GET("/mainData", controllers.GetMainData)
	r.GET("/fullData", controllers.GetFullData)
	r.GET("/firstData", controllers.GetFirstData)
	r.GET("/fullData/:id", controllers.GetFullDataById)
	r.Run()
}
