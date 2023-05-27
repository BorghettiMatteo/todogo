package apiHandlers

import (
	"fmt"
	"main/model"

	"github.com/gin-gonic/gin"
)

// questo è un handler perchè ha la signature del tipo: func (c * gin.Context){}
func GetOwnersTasks(c *gin.Context) {
	//devo prendere dal db tutte le task e poi schiaffarle dentro un JSON e pusharlo verso il client
	db, err := model.CreateDatabase()
	if err != nil {
		fmt.Println("ho panicato ")
	}

}
