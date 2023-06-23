package api

import (
	"time"

	"github.com/GeekCAMPvol2/QuizAPI/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	conn   *db.Client
	router *gin.Engine
}

// 新しいHTTPサーバを作り、ルーティングを設定する
func NewServer(conn *db.Client) *Server {
	server := &Server{conn: conn}
	router := gin.Default()
	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"https://price-quest.netlify.app/"},
			AllowMethods: []string{"GET", "Fetch"},
			AllowHeaders: []string{
				"Access-Control-Allow-Credentials",
				"Access-Control-Allow-Headers",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
			},
			AllowCredentials: false,
			MaxAge:           24 * time.Hour,
		}))
	server.router = router
	quizRotuer(server)
	return server
}

func quizRotuer(server *Server) {
	server.router.GET("/quiz", server.getQuiz)
}

// サーバを開始する
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// エラーレスポンス
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
