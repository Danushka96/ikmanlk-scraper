package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var adCollection *mongo.Collection

func init() {
	adCollection = GetDatabase().Collection("ads")
}

func SaveAd(ad IkmanAd) (*mongo.InsertOneResult, error) {
	return adCollection.InsertOne(context.Background(), ad)
}

func ExistAd(ad IkmanAd) bool {
	filter := bson.D{{"_id", ad.ID}}
	count, countError := adCollection.CountDocuments(context.TODO(), filter)
	if countError != nil {
		panic(countError)
	}
	return count > 0
}
