package main

import (
	"main/apiHandlers"
	"main/auth"
	"main/model"

	"github.com/gin-gonic/gin"
)

// prima faccio partire GIn, poi nel caso

func setupRouter() *gin.Engine {
	router := gin.Default()
	// grouping per rendere tutto più efficace
	headerGroup := router.Group("/api/v1")
	protected := headerGroup.Group("/protected")
	{
		protected.Use(auth.AuthMiddleware())
		// GET per prendere tutti i task di un dato owner
		//non serve passare *gin.Context perchè GetOwnerTask implementa implicitamente l'interfaccia func handler(*gin.Context)
		//la get sarà del tipo /api/v1/task?user=
		protected.GET("/task/", apiHandlers.GetTasks)
		protected.POST("/task", apiHandlers.PostTasks)
		protected.PUT("/task/:id", apiHandlers.UpdateTask)
		protected.PUT("/task", apiHandlers.UpdateWholeTask)
		// /api/v1/task?id=
		protected.DELETE("/task/", apiHandlers.DeleteTask)
		protected.GET("/health", apiHandlers.ReturnHealthAPI, apiHandlers.AnotherHealthFunc)
		protected.GET("/sampleAuth", apiHandlers.SampleAuth)

	}
	public := headerGroup.Group("/public")
	{
		public.POST("/login", apiHandlers.Login)
		public.POST("/register", apiHandlers.RegisterUser)
	}
	return router
}

func main() {

	// inizializzo il db, devo farmi passare anche l'errore eventualmente da stampare su return

	//creazione router
	currentRouter := setupRouter()
	err := model.CreateDatabase()
	model.Database.AutoMigrate(&model.ToDo{})
	model.Database.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	currentRouter.Run()
}
