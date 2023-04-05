package main

import (
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(cors.Default())

	routes.QuizRouter(router)

	router.Run()
}
