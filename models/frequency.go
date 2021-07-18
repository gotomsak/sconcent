package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MaxFrequencyData 中身
type MaxFrequencyData struct {
	MaxBlink     float64            `json:"max_blink"`
	MaxFaceMove  float64            `json:"max_face_move"`
	EarID        primitive.ObjectID `json:"ear_id"`
	FacePointAll []interface{}      `json:"face_point_all"`
}

// MaxFrequency bindしてそのまま保存
type MaxFrequency struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	UserID           int                `json:"user_id"`
	MaxFrequencyData MaxFrequencyData   `json:"max_frequency_data"`
	Date             time.Time          `json:"date" bson:"date"`
	// Environment      string             `json:"environment"`
}

// MinFrequencyData 中身
type MinFrequencyData struct {
	MinBlink     float64            `json:"min_blink"`
	MinFaceMove  float64            `json:"min_face_move"`
	EarID        primitive.ObjectID `json:"ear_id"`
	FacePointAll []interface{}      `json:"face_point_all"`
}

// MinFrequency bindしてそのまま保存
type MinFrequency struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	UserID           int                `json:"user_id"`
	MinFrequencyData MinFrequencyData   `json:"min_frequency_data"`
	Date             time.Time          `json:"date" bson:"date"`
	// Environment      string             `json:"environment"`
}

// GetFrequencyResData 集中度計測時に返すデータ
type GetFrequencyResData struct {
	MaxFrequency []MaxFrequency `json:"max_frequency"`
	MinFrequency []MinFrequency `json:"min_frequency"`
}
