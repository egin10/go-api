package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/routers"
)

func main() {
	port := ":9000"
	router := gin.Default()

	routers.BookRouters(router)

	router.Run(port)
	fmt.Println("Server running on port" + port)
}
