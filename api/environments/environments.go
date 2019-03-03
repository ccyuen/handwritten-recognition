// Package environments holds the values for the db to point to
package environments

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

// An Env represents the db environment to use
type Env struct {
	Host string // address of the db
	Key  string // access key
	DB   *mongo.Database
}

var local = Env{Host: "mongodb://localhost:27017", Key: ""}
var prod = Env{Host: "mongodb://localhost:27017", Key: ""} // the same for now

// GetLocal returns an Env object for the local environment
func GetLocal() Env {
	return local
}

// GetProd returns an Env object for the production environment
func GetProd() Env {
	return prod
}
