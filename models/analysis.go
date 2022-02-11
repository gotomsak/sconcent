package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// GetRecAllRes userの集中度等のデータをすべて返す（facepoint以外）
type GetRecAllRes struct {
	MaxFrequency  []MaxFrequency        `json:"maxFrequency"`
	MinFrequency  []MinFrequency        `json:"minFrequency"`
	Concentration []GetConcentrationRes `json:"concentration"`
}

type GetRecUserDateRes struct {

	// MaxFrequency  []MaxFrequency      `json:"maxFrequency"`
	// MinFrequency  []MinFrequency      `json:"minFrequency"`
	// GetFrequencyResData GetFrequencyResData `json:"requencys"`
	// GetEarDataRes       GetEarDataRes       `json:"ears"`
	GetEnvironmentRes   []EnvironmentRes    `json:"environments"`
	GetConcentrationRes GetConcentrationRes `json:"concentration"`
	FacePointAll        PostFacePointSave   `json:"facePointAll"`
}

type GetFacePointRes struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FacePointAll []interface{}      `json:"face_point_all" bson:"face_point_all"`
}

type GetSelectAnswerResultSection struct {
	Concentration       GetConcentrationRes `json:"concentration"`
	AnswerResultSection AnswerResultSection `json:"answer_result_section"`
	Questionnaire       Questionnaire       `json:"questionnaire"`
}
