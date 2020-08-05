package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Souvenir is a Model.
type Souvenir struct {
	//ID          string `json:"_id" bson:"_id"`
	Nombre      string `json:"nombre" bson:"nombre,omitempty"`
	Categoria   string `json:"categoria" bson:"categoria,omitempty"`
	Descripcion string `json:"descripcion" bson:"descripcion,omitempty"`
	Precio      string `json:"precio" bson:"precio,omitempty"`
	Stock       string `json:"stock" bson:"stock,omitempty"`
	Fecha       string `json:"fecha" bson:"fecha,omitempty"`
}

// SouvenirsMDB collecion of souvenirs MongoDB
var SouvenirsMDB *mongo.Collection

// ConnectMongoDB conexion a MongoDB
func ConnectMongoDB() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://nicolas17197:Qi5IKFhHo9oyUQLy@cluster0-xyuut.gcp.mongodb.net/AikenColores?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	SouvenirsMDB = client.Database("AikenColores").Collection("Souvenirs")
	fmt.Println("Connected to MongoDB!")
}

// GetSouvenirs get all souvenirs
func GetSouvenirs(w http.ResponseWriter, req *http.Request) {

	filter := bson.D{}

	var results []*Souvenir
	//
	cur, err := SouvenirsMDB.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var s Souvenir
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	json.NewEncoder(w).Encode(results)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}

// GetSouvenir get one souvenir by id
func GetSouvenir(w http.ResponseWriter, req *http.Request) {
	/*
		params := mux.Vars(req)
		for _, item := range souvenirs {
			if item._id == params["_id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		json.NewEncoder(w).Encode(&Souvenir{})
	*/
}

// CreateSouvenir create a souvenir
func CreateSouvenir(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	var souvenir Souvenir
	_ = json.NewDecoder(req.Body).Decode(&souvenir)
	fmt.Println("OBJECTO ", souvenir)
	insertResult, err := SouvenirsMDB.InsertOne(context.TODO(), souvenir)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(souvenir)
	fmt.Println("Inserted multiple documents: ", insertResult)
}

// DeleteSouvenir delete a souvenir by id
func DeleteSouvenir(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	//filter :=
	//fmt.Println(filter)
	idPrimitive, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	deleteResult, err := SouvenirsMDB.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DELETE ? %v", deleteResult.DeletedCount)
}

// UpdateSouvenir modificar un souvenir.
func UpdateSouvenir(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	idPrimitive, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		fmt.Print("Error de ID")
		log.Fatal(err)
	}
	var souvenir Souvenir
	_ = json.NewDecoder(req.Body).Decode(&souvenir)
	//fmt.Println(reflect.TypeOf(souvenir))

	filter := bson.M{"_id": idPrimitive}
	update := bson.M{"nombre": souvenir.Nombre}
	fmt.Println(souvenir.Nombre)
	result := SouvenirsMDB.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update})
	if result.Err() != nil {
		result.Err()
	}
}

// Inicio get.
func Inicio(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola"))
}
