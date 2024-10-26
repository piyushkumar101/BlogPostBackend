package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golangblogbackend/database"
	"golangblogbackend/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.OpenCollection(database.Client, "blogposts")

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := getAllPosts()
	json.NewEncoder(w).Encode(payload)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	post, err := getPostByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.BlogPost
	json.NewDecoder(r.Body).Decode(&post)
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	insertOnePost(post)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post models.BlogPost
	json.NewDecoder(r.Body).Decode(&post)
	post.UpdatedAt = time.Now()
	updatePost(params["id"], post)
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteOnePost(params["id"])
}

func getAllPosts() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	cur.Close(context.Background())
	return results
}

func getPostByID(id string) (primitive.M, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	var post bson.M
	if err := collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
		return nil, err
	}
	return post, nil
}

func insertOnePost(post models.BlogPost) {
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
}

func updatePost(id string, post models.BlogPost) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": post}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteOnePost(id string) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
