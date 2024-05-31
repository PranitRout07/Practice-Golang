package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PranitRout07/Practice-Golang/basic_api_using_mongodb/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDB_URI = "mongodb://host.docker.internal:27017/"
	DB_name     = "Netflix"
	col_name    = "movies"
)

var collections *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(MongoDB_URI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection is established successfully..")
	collections = client.Database(DB_name).Collection(col_name)

}

func addMovies(movie models.Movie) {
	inserted, err := collections.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted data", inserted)
}

func updateMovie(id string, r *http.Request) {
	var updatedMovie models.Movie
	json.NewDecoder(r.Body).Decode(&updatedMovie)
	covertedMovie, err := toBsonM(updatedMovie)
	if err != nil {
		log.Fatal(err)
	}
	movieId,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}
	filter := bson.M{"_id": movieId}
	update := bson.M{"$set": covertedMovie}

	res, err := collections.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", res)

}

func getMovieAll() []primitive.M {
	cur, err := collections.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies

}

func deleteMovie(id string){
	movieID,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}
	filter := bson.M{"_id": movieID}
	deleteCount,err := collections.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie deleted with deletecount :",deleteCount)
}

func deleteAllMovies(){
	filter := bson.D{{}}
	collections.DeleteMany(context.Background(),filter,nil)
}

func toBsonM(data interface{}) (bson.M, error) {
	// Marshal the struct to BSON bytes
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal the BSON bytes to a bson.M map
	var bsonMap bson.M
	err = bson.Unmarshal(bsonBytes, &bsonMap)
	if err != nil {
		return nil, err
	}

	return bsonMap, nil
}

//controller

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type","application/json")
	movies := getMovieAll()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Send some data")
	}

	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	addMovies(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	updateMovie(params["id"], r)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	deleteMovie(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	deleteAllMovies()
}