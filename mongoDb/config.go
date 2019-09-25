/*

package models

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient("mongodb://joshi:joshi123@ds121182.mlab.com:21182/gaurav")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	var Db = client.Database("gaurav")
}

*/