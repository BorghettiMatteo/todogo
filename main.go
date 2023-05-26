package main

import (
	"fmt"
	"main/model"
	"time"
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

	expiration, err := time.Parse("2006-01-02", "2023-05-27")

	samplequery := model.ToDo{
		Activity:      "giosuè carducci brigatista",
		ActivityOwner: "ma che davero?",
		Expiration:    expiration,
		IsDone:        true,
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	//data handling

	if err != nil {
		fmt.Println("strasuperdiomerda")
	}

	db.Debug().Create(&samplequery)
	newtodo := model.ToDo{
		Id: 6,
	}

	db.Take(&newtodo)
	fmt.Printf("new: %v\n", newtodo)
	//db.Last(&samplequery)

}
