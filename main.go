package main

import (
	"main/apiHandlers"
	"main/model"

	"github.com/gin-gonic/gin"
)

// prima faccio partire GIn, poi nel caso

func setupRouter() *gin.Engine {
	router := gin.Default()
	// grouping per rendere tutto più efficace
	v1 := router.Group("/api/v1")
	{
		// GET per prendere tutti i task di un dato owner
		//non serve passare *gin.Context perchè GetOwnerTask implementa implicitamente l'interfaccia func handler(*gin.Context)
		//la get sarà del tipo /api/v1/task?user=
		v1.GET("/task/", apiHandlers.GetTasks)
		v1.POST("/task", apiHandlers.PostTasks)
		v1.PUT("/task/:id", apiHandlers.UpdateTask)
		v1.PUT("/task", apiHandlers.UpdateWholeTask)
		// /api/v1/task?id=
		v1.DELETE("/task/", apiHandlers.DeleteTask)
		v1.GET("/health", apiHandlers.ReturnHealthAPI, apiHandlers.AnotherHealthFunc)
		v1.GET("/sampleAuth", apiHandlers.SampleAuth)

	}
	public := v1.Group("/public")
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
