package config

import (
	"log"
	"os"

	"github.com/murasame29/rakutenapi"
)

// 情報を初期化する
func InitRequest(id string) *rakutenapi.Service {
	var srv rakutenapi.Service

	if len(os.Getenv("AFFILIATE_ID")) == 0 {
		log.Println("AFFILIATE_IDは設定されていません。")
	} else {
		srv.AffiliateId = os.Getenv("AFFILIATE_ID")
	}

	srv.ApplicationId = loadApplicationId(id)
	srv.BasePath = "https://app.rakuten.co.jp"
	srv.RakutenSearch = rakutenapi.NewRakutenSearch(&srv)

	return &srv
}

func loadApplicationId(id string) string {
	if len(os.Getenv("APPLICATION_ID_"+id)) == 0 {
		log.Fatalf("APPLICATION_ID_%sが設定されていません", id)
	}
	return os.Getenv("APPLICATION_ID_" + id)
}
