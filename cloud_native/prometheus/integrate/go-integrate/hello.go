package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
http://0.0.0.0:8080/metrics
*/
func main() {
	r := gin.Default()
	r.GET("/metrics", PromHandler(promhttp.Handler()))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
