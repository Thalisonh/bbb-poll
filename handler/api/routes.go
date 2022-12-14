package api

import (
	"github.com/Thalisonh/bbb-poll/handler/api/poll"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	poll.Router(r.Group("/poll"))
}
