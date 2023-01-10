package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	Pnum string `bson:"pnum"`
}

func main() {
	MongoDB_URL := "mongodb://127.0.0.1:27017"
	if MongoDB_URL == "" {
		fmt.Println("Could not found MongoDB_URL")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDB_URL))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("go-mongo")
	personCollection := db.Collection("tPerson")

	// Create
	newPerson := Person{Name: "bewisesh91", Age: 32, Pnum: "01048304689"}
	insertResult, err := personCollection.InsertOne(context.TODO(), newPerson)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document insterted with ID: %s\n", insertResult.InsertedID)

	// Read
	targetName := "bewisesh91"
	filterForFind := bson.D{{Key: "name", Value: targetName}}
	var findResult bson.M
	err = personCollection.FindOne(context.TODO(), filterForFind).Decode(&findResult)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No Document was found with the name %s\n", targetName)
		return
	} else if err != nil {
		panic(err)
	}

	fmt.Println(findResult)
	if jsonData, err := json.MarshalIndent(findResult, "", "  "); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", jsonData)
	}

	filterNoOption := bson.D{}

	allPerson, err := personCollection.Find(context.TODO(), filterNoOption)
	if err != nil {
		panic(err)
	}

	var personArray []Person
	if err := allPerson.All(context.TODO(), &personArray); err != nil {
		panic(err)
	}

	for _, person := range personArray {
		allPerson.Decode(&person)
		output, err := json.MarshalIndent(person, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	// Update
	filterForUpdate := bson.D{{Key: "name", Value: "bewisesh91"}}
	update := bson.D{{"$set", bson.D{{Key: "age", Value: 31}}}}

	updateResult, err := personCollection.UpdateOne(context.TODO(), filterForUpdate, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document updated: %v\n", updateResult.ModifiedCount)

	// Delete

}
