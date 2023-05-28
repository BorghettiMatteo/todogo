package apiHandlers

import (
	"fmt"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// non mi devo preoccupare di sanificare l'input che viene dall'utente perchè gorm già escapa i caratteri usati in injection: https://gorm.io/docs/security.html
// questo è un handler perchè ha la signature del tipo: func (c * gin.Context){}
func GetTasks(c *gin.Context) {
	//devo prendere dal db tutte le task e poi schiaffarle dentro un JSON e pusharlo verso il client
	userName := c.Param("owner")
	//creazione di un array di ToDo poichè nulla vieta che un owner abbia uno o più todo
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
		c.AbortWithError(http.StatusInternalServerError, res.Error)
		return
	}
	// loading dei dati dentro il body della risposta come JSON indentato.
	c.IndentedJSON(http.StatusOK, tmpUser)

}

func PostTasks(c *gin.Context) {
	//devo usare shouldBindJson
	var tmpTodo model.ToDo

	res := c.BindJSON(&tmpTodo)
	// cosa abbiamo imparato, che se da postman faccio una quert senza rispettare lo schema di JSOn, si schiena tutto
	if res != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	// ora faccio tutti i controlli del caso su i dati passato
	//la data di fine deve essere > della data
	if tmpTodo.Expiration.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "expiration date cannot be less than creation date (as of today's date)"})
		return
	}
	// check se l'utente ha pushato un id
	if tmpTodo.Id != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "please do not insert the ID, postgres will handle it"})
		return
	}
	// se tutti i controlli sono andati bene, allora pusho la struttura nel DB
	err := model.Database.Create(&tmpTodo)

	if err.Error != nil {
		// inizio a metterli come placeholder, poi sistemerò
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "not able to push data into DB"})
		return
	}
	if err.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
}

func UpdateTask(c *gin.Context) {
	var tmpTodo model.ToDo
	type Boolway struct {
		Id   string `uri:"id" binding:"required"`
		Flag string `uri:"flag" binding:"required"`
	}
	var boolway Boolway
	err := c.ShouldBindUri(&boolway)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "whatever you have done, you have done bad"})
		return
	}
	// type check
	id, converr := strconv.Atoi(boolway.Id)
	if converr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "not able to use passed ID as primary Key"})
		return
	}
	flag, converr := strconv.ParseBool(boolway.Flag)
	if converr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "please pass a boolean value"})
		return
	}
	// check se ID è presente nel db
	exist := model.Database.Where("id = ?", id).First(&tmpTodo)
	if exist.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "the specified ID does not exist"})
		return
	}
	// update db record
	ret := model.Database.Model(&tmpTodo).Where("id = ?", id).Update("is_done", flag)
	if ret.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no action performed"})
		return
	}
}

func DeleteTask(c *gin.Context) {
	fmt.Println("a")
}

func UdateWholeTask(c *gin.Context) {
	fmt.Println("a")
}
