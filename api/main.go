// Package main initiates the connection to the db and the API routes
package main

import (
	// go router
	"github.com/gin-gonic/gin"

	// aws go driver
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	// std
	// "context"
	"fmt"
	// "log"
	// "time"

	// custom
	// "github.com/ccyuen/handwritten-recognition/api/users"
)

const myDB string = "myapp"
const userCol string = "users"

var router *gin.Engine

func main() {
	print("hello")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	print("yo")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Error in creating session ", err)
		return
	}
	svc := s3.New(sess)

	bucket := "handwritten-math-images"

	// Create input to put something in the bucket
	svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
	})

	putInput := &s3.PutBucketTaggingInput{
        Bucket: aws.String(bucket),
        Tagging: &s3.Tagging{
            TagSet: []*s3.Tag{
                {
                    Key:   aws.String("test"),
                    Value: aws.String("tag1"),
                },
            },
        },
    }

    _, err = svc.PutBucketTagging(putInput)
    if err != nil {
        fmt.Println("Error while trying to put item", err)
        return
	}
	
	// Create input for GetBucket method
    getInput := &s3.GetBucketTaggingInput{
        Bucket: aws.String(bucket),
    }

    result, err := svc.GetBucketTagging(getInput)
    if err != nil {
        fmt.Println("Error while trying to get item", err)
        return
    }

    fmt.Println(result.TagSet)
	// env := environments.GetLocal()
	// start HTTP routes
	// router = gin.Default() // or gin.new()
	// router.Use(Database(env))
	// initializeRoutes()
	// router.Run() // can specify port as parameter

	// err = client.Disconnect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connection to MongoDB closed.")
}

func initializeRoutes() {
	// basic auth
	// users.InitUserRoutes(router)
}

// Database provides the middle ware connection to Context in gin for access to DB object
// func Database(env environments.Env) gin.HandlerFunc {
	// var err error
	// // connect to mongo
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err = mongo.Connect(ctx, env.Host)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("okay connected to mongo")
	// env.DB = client.Database(myDB)

	// return func(c *gin.Context) {
	// 	c.Set("DB", env.DB)
	// 	c.Next()
	// }
// }
