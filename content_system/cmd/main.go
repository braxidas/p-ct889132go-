package cmd

import (
	"ContentSystem/internal/api"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api.CmsRouters(r)

	err := r.Run()
	if err != nil{
		fmt.Printf("run error = %v", err)
		return
	}
}