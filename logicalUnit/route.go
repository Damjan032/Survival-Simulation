package main

import (
	"./controllers"
	"./models"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
)

var router *mux.Router = nil
var r = gin.Default()

func GetHelloWorld(c *gin.Context) {
	var population = models.NewPopulation(1, 5)
	fmt.Println(json.Marshal(population))
	btResult, _ := json.Marshal(population)
	c.Data(http.StatusOK, "application/json", btResult)
	//c.JSONP(200, string(btResult) )

}
func PostElements(c *gin.Context) {
	//var p = c.PostForm("search-field")
	var numberOfGoods = c.PostForm("number-of-goods")
	fmt.Println(numberOfGoods)
	c.JSONP(200, gin.H{
		"message": "sve ok",
	})

}
func initRouter() {
	r.GET("/", GetHelloWorld)
	r.GET("/toTheEnd", controllers.ToTheEnd)
	r.POST("/", PostElements)
	r.POST("/setGlobals", controllers.SetGlobals)
	r.POST("/initData", controllers.SetInitData)
	r.GET("/nextEpoch", controllers.NextEpoch)
	r.GET("/currentData", controllers.CurrentData)
	r.Run()
	/*router = mux.NewRouter()
	router.HandleFunc("/api/start", simulation).Methods("GET")
	/*router.HandleFunc("/api/events", createEventEndpoint).Methods("POST")
	router.HandleFunc("/api/events/{id:[0-9]+}", updateEventEndpoint).Methods("PUT")
	router.HandleFunc("/api/events/{id:[0-9]+}", deleteEventEndpoint).Methods("DELETE")
	router.HandleFunc("/api/statistic/{start:[0-9]+}/{end:[0-9]+}", statisticEndpoint).Methods("GET")*/

	/*routerCors := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//AllowedHeaders:   []string{"Authorization", "Content-Type"},
		//AllowCredentials: true,
		//Debug:            true,
	})

	staticDir := "/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	server := http.Server{
		Addr:    ":8080",
		Handler: routerCors.Handler(router),
	}
	fmt.Println("START SERVER")
	log.Fatal(server.ListenAndServe())*/
}