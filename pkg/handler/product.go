package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "getAllProducts",
	})
}

func (h *Handler) getProductsById(c *gin.Context) {
	productId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": productId,
	})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	productId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": productId,
	})
}