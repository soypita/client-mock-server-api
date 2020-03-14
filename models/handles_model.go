package models

import "gopkg.in/mgo.v2/bson"

type HandlesModel struct {
	ID       bson.ObjectId `bson:"_id" json:"-"`
	Group    string        `bson:"group" json:"group"`
	HandleId string        `bson:"handleId" json:"handleId"`
}
