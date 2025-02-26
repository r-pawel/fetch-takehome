package main

import (
	"github.com/R-Pawel/fetch-takehome/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // TODO: conditionally set this based on ENV var

	r := router.NewRouter()
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
