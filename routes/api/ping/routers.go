package ping

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/ping", ping)
}