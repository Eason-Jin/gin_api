package apis

import (
	"fmt"
	"gin_api/moldels"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func GetSysDict(c *gin.Context) {

	p := moldels.SysDict{}
	sysdicts, err := p.GetSysDict()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(sysdicts)
	c.JSON(http.StatusOK, gin.H{
		"msg": "200",
	})
}

