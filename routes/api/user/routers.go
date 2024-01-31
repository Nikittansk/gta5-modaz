package user

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/users", getAll)
	r.GET("/users/:id", getUserById)
}