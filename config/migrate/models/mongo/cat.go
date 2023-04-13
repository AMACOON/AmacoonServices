package mongomodels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatTemp struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CatID      int                `bson:"id_gatos"`
	Registry   string             `bson:"registro"`
	CatName    string             `bson:"nome_do_gato"`
	FatherName string             `bson:"nome_do_pai"`
	MotherName string             `bson:"nome_da_mae"`
}


type CatTempId struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CatID      int                `bson:"id_gatos"`
	Registro   string             `bson:"registro"`
	CatName    string             `bson:"nome_do_gato"`
	FatherName string             `bson:"nome_do_pai"`
	MotherName string             `bson:"nome_da_mae"`
	FatherID   primitive.ObjectID `bson:"fatherID"`
	MotherID   primitive.ObjectID `bson:"motherID"`
}

type CatTempFull struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CatID      int                `bson:"id_gatos"`
	Registry   string             `bson:"registro"`
	CatName    string             `bson:"nome_do_gato"`
	FatherName string             `bson:"nome_do_pai"`
	MotherName string             `bson:"nome_da_mae"`
	FatherID   primitive.ObjectID `bson:"fatherID"`
	MotherID   primitive.ObjectID `bson:"motherID"`
}

type Cat struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

