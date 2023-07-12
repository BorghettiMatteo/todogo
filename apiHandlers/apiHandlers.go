package apiHandlers

import (
	"crypto/sha512"
	"fmt"
	"main/auth"
	"main/model"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func checkIDInput(id string) bool {
	digitReg, _ := regexp.Compile(`^[\d]+$`)
	return digitReg.MatchString(id)
}

// non mi devo preoccupare di sanificare l'input che viene dall'utente perchè gorm già escapa i caratteri usati in injection: https://gorm.io/docs/security.html
// questo è un handler perchè ha la signature del tipo: func (c * gin.Context){}
func GetTasks(c *gin.Context) {
	//devo prendere dal db tutte le task e poi schiaffarle dentro un JSON e pusharlo verso il client
	userName := c.Query("user")
	if userName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "valore nullo"})
		return
	}
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
	if !checkIDInput(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot pass non digit ID"})
		return
	}
	// check se ID è presente nel db
	exist := model.Database.Where("id = ?", id).First(&inputToDo)
	if exist.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": exist.Error.Error()})
		return
	}
	// update db record
	ret := model.Database.Model(&inputToDo).Update("is_done", !inputToDo.IsDone)
	if ret.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": exist.Error.Error()})
		return
	}
}

func DeleteTask(c *gin.Context) {
	var inputToDo model.ToDo
	id := c.Query("id")

	if !checkIDInput(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot pass non digit ID"})
		return
	}
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

func UpdateWholeTask(c *gin.Context) {
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

func AnotherHealthFunc(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "eccocil"})
}

func Login(c *gin.Context) {
	// check su db se c'è lo user
	var inputUser model.User
	var dbUser model.User

	res := c.BindJSON(&inputUser)
	if res != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": res.Error()})
		return
	}
	//
	dbret := model.Database.Where("username", inputUser.Username).Find(&dbUser)
	if dbret.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": dbret.Error})
		return
	} else {
		hasher := sha512.New()
		hasher.Write([]byte(inputUser.HashedPassword))
		a := hasher.Sum(nil)
		b := string(a)
		inputUser.HashedPassword = fmt.Sprintf("%x", b)
		if dbUser.HashedPassword == inputUser.HashedPassword {
			signedString, e := auth.GenerateToken(dbUser)
			if e == nil {
				c.Request.Header.Add("Bearer", signedString)
				c.IndentedJSON(http.StatusOK, gin.H{"token": signedString})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "wrong password, retry!"})
			return
		}
	}
}

func RegisterUser(c *gin.Context) {
	var newUser model.User

	res := c.BindJSON(&newUser)
	hasher := sha512.New()
	hasher.Write([]byte(newUser.HashedPassword))
	a := hasher.Sum(nil)
	b := string(a)
	newUser.HashedPassword = fmt.Sprintf("%x", b)
	if res != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": res.Error()})
		return
	}
	err := model.Database.Create(&newUser)
	if err.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	// qui devo gestire la fase di creazione del token
	signedString, e := auth.GenerateToken(newUser)
	if e == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"token": signedString})
	}

}

func SampleAuth(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "valore nullo"})
		return
	}
	if auth.VerifyToken(token) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "authok"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "authko"})
	}
}

func LogOff(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "valore nullo"})
		return
	}
	auth.SecretKey = []byte("orcolalamadonna")
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ciaociao"})
}
