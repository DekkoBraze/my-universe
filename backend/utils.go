package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"log"
	"strconv"
)

func profileFromBsonToJson(profileBson bson.Raw) interface{}{
	type ProfileData struct{
		Username        string	`json:"username"`
		Age				string 	`json:"age"`
		Country			string	`json:"country"`
		Status			string	`json:"status"`
		Description		string	`json:"description"`
	}

	
	dateOfBirth, err := timeFromStringToTime(profileBson.Lookup("dateOfBirth").StringValue())
	if err != nil {
		log.Print("profileFromBsonToJson", err)
	}
	age, _ := monthYearDiff(dateOfBirth, time.Now())
	profileData := ProfileData{ Username: profileBson.Lookup("username").StringValue(), 
								  Age: strconv.Itoa(age), 
								  Country: profileBson.Lookup("country").StringValue(), 
								  Status: profileBson.Lookup("status").StringValue(),
								  Description: profileBson.Lookup("description").StringValue() }

	return profileData
}

func timeFromStringToTime(timeString string) (time.Time, error) {
	dateOfBirthInTime, err := time.Parse(time.DateOnly, timeString)
	if err != nil {
		return time.Time{}, err
	}

	return dateOfBirthInTime, err
}


func monthYearDiff(a, b time.Time) (years, months int) {
    m := a.Month()
    for a.Before(b) {
        a = a.Add(time.Hour * 24)
        m2 := a.Month()
        if m2 != m {
            months++
        }
        m = m2
    }
    years = months / 12
    months = months % 12
    return
}