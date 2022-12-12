package main

import (
	"github.com/bryanmccarthy/deep-backend/db"
	"github.com/bryanmccarthy/deep-backend/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	r.Use(cors.Default())

	// routes
	users.Routes(r, h)

	r.Run(port)
}
