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
		v1.GET("/task/:owner", apiHandlers.GetTasks)
		v1.POST("/task", apiHandlers.PostTasks)
		v1.PUT("/task/:id", apiHandlers.UpdateTask)
		v1.DELETE("/task/:id", apiHandlers.DeleteTask)

	}
	return router
}

func main() {

	// inizializzo il db, devo farmi passare anche l'errore eventualmente da stampare su return

	model.CreateDatabase()

	//creazione router
	currentRouter := setupRouter()
	currentRouter.Run()
	/*
		//expiration, err := time.Parse("2006-01-02", "2023-05-27")
		samplequery := model.ToDo{
			Id: 5,
		}
		model.CreateDatabase()

		//data handling

			res := db.First(&samplequery)
			samplequery.ActivityOwner = "gesù terrorista"
			db.Save(&samplequery)

		res := (model.Database).Model(&samplequery).Update("activity", "tiratore pazzo")
		fmt.Printf("res: %v\n", res.RowsAffected)
	*/
}
