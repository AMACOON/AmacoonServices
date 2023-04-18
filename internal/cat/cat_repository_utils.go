package cat

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var database = "amacoon"
var catsCollection = "cats"
var breedsCollection = "breeds"
var colorsCollection = "colors"
var federationsCollection = "federations"
var catteriesCollection = "catteries"
var ownersCollection = "owners"
var countriesCollection = "countries"

func BuildPipelineWithLookups(matchStage bson.D, lookups []string) mongo.Pipeline {
	pipeline := mongo.Pipeline{matchStage}

	for _, lookup := range lookups {
		switch lookup {
		case "breed":
			pipeline = append(pipeline, LookupBreedStage(), UnwindBreedStage())
		case "color":
			pipeline = append(pipeline, LookupColorStage(), UnwindColorStage())
		case "father":
			pipeline = append(pipeline, LookupFatherStage(), UnwindFatherStage(), AddFatherNameAndRemoveFatherStage())
		case "country":
			pipeline = append(pipeline, LookupCountryStage(), UnwindCountryStage())
		case "mother":
			pipeline = append(pipeline, LookupMotherStage(), UnwindMotherStage(), AddMotherNameAndRemoveFatherStage())
		case "cattery":
			pipeline = append(pipeline, LookupCatteryStage(), UnwindCatteryStage())
		case "owner":
			pipeline = append(pipeline, LookupOwnerStage(), UnwindOwnerStage())
		case "federation":
			pipeline = append(pipeline, LookupFederationStage(), UnwindFederationStage())
		}
	}

	return pipeline
}


func LookupFatherStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         catsCollection,
			"localField":   "fatherId",
			"foreignField": "_id",
			"as":           "father",
		},
	}}
}

func UnwindFatherStage() bson.D {
	return bson.D{{
		Key: "$unwind",
		Value: bson.M{
			"path":                       "$father",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}

func AddFatherNameAndRemoveFatherStage() bson.D {
	return bson.D{{
		Key: "$addFields",
		Value: bson.M{
			"fatherName": "$father.name",
			"father":     nil,
		},
	}}
}



func LookupCountryStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         countriesCollection,
			"localField":   "countryId",
			"foreignField": "_id",
			"as":           "country",
		},
	}}
}

func UnwindCountryStage() bson.D {
	return bson.D{{
		Key:   "$unwind",
		Value: bson.M{
			"path":                       "$country",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}

func LookupMotherStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         catsCollection,
			"localField":   "motherId",
			"foreignField": "_id",
			"as":           "mother",
		},
	}}
}

func UnwindMotherStage() bson.D {
	return bson.D{{
		Key: "$unwind",
		Value: bson.M{
			"path":                       "$mother",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}


func AddMotherNameAndRemoveFatherStage() bson.D {
	return bson.D{{
		Key: "$addFields",
		Value: bson.M{
			"motherName": "$mother.name",
			"mother":     nil,
		},
	}}
}

func LookupBreedStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         breedsCollection,
			"localField":   "breedId",
			"foreignField": "_id",
			"as":           "breed",
		},
	}}
}

func UnwindBreedStage() bson.D {
	return bson.D{{
		Key:   "$unwind",
		Value: bson.M{
			"path":                       "$breed",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}

func LookupColorStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         colorsCollection,
			"localField":   "colorId",
			"foreignField": "_id",
			"as":           "color",
		},
	}}
}

func UnwindColorStage() bson.D {
	return bson.D{{
		Key:   "$unwind",
		Value: bson.M{
			"path":                       "$color",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}

func LookupCatteryStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         catteriesCollection,
			"localField":   "catteryId",
			"foreignField": "_id",
			"as":           "cattery",
		},
	}}
}

func UnwindCatteryStage() bson.D {
	return bson.D{{
		Key:   "$unwind",
		Value: bson.M{
			"path":                       "$cattery",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}

func LookupOwnerStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         ownersCollection,
			"localField":   "ownerId",
			"foreignField": "_id",
			"as":           "owner",
		},
	}}
}

func UnwindOwnerStage() bson.D {
	return bson.D{{
		Key:   "$unwind",
		Value: bson.M{
			"path":                       "$owner",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}

func LookupFederationStage() bson.D {
	return bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         federationsCollection,
			"localField":   "federationId",
			"foreignField": "_id",
			"as":           "federation",
		},
	}}
}

func UnwindFederationStage() bson.D {
	return bson.D{{
		Key:   "$unwind",
		Value: bson.M{
			"path":                       "$federation",
			"preserveNullAndEmptyArrays": true,
		},
	}}
}
