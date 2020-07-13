package models

import "gopkg.in/mgo.v2/bson"

type CargoModel struct {
	ID                   bson.ObjectId `bson:"_id" json:"-"`
	RegistryNumber       string        `bson:"registryNumber" json:"registryNumber"`
	RegistryCreationDate string        `bson:"registryCreationDate" json:"registryCreationDate"`
	Name                 string        `bson:"name" json:"name"`
	Category             string        `bson:"category" json:"category"`
}