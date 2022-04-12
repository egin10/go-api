package routers

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

func BookRouters(router *gin.Engine) {
	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/book/:id", controllers.ShowBook)
	router.DELETE("/book/delete/:id", controllers.DeleteBook)
	router.PATCH("/book/update/:id", controllers.UpdateBook)
}
