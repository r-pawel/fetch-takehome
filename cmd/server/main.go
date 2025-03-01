package main

import (
	"github.com/gin-gonic/gin"
	"github.com/r-pawel/fetch-takehome/internal/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // TODO: conditionally set this based on ENV var

	r := router.NewRouter()
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
