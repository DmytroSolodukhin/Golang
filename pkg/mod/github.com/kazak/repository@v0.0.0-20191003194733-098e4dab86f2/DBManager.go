package db

import (
	api "github.com/kazak/grpcapi"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

//Get single Port from db
func (repository *Repository) GetByID(portID string) *api.Port {
	filter := bson.D{{"portid", portID}}
	var seaPort *api.Port
	err := repository.collection.FindOne(context.TODO(), filter).Decode(&seaPort)
	if err != nil {
		log.Fatal(err)
	}
	return seaPort
}

//Select port object
func (repository *Repository) GetAll() []*api.Port {
	options := options.Find()
	options.SetLimit(100)
	var seaPorts []*api.Port
	cur, err := repository.collection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem api.Port
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		seaPorts = append(seaPorts, &elem)
	}
	return seaPorts
}

//Saving port in DB
func (repository *Repository) Save(seaPort *api.Port) bool {
	filter := bson.D{{"portid", seaPort.PortId}}
	update := bson.D{
		{"$set", seaPort},
	}
	_, err := repository.collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))

	if err != nil {
		return false
	}

	return true
}

//Remove port from db
func (repository *Repository) Delete(portID string) bool {
	filter := bson.D{{"portid", portID}}
	res, err := repository.collection.DeleteOne(context.TODO(), filter)

	if err != nil || res.DeletedCount < 1 {
		return false
	}

	return true
}

//Connecting to mongoDB
func ConnectToDB(MongoDBHost string, dbName string) *Repository{
	clientOptions := options.Client().ApplyURI(MongoDBHost)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Connected error", err)
	}
	collection := client.Database(dbName).Collection("ports")

	repository := &Repository{
		collection: collection,
	}

	return repository
}