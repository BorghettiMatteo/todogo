package model

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDatabase() (*gorm.DB, error) {
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

	superdb, err := db.DB()

	if err != nil {
		fmt.Println("diomerda sta andado tutto a fuoco")
	}
	err = superdb.Ping()
	if err != nil {
		fmt.Printf("\"superdiomerda\": %v\n", "superdiomerda")
		return nil, err
	}

	return db, err
}
