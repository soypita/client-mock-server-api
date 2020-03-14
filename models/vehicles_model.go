package models

import "gopkg.in/mgo.v2/bson"

type VehiclesModel struct {
	ID                   bson.ObjectId `bson:"_id" json:"-"`
	RegistryNumber       string        `bson:"registryNumber" json:"registryNumber"`
	RegistryCreationDate string        `bson:"registryCreationDate" json:"registryCreationDate"`
	Type                 string        `bson:"type" json:"type"`
	Model                string        `bson:"model" json:"model"`
	VIN                  string        `bson:"vin" json:"vin"`
	AddressComp          string        `bson:"companyAddress" json:"companyAddress"`
	AddressFiz           string        `bson:"fizAddress" json:"fizAddress"`
	OrgForm              string        `bson:"orgForm" json:"orgForm"`
	RegistrationId       string        `bson:"registrationId" json:"registrationId"`
	EgrulCreationDate    string        `bson:"egrulCreationDate" json:"egrulCreationDate"`
	Category             string        `bson:"category" json:"category"`
	CategoryCreationDate string        `bson:"categoryCreationDate" json:"categoryCreationDate"`
	RegistryCancelDate   string        `bson:"registryExpireDate" json:"registryExpireDate"`
	CancelDescription	 string			`bson:"cancelDescription" json:"cancelDescription"`
}
