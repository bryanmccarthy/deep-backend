package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	r := gin.Default()

	r.Run(port)
}
