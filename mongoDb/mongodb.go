package mongoDb

import (
	. "GoMysql/models"
	. "GoMysql/service"

	"context"
	"fmt"
	"log"
	"mongo-go-driver-master/bson"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

//client, err := mongo.NewClient("mongodb://<joshi>:<joshi123>@ds121182.mlab.com:21182/gaurav")

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var client, err = mongo.Connect(context.TODO(), "mongodb://joshi:joshi123@ds121182.mlab.com:21182/gaurav")

var db = client.Database("gaurav")

func Insert(person Person) {
	hashValue, _ := HashPassword(person.GetPassword())
	person.SetPassword(hashValue)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	//person.ID = bson.TypeObjectID()

	//doc := db.Collection(COLLNAME).FindOneAndUpdate

	res, _ := db.Collection("person1").InsertOne(ctx, &person)

	if err != nil {
		fmt.Println(" ft gyga")
	}
	id := res.InsertedID

	fmt.Println("id is ", id)
	return
}

func FindByUserName(userName string) []*Person {

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	option := options.Find()
	option.SetLimit(3)

	var results []*Person
	filter := bson.M{"userName": userName}
	fmt.Println("The filter is", userName)

	cur, err := db.Collection("person1").Find(ctx, filter, option)

	if err != nil {
		fmt.Println("error in db")
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var elem Person
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error in doc")
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	log.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results

}
