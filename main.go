package main

import (
	"fmt"
	"main/todo"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	connStr := "user=matteo dbname=ToDo password=password host=172.28.120.162 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Europe/Rome")
			return time.Now().In(ti)
		},
	})

	if err != nil {
		fmt.Println("si è schienata la configurazione con il db")
	}

	expiration, err := time.Parse("2006-01-02", "2023-05-27")

	samplequery := todo.ToDo{
		Activity:      "falcidiare bambini 13",
		ActivityOwner: "matteo",
		Expiration:    expiration,
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	db.Debug().AutoMigrate(&todo.ToDo{})

	superdb, err := db.DB()

	if err != nil {
		fmt.Println("diomerda sta andado tutto a fuoco")
	}
	err = superdb.Ping()
	if err != nil {
		fmt.Printf("\"superdiomerda\": %v\n", "superdiomerda")
		return
	}

	//data handling

	if err != nil {
		fmt.Println("strasuperdiomerda")
	}

	db.Debug().Create(&samplequery)
	newtodo := todo.ToDo{
		Id: 6,
	}

	db.Take(&newtodo)
	fmt.Printf("new: %v\n", newtodo)
	//db.Last(&samplequery)

}
