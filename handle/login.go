package handle

import (
	"github.com/gin-gonic/gin"
	"tgin/utils"
)

func UCLogin(c *gin.Context) {
	words, _ := c.GetQuery("words")
	content := utils.Filter.Replace(words, '*')
	c.JSON(200, gin.H{"data": content})
}
