package main

import (
	"os"
	"securityMS/cmd/utils"
	"securityMS/pkg/routes"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {

	config := read()

	port := os.Getenv("PORT")
	if port == "" {
		config.Host.Port = port
	}

	//waits for a collection of goroutines to finish
	var waitgroup sync.WaitGroup

	//takes the number of goroutines
	waitgroup.Add(2)

	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger())

	go func() {
		routes.InitializeLogin(*router)
		waitgroup.Done()
	}()

	go func() {
		routes.InitializeAuth(*router)
		waitgroup.Done()
	}()

	//waits for all the calls to finish
	waitgroup.Wait()

	router.Run(":" + config.Host.Port + "8080")
}

func read() utils.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}

	return config
}
