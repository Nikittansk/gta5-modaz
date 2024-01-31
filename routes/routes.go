package routes

import (
	"github.com/Nikittansk/gta5-modaz/routes/api/ping"
	"github.com/Nikittansk/gta5-modaz/routes/api/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	
	setUpRouter(router)

	router.Run()
}

func setUpRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		ping.RegisterRouter(api)
		user.RegisterRouter(api)
	}
}

