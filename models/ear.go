package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// EarData bindしてそのまま保存
type EarData struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	UserID int                `json:"user_id"`

	RightEarList []float64 `json:"right_ear_list"`
	RightEarT    float64   `json:"right_ear_t"`
	LeftEarList  []float64 `json:"left_ear_list"`
	LeftEarT     float64   `json:"left_ear_t"`
	Date         string    `json:"date" bson:"date"`
	// Environment string             `json:"environment"`
}

type GetEarDataRes struct {
	EarData []EarData `json:"earData"`
}
