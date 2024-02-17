package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojotx/gin-web/pkg/dice"
	"github.com/mojotx/gin-web/pkg/logging"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/rs/zerolog/log"
)

var srv *http.Server

func main() {

	logging.InitZerolog()

	r := gin.Default()
	_ = r.SetTrustedProxies([]string{})
	r.Use(cors.AllowAll())
	r.GET("/dice/:diceCmd", HandleDice)
	r.StaticFile("/", "./web/index.html")
	r.GET("/shutdown", HandleShutdown)

	srv = &http.Server{
		Addr:    "192.168.1.21:8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("failure with running router")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	stopServer()

}

func HandleDice(c *gin.Context) {
	diceCmd := c.Param("diceCmd")

	result := dice.ParseDiceCmd(diceCmd)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func HandleShutdown(c *gin.Context) {
	stopServer()
}

func stopServer() {
	log.Debug().Msg("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
