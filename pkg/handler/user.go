package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "getAllUsers",
	})
}

func (h *Handler) getUsersById(c *gin.Context) {
	usersId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": usersId,
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	usersId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": usersId,
	})
}

