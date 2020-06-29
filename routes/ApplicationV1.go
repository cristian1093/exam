package routes

import (
	"github.com/controllers"
	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {

	v1 := router.Group("/v1")
	{

		v1.POST("/upload", controllers.UploadLayout)

		v1.POST("/upload_file", controllers.UploadFiles)

		v1.POST("/registration", controllers.Registration)
		v1.POST("/login", controllers.Login)
		v1.POST("/login2", controllers.Login2)
		v1.POST("/membresia", controllers.Membresia)
	}

}
