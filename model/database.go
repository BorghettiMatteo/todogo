package model

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func CreateDatabase() error {
	connStr := "user=matteo dbname=ToDo password=password host=172.28.120.162 port=5432 sslmode=disable"
	var err error
	// qui non riuscivo a passare il db poichè facevo lo shadowing della varibile Database facendo l'assegnazione implicita di quest'ultima mi scassava tutto
	// perchè creava una variabile Database che aveva come scope la funzione CreaDatabase() e di conseguenza non potevo esportarla

	Database, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Europe/Rome")
			return time.Now().In(ti)
		},
	})
	fmt.Printf("Database: %v\n", Database)

	if err != nil {
		fmt.Println("error starting DB")
		return err
	}
	superdb, err := Database.DB()

	if err != nil {
		fmt.Println("no DB connection")
		return err
	}
	err = superdb.Ping()
	if err != nil {
		fmt.Printf("no ping to DB")
		return err
	}
	return nil
}
