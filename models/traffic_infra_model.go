package models

import (
	"gopkg.in/mgo.v2/bson"
)

type TrafficInfraModel struct {
	ID                   bson.ObjectId `bson:"_id" json:"-"`
	RegistryNumber       string        `bson:"registryNumber" json:"registryNumber"`
	RegistryCreationDate string        `bson:"registryCreationDate" json:"registryCreationDate"`
	Name                 string        `bson:"name" json:"name"`
	Latitude             string        `bson:"latitude" json:"latitude"`
	Longitude            string        `bson:"longitude" json:"longitude"`
	LegalInfo            string        `bson:"legalInfo" json:"legalInfo"`
	Category             string        `bson:"category" json:"category"`
	CategoryCreationDate string        `bson:"categoryCreationDate" json:"categoryCreationDate"`
	CategoryRefreshDate  string        `bson:"categoryRefreshDate" json:"categoryRefreshDate"`
	RegistryCancelDate   string        `bson:"registryExpireDate" json:"registryExpireDate"`
	CancelDescription    string        `bson:"cancelDescription" json:"cancelDescription"`
}
