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
var profiles *mongo.Collection 
var ctx = context.TODO()

type User struct {
	Email           string  	`bson:"email"`
	Username        string		`bson:"username"`
	HashedPassword  []byte  	`bson:"hashedPassword"`
	CreatedAt		time.Time	`bson:"createdAt"`
	SessionValue	string		`bson:"sessionValue"`
	Age				string		`bson:"age"`
	Country			string		`bson:"country"`
	Status			string		`bson:"status"`
	Description		string		`bson:"description"`
}

func Init() error{
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
	profiles = database.Collection("profiles")
	return nil
}

func InsertNewUser(email string, username string, hashedPassword []byte) error{
	filter := bson.D{
		{"$or",
		   bson.A{
			bson.D{{"email", email}},
			bson.D{{"email", username}},
		   },
		},
	 }
	cursor := users.FindOne(ctx, filter)
	_, err:= cursor.Raw()
	if err == mongo.ErrNoDocuments {
		newUser := User {
			Email: email,
			Username: username,
			HashedPassword: hashedPassword, 
			CreatedAt: time.Now(),
			Age: "",
			Country: "", 
			Status: "", 
			Description: "" }
		_, err = users.InsertOne(ctx, newUser)
		if err != nil {
			return err
		}
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

func GetPasswordByEmail(email string) ([]byte, error){
	filter := bson.D{{"email", email}}
	cursor := users.FindOne(ctx, filter)
	result, err := cursor.Raw()
	if err != nil {
		return nil, errors.New("Email dosen't exist")
	}

	return result.Lookup("hashedPassword").Value, nil
}

func SetUserSessionValue(email string, sessionValue string) (error){
	filter := bson.D{{"email", email}}
	update := bson.D{{"$set", bson.D{{"sessionValue", sessionValue}}}}
	_, err := users.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetProfileByUsername(username string) (bson.Raw, error){
	filter := bson.D{{"username", username}}
	cursor := users.FindOne(ctx, filter)
	result, err := cursor.Raw()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProfileBySessionValue(value string) (bson.Raw, error){
	filter := bson.D{{"sessionValue", value}}
	cursor := users.FindOne(ctx, filter)
	result, err := cursor.Raw()
	if err != nil {
		return nil, err
	}

	return result, nil
}