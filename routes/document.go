package routes

import (
	"github.com/7lawa9111/Project-Based-Mentorship-Go-App/issues/controller"
	"github.com/gin-gonic/gin"
)

func DocumentRoute(router *gin.Engine) {
	router.GET("/", controller.GetDocuments)
}
