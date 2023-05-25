package apiHandlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// questo è un handler perchè ha la signature del tipo: func (c * gin.Context){}
func GetOwnersTasks(c *gin.Context) {
	a := 1
	fmt.Printf("a: %v\n", a)
}
