package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		UserID     bson.ObjectId `json:"userid" xml:"userid" bson:"userid"`
		Name       string        `json:"name" xml:"name" form:"name" query:"name" bson:"name"`
		Email      string        `json:"email" xml:"email" form:"email" query:"email" bson:"email"`
		CreateDate string        `json:"createDate" bson:"createDate"`
	}
)
