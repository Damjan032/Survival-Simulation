package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoDocumentDto struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	MainParams MainParams         `bson:"MainParams "json:"MainParams"`
	Epoches    []*Epoch           `bson:"Epoches "json:"Epoches"`
}

type MainParams struct {
	Id                  primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	StartDate           string             `bson:"startDate" json:"startDate"`
	NumberOfFoodSources int                `bson:"NumberOfFoodSources" json:"NumberOfFoodSources"`
	FinalEpoch          int                `bson:"FinalEpoch" json:"FinalEpoch"`
	Level               string             `bson:"Level" json:"Level"`
}
