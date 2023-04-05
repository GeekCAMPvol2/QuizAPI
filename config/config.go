package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/murasame29/rakutenapi"
)

// 情報を初期化する
func InitRequest() *rakutenapi.Service {
	var srv rakutenapi.Service
	err := godotenv.Load()
	if err != nil {
		log.Fatal("エラー！.envファイルが読み込めませんでした")
	}
	if len(os.Getenv("APPLICATION_ID")) == 0 {
		log.Panic("APPLICATION_IDが設定されていません")
		return nil
	}

	if len(os.Getenv("AFFILIATE_ID")) == 0 {
		log.Println("AFFILIATE_IDは設定されていません。")
	} else {
		srv.AffiliateId = os.Getenv("AFFILIATE_ID")
	}

	srv.ApplicationId = os.Getenv("APPLICATION_ID")
	srv.BasePath = "https://app.rakuten.co.jp"
	srv.RakutenSearch = rakutenapi.NewRakutenSearch(&srv)

	return &srv
}
