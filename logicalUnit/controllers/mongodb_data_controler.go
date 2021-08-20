package controllers

import (
	"../services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMainData(c *gin.Context) {
	btResult, _ := json.Marshal(services.ReadAllMainParams())
	c.Data(http.StatusOK, "application/json", btResult)
}

func GetFullData(c *gin.Context) {
	btResult, _ := json.Marshal(services.ReadFullDataDocument())
	c.Data(http.StatusOK, "application/json", btResult)
}

func GetFirstData(c *gin.Context) {
	btResult, _ := json.Marshal(services.ReadFullFirstData())
	c.Data(http.StatusOK, "application/json", btResult)
}

func GetFullDataById(c *gin.Context) {
	btResult, _ := json.Marshal(services.FindAllDataById(c.Param("id")))
	c.Data(http.StatusOK, "application/json", btResult)
}
