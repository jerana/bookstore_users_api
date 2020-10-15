package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//"StartApplication: Entry point of bookstore application"
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
