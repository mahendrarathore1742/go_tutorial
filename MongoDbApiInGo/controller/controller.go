package controller;


import (	
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

const connection = "mongodb+srv://msr:msr12345@@cluster0.f5zbcgu.mongodb.net/?retryWrites=true&w=majority"
const DbName = "Netflix";

const colName = "watchlist"

//Most Important
var collection *mongo.Collection


//connect with mongoDB

func init(){

	//client option
	clientOption  := options.Client().ApplyURI(connection)


	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(),clientOption)


	if err != nil{
		log.Fatal(err)
	}


	fmt.Println("mongodb Connected success")

	collection = client.Database(DbName).Collection(colName)

	//Collection instanse
	fmt.Println("Collection instanse is ready")

}

