package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/drver/mongocrypt/options"
	"log"
)

const connection = "mongodb+srv://msr:msr12345@@cluster0.f5zbcgu.mongodb.net/?retryWrites=true&w=majority"
const DbName = "Netflix"

const colName = "watchlist"

// Most Important
var collection *mongo.Collection

//connect with mongoDB

func init() {

	//client option
	clientOption := options.Client().ApplyURI(connection)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
 
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongodb Connected success")

	collection = client.Database(DbName).Collection(colName)

	//Collection instanse 
	fmt.Println("Collection instanse is ready")

}

func insertOneMovie(movie model.netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	} 

	fmt.Println(inserted.InsertedID)
}

//update ine record

func updateOneMovies(movieId string){

	id,_ := primitive.ObjectIDFromHex(movieId);
	filter := bson.M{"_id": id};
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

fmt.Println("modified count: ", result.ModifiedCount)

}
