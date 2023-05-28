package apiHandlers

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// questo è un handler perchè ha la signature del tipo: func (c * gin.Context){}
func GetTasks(c *gin.Context) {
	//devo prendere dal db tutte le task e poi schiaffarle dentro un JSON e pusharlo verso il client
	userName := c.Param("owner")

	tmpUser := []model.ToDo{}
	// res ha la risposta che è del tipo *DB, quindi per l'errore devo accedere al campo res.Error
	res := model.Database.Where("activity_owner", userName).Find(&tmpUser)

	// check se l'owner è nel db, se così non fosse, 404
	if res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "user not in db"})
		return
	}

	if res.Error != nil {
		fmt.Println("errore db")
	}
	// marshalling dei dati di ritorno dal db
	c.IndentedJSON(http.StatusOK, tmpUser)

}

func PostTasks(c *gin.Context) {
	fmt.Println("a")
}

func UpdateTask(c *gin.Context) {
	fmt.Println("a")

}

func DeleteTask(c *gin.Context) {
	fmt.Println("a")
}
