package main

import (
	"./controllers"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var r = gin.Default()

func initRouter() {
	r.GET("/toTheEnd", controllers.ToTheEnd)
	r.POST("/setGlobals", controllers.SetGlobals)
	r.POST("/initData", controllers.SetInitData)
	r.GET("/nextEpoch", controllers.NextEpoch)
	r.GET("/currentData", controllers.CurrentData)
	r.GET("/mainData", controllers.GetMainData)
	r.GET("/fullData", controllers.GetFullData)
	r.GET("/fullData/:id", controllers.GetFullDataById)
	r.Run()
}
