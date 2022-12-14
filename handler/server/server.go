package server

import (
	"github.com/Thalisonh/bbb-poll/handler/api"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	api.Router(r)
	r.Run()
}
