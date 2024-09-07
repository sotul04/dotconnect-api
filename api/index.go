package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API ready to use!")
	})
}

func init() {
	app = gin.New()
	r := app.Group("/")
	myRoute(r)
}


func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}