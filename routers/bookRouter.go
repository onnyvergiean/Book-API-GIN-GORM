package routers

import (
	"tugas8/controllers"

	_ "tugas8/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API
// @version 1.0
// @description This is a sample server book server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func StartServer() *gin.Engine{
	router := gin.Default()

	// Read All
	router.GET("/books", controllers.GetBooks)
	// Read
	router.GET("/books/:id", controllers.GetBook)
	// Create
	router.POST("/books", controllers.CreateBook)
	// Update
	router.PUT("/books/:id", controllers.UpdateBook)
	// Delete	
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}