package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

var users *mongo.Collection 
var ctx = context.TODO()

type User struct {
	Email           string  	`bson:"email"`
	Username        string		`bson:"username"`
	HashedPassword  []byte  	`bson:"hashedPassword"`
	CreatedAt		time.Time	`bson:"createdAt"`
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
	return nil
}

func InsertNewUser(email string, username string, hashedPassword []byte) error{
	isEmailBusy, err := IsEmailExists(email)
	if err != nil {
		return err
	}
	if isEmailBusy {
		return errors.New("email already exists")
	}

	isUsernameBusy, err := IsUsernameExists(username)
	if err != nil {
		return err
	}
	if isUsernameBusy {
		return errors.New("username already exists")
	}

	newUser := User {
		Email: email,
		Username: username,
		HashedPassword: hashedPassword, 
		CreatedAt: time.Now() }
	_, err = users.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

func IsEmailExists(email string) (bool, error) {
	filter := bson.D{{"email", email}}
	cursor := users.FindOne(ctx, filter)
	_, err:= cursor.Raw()
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func IsUsernameExists(username string) (bool, error) {
	filter := bson.D{{"username", username}}
	cursor := users.FindOne(ctx, filter)
	_, err:= cursor.Raw()
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func GetUser(email string) {

}