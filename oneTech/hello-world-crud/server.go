package main

import (
	"github.com/gin-gonic/gin"
	"learning/hello-world-crud/controllers"
)


func handleRequest(r *gin.Engine) {
	r.GET("/user", controllers.GetUser)
	r.GET("/user/:id", controllers.GetUserByID)
	r.POST("/user", controllers.PostUser)
	r.PUT("/user", controllers.PutUser)
	r.DELETE("/user", controllers.DeleteUser)
}


func main() {
	r := gin.Default()
	handleRequest(r)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//type Product struct {
//	gorm.Model
//	Code string
//	Price uint
//}
//
//func main() {
//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	db.AutoMigrate(&Product{})
//	//
//	db.Create(&Product{Code: "D42", Price: 100})
//
//	var product Product
//	db.First(&product, 1)
//	db.First(&product, "code=?", "D42")
//	fmt.Printf("%v", product)
//}

