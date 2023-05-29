package apiHandlers

import (
	"main/model"
	"net/http"
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
	if res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "user not in db"})
		return
	}
	//se tutto ok
	// loading dei dati dentro il body della risposta come JSON indentato.
	c.IndentedJSON(http.StatusOK, tmpUser)

}

func PostTasks(c *gin.Context) {
	//devo usare shouldBindJson
	var inputToDo model.ToDo

	res := c.BindJSON(&inputToDo)
	// cosa abbiamo imparato, che se da postman faccio una quert senza rispettare lo schema di JSOn, si schiena tutto
	if res != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": res.Error()})
		return
	}
	// ora faccio tutti i controlli del caso su i dati passato
	//la data di fine deve essere > della data
	if inputToDo.Expiration.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "expiration date cannot be less than creation date (as of today's date)"})
		return
	}
	// check se l'utente ha pushato un id che giùesiste me ne frego e passo olte

	if inputToDo.Id != 0 {
		inputToDo.Id = 0
	}

	// se tutti i controlli sono andati bene, allora pusho la struttura nel DB
	err := model.Database.Create(&inputToDo)

	if err.Error != nil {
		// inizio a metterli come placeholder, poi sistemerò
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error.Error()})
		return
	}
	// davvero necessario?
	if err.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
}

func UpdateTask(c *gin.Context) {
	var inputToDo model.ToDo
	id := c.Param("id")

	// check se ID è presente nel db
	exist := model.Database.Where("id = ?", id).First(&inputToDo)
	if exist.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": exist.Error.Error()})
		return
	}
	// update db record
	if inputToDo.IsDone {
		ret := model.Database.Model(&inputToDo).Update("is_done", false)
		if ret.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": exist.Error.Error()})
			return
		}
	} else {
		ret := model.Database.Model(&inputToDo).Update("is_done", true)
		if ret.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": exist.Error.Error()})
			return
		}
	}
}

func DeleteTask(c *gin.Context) {
	var inputToDo model.ToDo
	id := c.Param("id")
	//parse id to int
	err := model.Database.Where("Id = ?", id).First(&inputToDo)
	// se errore != null vuol dire che non è riuscita a fare la delete oppure l'id non esiste
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error.Error()})
		return
	}
	ret := model.Database.Delete(&inputToDo)
	if ret.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": ret.Error.Error()})
		return
	}

}

func UdateWholeTask(c *gin.Context) {
	var sampleToDo model.ToDo
	var passedTodo model.ToDo
	res := c.BindJSON(&passedTodo)

	if res != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": res.Error()})
		return
	}
	//controllo che l' entità esista:
	err := model.Database.Where("id = ?", passedTodo.Id).First(&sampleToDo)
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error.Error()})
		return
	}
	// controllo che i dati passati siano consistenti
	//controllo data expiration > data attuale
	if passedTodo.Expiration.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "expiration date cannot be less than creation date (as of today's date)"})
		return
	}
	// che sia passedToDo o tempToDo è indifferente da passare a Model()
	ret := model.Database.Model(&sampleToDo).Updates(&passedTodo)
	if ret.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": ret.Error.Error()})
		return
	}
}

func ReturnHealthAPI(c *gin.Context) {
	superdb, err := model.Database.DB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = superdb.Ping()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "db is ok"})
}
