package main

import (
	"go.mongodb.org/mongo-driver/bson"
)

func profileFromBsonToJson(profileBson bson.Raw) interface{}{
	type ProfileData struct{
		Username        string	`json:"username"`
		Age				string  `json:"age"`
		Country			string	`json:"country"`
		Status			string	`json:"status"`
		Description		string	`json:"description"`
	}

	profileData := ProfileData{ Username: profileBson.Lookup("username").StringValue(), 
								  Age: profileBson.Lookup("age").StringValue(), 
								  Country: profileBson.Lookup("country").StringValue(), 
								  Status: profileBson.Lookup("status").StringValue(),
								  Description: profileBson.Lookup("description").StringValue() }

	return profileData
}