package models

// GetRecAllRes userの集中度等のデータをすべて返す（facepoint以外）
type GetRecAllRes struct {
	MaxFrequency  []MaxFrequency        `json:"maxFrequency"`
	MinFrequency  []MinFrequency        `json:"minFrequency"`
	Concentration []GetConcentrationRes `json:"concentration"`
}
