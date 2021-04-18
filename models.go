package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveConcentration postされてきたdataのbind
type SaveConcentrationBind struct {
	Type          string             `json:"type"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Measurement   string             `json:"measurement"`
	Concentration []interface{}      `json:"concentration"`
}

// GetSaveImagesID postされてきたdataのbind
type GetSaveImagesIDBind struct {
	Type string `json:"type"`
}

// GetID postされてきたdataのbind
type GetIDBind struct {
	Type string `json:"type"`
	// ID       primitive.ObjectID `json:"id" bson:"_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Concentration []interface{} `json:"concentration"`
}

// GetIDSave 保存
type GetIDSave struct {
	Type string             `json:"type"`
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Concentration []interface{} `json:"concentration"`
}

// GetIDRes レスポンス
type GetIDRes struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

// GetSaveImagesIDRes レスポンス
type GetSaveImagesIDRes struct {
	ID uint64 `json:"id"`
}
