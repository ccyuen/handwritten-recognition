// Package main initiates the connection to the db and the API routes
package main

import (
	// go router
	"github.com/gin-gonic/gin"
	//"net/http"

	// mongodb
	"github.com/mongodb/mongo-go-driver/mongo"

	//jwt
	//"github.com/dgrijalva/jwt-go"

	// std
	"context"
	"fmt"
	"log"
	"time"

	// custom
	"github.com/ccyuen/api/environments"
	"github.com/ccyuen/api/users"
)

const myDB string = "myapp"
const userCol string = "users"

var router *gin.Engine

var client *mongo.Client
var collection *mongo.Collection

func main() {
	env := environments.GetLocal()
	// start HTTP routes
	router = gin.Default() // or gin.new()
	router.Use(Database(env))
	initializeRoutes()
	router.Run() // can specify port as parameter

	// err = client.Disconnect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connection to MongoDB closed.")
}

func initializeRoutes() {
	// basic auth
	users.InitUserRoutes(router)
}

// Database provides the middle ware connection to Context in gin for access to DB object
func Database(env environments.Env) gin.HandlerFunc {
	var err error
	// connect to mongo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, env.Host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("okay connected to mongo")
	env.DB = client.Database(myDB)

	return func(c *gin.Context) {
		c.Set("DB", env.DB)
		c.Next()
	}
}
