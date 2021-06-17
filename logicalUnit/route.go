package main

import (
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
	var population = NewPopulation(1, 5)
	fmt.Println(json.Marshal(population))
	btResult, _ := json.Marshal(population)
	c.Data(http.StatusOK, "application/json", btResult)
	//c.JSONP(200, string(btResult) )

}
func PutElements(c *gin.Context) {
	var population = NewPopulation(1, 5)

	var p = c.PostForm("search-field")
	fmt.Println(p)
	c.JSONP(200, gin.H{
		"message": population,
	})

}
func initRouter() {
	r.GET("/", GetHelloWorld)
	r.POST("/", PutElements)
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
