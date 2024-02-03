package handler

import "github.com/gin-gonic/gin"

type Handler struct {}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUsersById)
			users.DELETE("/:id", h.deleteUser)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", h.getAllCategories)
			categories.GET("/:id", h.getCategoriesById)
			categories.DELETE("/:id", h.deleteCategory)
		}

		products := api.Group("/products")
		{
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProductsById)
			products.DELETE("/:id", h.deleteProduct)
		}
	}

	return router
}