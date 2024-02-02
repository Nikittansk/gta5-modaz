package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "getAllCategories",
	})
}

func (h *Handler) getCategoriesById(c *gin.Context) {
	categoryId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": categoryId,
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": categoryId,
	})
}