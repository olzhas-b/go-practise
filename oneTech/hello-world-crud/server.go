package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olzhas-b/hello-world-crud/controllers"
	"github.com/olzhas-b/hello-world-crud/db"
	"gorm.io/gorm"
)

func handleRequest(r *gin.Engine) {
	r.GET("/user", controllers.GetUser)
	r.GET("/user/:id", controllers.GetUserByID)
	r.POST("/user", controllers.PostUser)
	r.PUT("/user", controllers.PutUser)
	r.DELETE("/user", controllers.DeleteUser)
}
var dataBase *gorm.DB
func main() {
	r := gin.Default()
	dataBase = db.InitDataBase()
	handleRequest(r)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
