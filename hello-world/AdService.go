package main

import (
	"context"
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
	var currentAd IkmanAd
	findErr := adCollection.FindOne(context.TODO(), ad.ID).Decode(currentAd)
	return findErr == nil
}
