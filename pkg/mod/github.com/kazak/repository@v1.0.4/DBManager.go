package db

import (
	"context"
	"fmt"
	api "github.com/kazak/grpcapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository Used interface repository
type Repository struct {
	collection *mongo.Collection
}

// GetByID - Get single Port from db
func (repository *Repository) GetByID(portID string) *api.Port {
	filter := bson.D{{"portid", portID}}
	var seaPort *api.Port
	_ = repository.collection.FindOne(context.TODO(), filter).Decode(&seaPort)

	return seaPort
}

// GetAll - Select port object
func (repository *Repository) GetAll() []*api.Port {
	options := options.Find()
	options.SetLimit(100)
	var seaPorts []*api.Port
	cur, _ := repository.collection.Find(context.TODO(), bson.D{}, options)

	for cur.Next(context.TODO()) {
		var elem api.Port
		_ = cur.Decode(&elem)

		seaPorts = append(seaPorts, &elem)
	}
	return seaPorts
}

// Save - Saving port in DB
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

// Delete - Remove port from db
func (repository *Repository) Delete(portID string) bool {
	filter := bson.D{{"portid", portID}}
	res, err := repository.collection.DeleteOne(context.TODO(), filter)

	if err != nil || res.DeletedCount < 1 {
		return false
	}

	return true
}

// ConnectToDB - initialize connect to db
func ConnectToDB(MongoDBHost string, dbName string) *Repository{
	clientOptions := options.Client().ApplyURI(MongoDBHost)
	client, _ := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Connected error", err)
	}
	collection := client.Database(dbName).Collection("ports")

	repository := &Repository{
		collection: collection,
	}

	return repository
}