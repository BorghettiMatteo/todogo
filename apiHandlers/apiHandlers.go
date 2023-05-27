package apiHandlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// questo è un handler perchè ha la signature del tipo: func (c * gin.Context){}
func GetTasks(c *gin.Context) {
	//devo prendere dal db tutte le task e poi schiaffarle dentro un JSON e pusharlo verso il client
	fmt.Printf("c.Accepted: %v\n", c.Accepted)
}

func PostTasks(c *gin.Context) {

}

func UpdateTask(c *gin.Context) {

}

func DeleteTask(c *gin.Context)
