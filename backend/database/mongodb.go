package database

import (
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

var users *mongo.Collection
var items *mongo.Collection
var ctx = context.TODO()

type User struct {
	Email          	string    	`bson:"email"`
	Username       	string    	`bson:"username"`
	HashedPassword 	[]byte    	`bson:"hashedPassword"`
	CreatedAt      	time.Time 	`bson:"createdAt"`
	SessionValue   	string    	`bson:"sessionValue"`
	DateOfBirth    	string 	 	`bson:"dateOfBirth"`
	Gender			string		`bson:"gender"`
	Country        	string    	`bson:"country"`
	Status         	string    	`bson:"status"`
	Description    	string    	`bson:"description"`
}

type Item struct{
	Username        		string		`bson:"username"`
	CreatedAt      			time.Time 	`bson:"createdAt"`
	ItemType				string		`bson:"itemType"`
	ItemId					string  	`bson:"itemId"`
	ItemName				string 		`bson:"itemName"`
	ItemImage				string		`bson:"itemImage"`
	Rating					int			`bson:"rating"`
	Comment					string		`bson:"comment"`
}

func Init() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	database := client.Database("my-universe")
	users = database.Collection("users")
	items = database.Collection("items")
	return nil
}

func InsertNewUser(email string, username string, hashedPassword []byte, dateOfBirth string, gender string, country string, status string, description string) error {
	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"email", email}},
				bson.D{{"username", username}},
			},
		},
	}
	cursor := users.FindOne(ctx, filter)
	_, err := cursor.Raw()
	if err == mongo.ErrNoDocuments {
		newUser := User{
			Email:          email,
			Username:       username,
			HashedPassword: hashedPassword,
			CreatedAt:      time.Now(),
			DateOfBirth:    dateOfBirth,
			Gender:			gender,
			Country:        country,
			Status:         status,
			Description:    description}
		_, err = users.InsertOne(ctx, newUser)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	} else {
		return errors.New("Username or email already exists")
	}
}

func GetPasswordByEmail(email string) ([]byte, error) {
	filter := bson.D{{"email", email}}
	cursor := users.FindOne(ctx, filter)
	result, err := cursor.Raw()
	if err != nil {
		return nil, errors.New("Email dosen't exist")
	}

	return result.Lookup("hashedPassword").Value, nil
}

func SetUserSessionValue(email string, sessionValue string) error {
	filter := bson.D{{"email", email}}
	update := bson.D{{"$set", bson.D{{"sessionValue", sessionValue}}}}
	_, err := users.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetProfileByUsername(username string) (bson.Raw, error) {
	filter := bson.D{{"username", username}}
	cursor := users.FindOne(ctx, filter)
	result, err := cursor.Raw()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProfileBySessionValue(value string) (bson.Raw, error) {
	filter := bson.D{{"sessionValue", value}}
	cursor := users.FindOne(ctx, filter)
	result, err := cursor.Raw()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func AddItem(username string, itemType string, itemId string, itemName string, itemImage string, rating int, comment string) error{
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
				bson.D{{"itemId", itemId}},
			},
		},
	}
	cursor := items.FindOne(ctx, filter)
	_, err := cursor.Raw()
	if err == mongo.ErrNoDocuments {
		newItem := Item{
			Username:   	username,
			CreatedAt:      time.Now(),
			ItemType:		itemType,
			ItemId: 		itemId,
			ItemName: 		itemName,
			ItemImage:  	itemImage,
			Rating:     	rating,
			Comment:    	comment}
		_, err = items.InsertOne(ctx, newItem)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	} else {
		update := bson.D{{"$set", bson.D{{"rating", rating}, {"comment", comment}}}}
		_, err = items.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		return nil
	}
}

func GetUserItems(username string) ([]Item, error){
	filter := bson.D{{"username", username}}
	cursor, err := items.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	
	var results []Item
   	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func DeleteUserItem(username string, itemId string) error{
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
				bson.D{{"itemId", itemId}},
			},
		},
	}

	_, err := items.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
