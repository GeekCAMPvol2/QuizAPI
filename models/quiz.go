package models

import "github.com/murasame29/rakutenapi"

type ReturnData struct {
	Quiz          string             `json:"quiz"`
	Answer        int                `json:"answer"`
	Images        []rakutenapi.Image `json:"images"`
	AffiliateLink string             `json:"affiliatelink"`
}
