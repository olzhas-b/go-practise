package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olzhas-b/hello-world-crud/models"
	"net/http"
)
func GetUser(c *gin.Context)  {
	//dp := db.GetDataBase()
	c.JSON(200, gin.H{
		"message": "userData must be here",
	})
	return
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Hello World from user %s", id)
	return
}

func PostUser(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("id: %d; name: %s; login: %s; phone: %s", user.ID, user.Name, user.Email, user.Phone)
	c.String(200, fmt.Sprintf("id: %d; name: %s; login: %s; phone: %s", user.ID, user.Name, user.Email, user.Phone))
	return
}

func PutUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(200, fmt.Sprintf("id: %d; name: %s; login: %s; phone: %s", user.ID, user.Name, user.Email, user.Phone))
	return
}

func DeleteUser(c *gin.Context) {
	var ID models.User
	if err := c.ShouldBindJSON(&ID); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	}
	fmt.Println(ID)
	c.String(200, fmt.Sprintf("deleted item ID : %d", ID.ID))
	return
}
