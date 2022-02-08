package application

import "github.com/gin-gonic/gin"

var r = gin.Default()

func StartApplication() {

	mapUrls()
	r.Run(":9000")
}
