package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/mahendrarathore1742/MongoDbApiInGo/model"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection = "Your Link"
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

func insertOneMovie(movie model.Netflix) {

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


func deleteOneMovie(movieId string){

	 id,_ := primitive.ObjectIDFromHex(movieId);

	 filter := bson.M{"_id" : id}
	 deleteCount ,_ :=  collection.DeleteOne(context.Background(),filter);

	 if err != nil{
	 	log.Fatal(err)
	 }

	 fmt.Println("Movie got deleted:", deleteCount);

}

// delere all record for database 


func deleteAllMovie() int64 {

	DeleteResult,err  := collection.DeleteMany(context.Background(),bson.D{{}},nil);

	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("Number of movie delete: ",DeleteResult.DeletedCount);
	return DeleteResult.DeletedCount

}


// get all Movies form database

func getAllMovies() []primitive.M { 


	cur,err   := collection.Find(context.Background(),bson.D{{}},nil);

	if err != nil{
		log.Fatal(err);
	}


	var movies []primitive.M;

	for cur.Next(context.Background()){
		var movie bson.M;

		err := cur.Decode(&movie);
		log.Fatal(err);
	}

	movies  = append(movies,movie);
	defer cur.Close(context.Background());
	return movies;

}

// Actual controller  - file

func GetMyAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header("Content-Type","application/x-www-form-urlencode");
	allMovies := getAllMovies();
	json.NewEncoder(w).Encode(allMovies);
}

func CreateMovie(w http.ResponseWriter,r *http.Request){

	w.Header("Content-Type","application/x-www-form-urlencode");
	w.Header("Allow-Control-Allow-Methods","POST");

	var movie model.Netflix
	_ = json.NewDecoder(w).Decode(movie);
	
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request){

	w.Header("Content-Type","application/x-www-form-urlencode");
	w.Header("Allow-Control-Allow-Methods","PUT");	

	params := mux.Vars(r);

	updateOneMovies(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){

	w.Header("Content-Type","application/x-www-form-urlencode");
	w.Header("Allow-Control-Allow-Methods","DELETE");	

	params := mux.Vars(r);
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"]);

}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request){

	w.Header("Content-Type","application/x-www-form-urlencode");
	w.Header("Allow-Control-Allow-Methods","DELETE");	

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count);

}