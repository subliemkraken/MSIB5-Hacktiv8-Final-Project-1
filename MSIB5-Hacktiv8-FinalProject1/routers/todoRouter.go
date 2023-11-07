package routers

import (
	"MSIB5-Hacktiv8-FinalProject1/controllers"

	"github.com/gin-gonic/gin"

	_ "MSIB5-Hacktiv8-FinalProject1/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Todo Application
// @version 1.0
// @description This is a todo application
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSES-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/todo", controllers.GetAllTodo)

	router.POST("/todo", controllers.CreateTodo)

	router.PUT("/todo/:todoID", controllers.UpdateTodo)

	router.GET("/todo/:todoID", controllers.GetTodoWithID)

	router.DELETE("/todo/:todoID", controllers.DeleteTodo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
