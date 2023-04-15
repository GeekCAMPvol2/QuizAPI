package main

import (
	"log"
	"serv/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("エラー！.envファイルが読み込めませんでした")
	}

	router := gin.New()
	router.Use(cors.Default())

	routes.QuizRouter(router)
	router.Run()
}
