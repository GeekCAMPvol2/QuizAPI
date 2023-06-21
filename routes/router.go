package routes

import (
	"serv/controller"

	"github.com/gin-gonic/gin"
)

// クイズに関するルーティング
func QuizRouter(router *gin.Engine) {
	// GET /quiz　で問題を取得できる
	// =クエリパラメタ=
	// keyword : キーワードを設定して検索できる。UTF-8でエンコードした文字列
	// genreid : 楽天市場におけるジャンルを検索するためのID
	// page    : 楽天検索からページを設定できる
	// sort	   : 0~6で設定可能
	//		0 : 楽天標準ソート <- default
	//		1 : 発売日順(降順)
	//		2 : 発売日順(昇順)
	//		3 : 売上順(降順)
	//		4 : 売上順(昇順)
	//		5 : 満足順（降順）
	//		6 : 満足順（昇順）
	router.GET("/quiz", controller.GetQuiztemp)
	router.GET("/quizlake", controller.GetQuizLake)
	router.GET("/all", controller.GetAll)
}
