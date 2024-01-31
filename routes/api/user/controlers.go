package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []user{
	{
		ID:       1,
		Username: "Nikittansk",
		Password: "123456789",
		Role:     "admin",
	},
	{
		ID:       2,
		Username: "Anton",
		Password: "1234567890",
		Role:     "user",
	},
}

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	for _, v := range users {
		if v.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": v,
			})
			return
		}
	}
}