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
	FaceAngleAll []interface{}      `json:"face_angle_all"`
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
	FaceAngleAll []interface{}      `json:"face_angle_all"`
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

// type ReMaxFrequencyBind struct {
// 	UserID           int                `json:"user_id"`
// 	MaxFrequencyData MaxFrequencyData   `json:"max_frequency_data"`
// 	EnvironmentID    primitive.ObjectID `json:"environment_id"`
// 	RootMaxFreqID    primitive.ObjectID `json:"root_max_freq_id"`
// 	Date             time.Time          `json:"date" bson:"date"`
// }

type ReMaxFrequencySave struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	SeparationNum    int                `json:"separation_num"`
	UserID           int                `json:"user_id"`
	MaxFrequencyData MaxFrequencyData   `json:"max_frequency_data"`
	EnvironmentID    primitive.ObjectID `json:"environment_id"`
	RootMaxFreqID    primitive.ObjectID `json:"root_max_freq_id"`
	Date             time.Time          `json:"date" bson:"date"`
}

type ReMinFrequencySave struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	UserID           int                `json:"user_id"`
	MinFrequencyData MinFrequencyData   `json:"min_frequency_data"`
	EnvironmentID    primitive.ObjectID `json:"environment_id"`
	RootMinFreqID    primitive.ObjectID `json:"root_min_freq_id"`
	Date             time.Time          `json:"date" bson:"date"`
}
