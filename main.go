package main

import (
	"fmt"
	"main/todo"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

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
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	expiration, err := time.Parse("2006-01-02", "2023-05-27")

	samplequery := todo.ToDo{
		Activity:      "falcidiare bambini 2",
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

	db.Debug().Table(samplequery.TableName()).Create(&samplequery)
	//db.Last(&samplequery)

}

/*

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

//r := setupRouter()
	// creare il db

	connStr := "user=matteo dbname=ToDo password=password host=172.28.120.162 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		// il server si è schienato, ritorna
		fmt.Println("diomerda")
		return
	}
	// così mi ricordo di chiudere il db
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("merdacagata")
		panic(err)
	}

	statement := `INSERT INTO todo (activity, activityowner, expiration) VALUES ($1, $2,$3)`
	// usare 2006-01-02 come magical date
	expiration, err := time.Parse("2006-01-02", "2023-05-27")
	if err != nil {
		fmt.Println("strasuperdiomerda")
	}
	//creation, _ := time.Parse("YY-MM-DD", "2023-05-25")
	ret, err := db.Exec(statement, "matteo", "matteo", expiration)
	if err != nil {
		fmt.Println("diostramerda")
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("ret: %v\n", ret)
	/*


		var sampletodo todo.ToDo
		//creation, _ := time.Parse("YY-MM-DD", "2023-05-25")
		expiration, _ := time.Parse("YY-MM-DD", "2023-05-25")
		sampleRow := sampletodo.CreateToDo("impiccarsi", "matteo", expiration, false)
		fmt.Printf("sampleRow: %v\n", sampleRow)
		postdb.AutoMigrate(&todo.ToDo{})

		postdb.Select("activity", "activityowner", "expiration", "isdone").Create(&sampleRow)

		//fmt.Printf("result: %v\n", result.Error)

		//ora postgredb contiene il database effettivo con lo schema caricato
		//r.Run() // listen and serve on 0.0.0.0:8080
*/
