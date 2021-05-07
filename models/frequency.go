package models

// MaxFrequencyData 中身
type MaxFrequencyData struct {
	MaxBlink     float64       `json:"max_blink"`
	MaxFaceMove  float64       `json:"max_face_move"`
	FacePointAll []interface{} `json:"face_point_all"`
}

// MaxFrequency bindしてそのまま保存
type MaxFrequency struct {
	UserID           int              `json:"user_id"`
	MaxFrequencyData MaxFrequencyData `json:"max_frequency_data"`
	Environment      string           `json:"environment"`
}

// MinFrequencyData 中身
type MinFrequencyData struct {
	MinBlink     float64       `json:"min_blink"`
	MinFaceMove  float64       `json:"min_face_move"`
	FacePointAll []interface{} `json:"face_point_all"`
}

// MinFrequency bindしてそのまま保存
type MinFrequency struct {
	UserID           int              `json:"user_id"`
	MinFrequencyData MinFrequencyData `json:"min_frequency_data"`
	Environment      string           `json:"environment"`
}

type GetFrequencyResData struct {
	MaxFrequency []MaxFrequency `json:"max_frequency"`
	MinFrequency []MinFrequency `json:"min_frequency"`
}
