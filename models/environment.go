package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Environment struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name"`
	UserID       int                `json:"user_id"`
	EarID        primitive.ObjectID `json:"ear_id"`
	MaxFreqID    primitive.ObjectID `json:"max_freq_id" bson:"max_freq_id"`
	MinFreqID    primitive.ObjectID `json:"min_freq_id" bson:"min_freq_id"`
	ReMaxFreqIDs []string           `json:"re_max_freq_ids"`
	ReMinFreqIDs []string           `json:"re_min_freq_ids"`
	Date         time.Time          `json:"date"`
}

type EnvironmentRes struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name"`
	UserID  int                `json:"user_id"`
	Ear     EarData            `json:"ear"`
	MaxFreq MaxFrequency       `json:"max_freq" bson:"max_freq"`
	MinFreq MinFrequency       `json:"min_freq" bson:"min_freq"`
	Date    time.Time          `json:"date"`
}

type GetEnvironmentRes struct {
	Environments []EnvironmentRes `json:"environments"`
}
