package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojotx/gin-web/pkg/logging"
)

func main() {

	logging.InitZerolog()

	r := gin.Default()
	r.GET("/dice/:diceCmd", HandleDice)

	/*
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	*/
	r.Run() // listen and serve on 0.0.0.0:8080

}

func HandleDice(c *gin.Context) {
	diceCmd := c.Param("diceCmd")

	c.JSON(http.StatusOK, gin.H{
		"result": 42,
	})
}

func HandleWeb(c *gin.Context) {
	c.HTML(http.StatusOK)
}
