package models

import (
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveConcentration postされてきたdataのbind
type SaveConcentrationBind struct {
	Type          string             `json:"type"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FacePointAll  primitive.ObjectID `json:"face_point_id" bson:"_face_point_id"`
	UserID        int                `json:"user_id"`
	Measurement   string             `json:"measurement"`
	Concentration []interface{}      `json:"concentration"`
}

// GetSaveImagesID postされてきたdataのbind
type GetSaveImagesIDBind struct {
	Type string `json:"type"`
}

// GetID postされてきたdataのbind
type GetIDBind struct {
	Type   string `json:"type"`
	UserID int    `json:"user_id"`
	// ID       primitive.ObjectID `json:"id" bson:"_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Concentration []interface{} `json:"concentration"`
}

// GetIDSave 保存
type GetIDSave struct {
	Type         string             `json:"type"`
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FacePointAll primitive.ObjectID `json:"face_point_id" bson:"_face_point_id"`
	UserID       int                `json:"user_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Concentration []interface{} `json:"concentration"`
}

// GetIDLog IDを取得したときの履歴
type GetIDLog struct {
	gorm.Model
	UserID     int    `json:"user_id" gorm:"not null"`
	ConcDataID string `json:"conc_data_id" bson:"conc_data_id" gorm:"null"`
}

// GetIDRes レスポンス
type GetIDRes struct {
	ConcDataID  primitive.ObjectID `json:"conc_id" bson:"_conc_id"`
	FacePointID primitive.ObjectID `json:"face_point_id" bson:"_face_point_id"`
}

type GetFacePointIDSave struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FacePointAll []interface{}      `json:"face_point_all" bson:"face_point_all"`
}

type PostFacePointSave struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FacePointAll []interface{}      `json:"face_point_all" bson:"face_point_all"`
}

// // GetSaveImagesIDRes レスポンス
// type GetSaveImagesIDRes struct {
// 	ConcDataID int64 `json:"id"`
// }
