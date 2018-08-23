package repository

import (
	"context"
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

type repository struct {
	connectionString string
}

func NewRepository(connectionString string) *repository {
	r := repository{}
	return &r
}

func (r *repository) GetItems(count int) ([]string, error) {

	collection, err := r.getCollection()

	if err != nil {
		return []string{}, err
	}

	elem := bson.NewDocument()
	collection.FindOne(context.Background(), bson.NewDocument()).Decode(elem)

	value := elem.Lookup("value")

	if value == nil {
		return nil, errors.New("Failed to fetch items")
	}

	return []string{value.Interface().(string)}, nil
}

func (r *repository) GetItem(id string) (string, error) {
	collection, err := r.getCollection()

	if err != nil {
		return "", err
	}

	objectId, err := objectid.FromHex(id)

	if err != nil {
		return "", err
	}

	filter := bson.NewDocument(bson.EC.ObjectID("_id", objectId))
	elem := bson.NewDocument()

	collection.FindOne(context.Background(), filter).Decode(elem)

	value := elem.Lookup("value")

	return value.Interface().(string), nil
}

func (r *repository) SetItem(s string) (string, error) {
	collection, err := r.getCollection()

	if err != nil {
		return "", err
	}

	insertResult, err := collection.InsertOne(context.Background(), map[string]string{"value": s})

	if err != nil {
		return "", err
	}

	objectId, success := insertResult.InsertedID.(objectid.ObjectID)

	if !success {
		return "", errors.New("Failed insert to db.")
	}

	log.Println(objectId.Hex())
	// log.Println("Stored " + s + " in id " + insertResult.InsertedID)
	return objectId.Hex(), nil
}

func (r *repository) getCollection() (*mongo.Collection, error) {
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	collection := client.Database("TestDb").Collection("testcollection")

	return collection, nil
}
