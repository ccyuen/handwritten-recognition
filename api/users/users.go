// Package users holds the functions related to users endpoints to which the HTTP router references
package users

import (
	// "fmt"
	// "log"
	// "context"

	"github.com/gin-gonic/gin"
)

// A User represents login information for a single account
type User struct {
	Email    string
	Password string
}

var userCol = "users"

// InitUserRoutes initializes the routes specific to user endpoints
func InitUserRoutes(router *gin.Engine) {
	//router.GET("/users", getUsers)
	router.POST("/users/:uid", createUser)
	//router.DELETE("/users/:uid", deleteUser)
	router.GET("/users/:email", getUser)
}

// func getUsers(c *gin.Context) []User {

// }

// CreateUser creates a new account and registers it to the DB
func createUser(c *gin.Context) {
	// email := "test@gmail.com"
	// password := "test123"
	// newUser := User{Email: email, Password: password}

	// insertResult, err := collection.InsertOne(context.TODO(), newUser)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Inserted a single document: ", insertResult.InsertedID)
}

// DeleteUser deletes a user from the DB
func deleteUser(c *gin.Context) {

}

// GetUser returns the User object with some uid
func getUser(c *gin.Context) {
	// collection := c.MustGet("DB").(*mongo.Database).Collection(userCol)
	// var result User
	// filter := bson.D{{"email", c.Param("email")}}
	// err := collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Got this result: ", result)
}
