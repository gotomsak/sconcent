package models

import (
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Concentration 集中度の手法に関するデータ
type Concentration struct {
	C1            []float64          `json:"c1" bson:"c1"`
	C2            []float64          `json:"c2" bson:"c2"`
	C3            []float64          `json:"c3" bson:"c3"`
	W             []float64          `json:"W" bson:"w"`
	Date          []string           `json:"date" bson:"date"`
	EarID         primitive.ObjectID `json:"ear_id" bson:"ear_id"`
	MaxFreqID     primitive.ObjectID `json:"max_freq_id" bson:"max_freq_id"`
	MinFreqID     primitive.ObjectID `json:"min_freq_id" bson:"min_freq_id"`
	EnvironmentID primitive.ObjectID `json:"environment_id"`
	FacePointAll  primitive.ObjectID `json:"face_point_id" bson:"_face_point_id"`
}

// SaveConcentration postされてきたdataのbind
type SaveConcentrationBind struct {
	Type          string             `json:"type"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        int                `json:"user_id"`
	Measurement   string             `json:"measurement"`
	Work          string             `json:"work"`
	Memo          string             `json:"memo"`
	Concentration Concentration      `json:"concentration" bson:"concentration"`
}

// GetConcentrationRes get_rec_all時のユーザーの集中度情報をすべて返す
type GetConcentrationRes struct {
	Type          string             `json:"type"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        int                `json:"user_id"`
	Measurement   string             `json:"measurement"`
	Work          string             `json:"work"`
	Memo          string             `json:"memo"`
	Concentration Concentration      `json:"concentration"`
}

// GetSaveImagesID postされてきたdataのbind
type GetSaveImagesIDBind struct {
	Type string `json:"type"`
}

// GetID get_id時にpostされてきたdataのbind
type GetIDBind struct {
	Type   string `json:"type"`
	UserID int    `json:"user_id"`
	// ID       primitive.ObjectID `json:"id" bson:"_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Work          string        `json:"work"`
	Memo          string        `json:"memo"`
	Concentration Concentration `json:"concentration"`
}

// GetIDSave 保存
type GetIDSave struct {
	Type string             `json:"type"`
	ID   primitive.ObjectID `json:"id" bson:"_id"`

	UserID int `json:"user_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Work          string        `json:"work"`
	Memo          string        `json:"memo"`
	Concentration Concentration `json:"concentration"`
}

// GetIDLog IDを取得したときの履歴
type GetIDLog struct {
	gorm.Model
	UserID     int    `json:"user_id" gorm:"not null"`
	ConcDataID string `json:"conc_data_id" bson:"conc_data_id" gorm:"null"`
}

// GetIDRes get_idのレスポンス
type GetIDRes struct {
	ConcDataID primitive.ObjectID `json:"conc_id" bson:"_conc_id"`
	// EarData     EarData            `json:"earData"`
	FacePointID primitive.ObjectID `json:"face_point_id" bson:"_face_point_id"`
}

// GetFacePointIDSave get_id時にidを生成して保存
type GetFacePointIDSave struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FacePointAll []interface{}      `json:"face_point_all" bson:"face_point_all"`
}

// PostFacePointSave facepointの保存
type PostFacePointSave struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FacePointAll []interface{}      `json:"face_point_all" bson:"face_point_all"`
}

// PostConcentSplitBind
type PostConcentSplitBind struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Type          string             `json:"type"`
	Measurement   string             `json:"measurement"`
	Concentration Concentration      `json:"concentration" bson:"_id"`
}

// // GetSaveImagesIDRes レスポンス
// type GetSaveImagesIDRes struct {
// 	ConcDataID int64 `json:"id"`
// }
