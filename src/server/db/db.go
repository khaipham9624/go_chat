package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

const DbName = "chat_app"

func GetClient() *mongo.Client {
	if client != nil {
		return client
	}
	client, err := Connect("mongodb://localhost:27017/")
	if err != nil {
		return nil
	}
	Ping(client, DbName)
	return client
}

func Connect(uri string) (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	var err error
	client, err = mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	return client, err
}

func Ping(client *mongo.Client, dbName string) {
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func InsertOne(client *mongo.Client, dbName, collection string, data bson.M) (*mongo.InsertOneResult, error) {
	insertResult, err := client.Database(dbName).Collection(collection).InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func InsertMany(client *mongo.Client, dbName, collection string, data []bson.D) (*mongo.InsertManyResult, error) {
	insertManyResult, err := client.Database(dbName).Collection(collection).InsertMany(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return insertManyResult, nil
}

func FindOne(client *mongo.Client, dbName, collection string, filter bson.M) (bson.M, error) {
	var result bson.M
	err := client.Database(dbName).Collection(collection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindMany(client *mongo.Client, dbName, collection string, filter bson.M) ([]bson.M, error) {
	var result []bson.M
	cursor, err := client.Database(dbName).Collection(collection).Find(context.TODO(), filter)
	if err != nil {
		return result, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateOne(client *mongo.Client, dbName, collection string, filter, update bson.M) (*mongo.UpdateResult, error) {
	updateResult, err := client.Database(dbName).Collection(collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func UpdateMany(client *mongo.Client, dbName, collection string, filter, update bson.M) (*mongo.UpdateResult, error) {
	updateResult, err := client.Database(dbName).Collection(collection).UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeleteOne(client *mongo.Client, dbName, collection string, filter bson.M) (*mongo.DeleteResult, error) {
	deleteResult, err := client.Database(dbName).Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}

func DeleteMany(client *mongo.Client, dbName, collection string, filter bson.M) (*mongo.DeleteResult, error) {
	deleteResult, err := client.Database(dbName).Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
