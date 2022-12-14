package poll

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPollController interface {
	Vote(ctx *gin.Context)
}

type PollController struct {
}

func NewPollController() IPollController {
	return &PollController{}
}

func (c *PollController) Vote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "It's working"})
}
