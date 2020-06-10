package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongo_url = "mongodb+srv://huifer:huifer@cluster0-1boxu.azure.mongodb.net/big_data?authSource=admin&replicaSet=Cluster0-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))
	if err != nil {
	}
}
