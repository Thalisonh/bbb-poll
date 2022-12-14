package poll

import "github.com/gin-gonic/gin"

func Router(noAuthReq *gin.RouterGroup) {
	c := NewPollController()

	noAuthReq.POST("/", c.Vote)
}
