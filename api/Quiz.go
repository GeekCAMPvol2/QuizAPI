package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getQuizRequest struct {
	Hits int64 `form:"hits" binding:"min=1,max=10"`
}

func (server *Server) getQuiz(ctx *gin.Context) {
	var req getQuizRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Println(req.Hits)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Hits == 0 {
		req.Hits++
	}

	data, err := server.conn.GetQuiz(req.hits)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
