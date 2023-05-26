package main

import (
	"fmt"
	"main/model"
)

// prima faccio partire GIn, poi nel caso

/*
	func setupRouter() *gin.Engine {
		router := gin.Default()
		// grouping per rendere tutto più efficace
		v1 := router.Group("/api/v1")
		{
			// GET per prendere tutti i task di un dato owner
			v1.GET("/task/:owner", apiHandlers.GetOwnersTasks)
		}
		return router
	}
*/

func main() {

	db, _ := model.CreateDatabase()

	//expiration, err := time.Parse("2006-01-02", "2023-05-27")
	samplequery := model.ToDo{
		Id: 4,
	}

	//data handling
	/*
		res := db.First(&samplequery)
		samplequery.ActivityOwner = "gesù terrorista"
		db.Save(&samplequery)
	*/
	res := db.Model(&samplequery).Update("activity", "tiratore pazzo")
	fmt.Printf("res: %v\n", res.RowsAffected)

}
