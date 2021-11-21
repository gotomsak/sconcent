package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetQuestionIDs 解く問題のIDと解いた問題のIDのstruct
type GetQuestionIDs struct {
	QuestionIDs []uint `json:"question_ids"`
	SolvedIDs   []uint `json:"solved_ids"`
}

// Question questionテーブルのstruct
type Question struct {
	gorm.Model
	Question    string `json:"question"`
	QimgPath    string `json:"qimg_path"`
	Mistake1    string `json:"mistake1"`
	Mistake2    string `json:"mistake2"`
	Mistake3    string `json:"mistake3"`
	Ans         string `json:"ans"`
	MimgPath1   string `json:"mimg_path1"`
	MimgPath2   string `json:"mimg_path2"`
	MimgPath3   string `json:"mimg_path3"`
	AimgPath    string `json:"aimg_path"`
	Season      string `json:"season"`
	QuestionNum string `json:"question_num"`
	Genre       string `json:"genre"`
}

// QuestionSend クライアントに送信する問題のstruct
type QuestionSend struct {
	QuestionID  uint     `json:"question_id"`
	Question    string   `json:"question"`
	QimgPath    []string `json:"qimg_path"`
	AnsList     []string `json:"ans_list"`
	AimgList    []string `json:"aimg_list"`
	Season      string   `json:"season"`
	QuestionNum string   `json:"question_num"`
	Genre       string   `json:"genre"`
}

// AnswerResultSectionIDSend クライアントに送信するsectionID
type AnswerResultSectionIDSend struct {
	AnswerResultSectionID uint `json:"answer_result_section_id"`
}

// AnswerResultSend checkAnswerのレスポンスをまとめたstruct
type AnswerResultSend struct {
	AnswerResultID uint   `json:"answer_result_id"`
	Result         string `json:"result"`
	Answer         string `json:"answer"`
}

// AnswerResult 解答の結果を保存するテーブルのstruct
type AnswerResult struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	UserAnswer   string `gorm:"not null"` // userの選んだ答え
	AnswerResult string `gorm:"not null"` // correctかincorrect
	// ConcentrationData string
	MemoLog          string    `gorm:"type:text;"`
	OtherFocusSecond uint      `json:"other_focus_second"`
	QuestionID       uint      `gorm:"not null"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
}

// AnswerResultSection 解答の結果のまとめを保存するテーブルのstruct
type AnswerResultSection struct {
	gorm.Model
	UserID              uint      `json:"user_id" gorm:"not null"`
	AnswerResultIDs     string    `json:"answer_result_ids"`
	SelectQuestionID    string    `json:"select_question_id"`
	CorrectAnswerNumber uint      `json:"correct_answer_number" gorm:"not null"`
	ConcID              string    `json:"conc_id"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
}

// CheckAnswerSection sectionごとの問題が送信されてきた時のbindのstruct
type CheckAnswerSectionBind struct {
	UserID              uint     `json:"user_id" gorm:"not null"`
	AnswerResultIDs     []uint64 `json:"answer_result_ids" gorm:"type:text;not null"`
	ConcID              string   `json:"conc_id"`
	SelectQuestionID    string   `json:"select_question_id"`
	CorrectAnswerNumber uint     `json:"correct_answer_number" gorm:"not null"`
	StartTime           string   `json:"start_time"`
	EndTime             string   `json:"end_time"`
}

// CheckAnswer postされてきたdataのbind
type CheckAnswerBind struct {
	UserID           uint   `json:"user_id"`
	UserAnswer       string `json:"user_answer"`
	MemoLog          string `json:"memo_log"`
	OtherFocusSecond uint   `json:"other_focus_second"`
	QuestionID       uint   `json:"question_id"`
	// ConcentrationData []interface{} `json:"concentration_data"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// Questionnaire アンケート結果を保存するテーブルのstruct
type Questionnaire struct {
	gorm.Model
	AnswerResultSectionID uint `json:"answer_result_section_id"`
	UserID                uint `json:"user_id" gorm:"not null"`
	Concentration         int  `json:"concentration"` // 集中
	WhileDoing            int  `json:"while_doing"`   // しながら
	Cheating              int  `json:"cheating"`      // カンニング
	Nonsense              int  `json:"nonsense"`      // デタラメ
}

type Genre struct {
	gorm.Model
	Genre string `json:"genre"`
}

type Season struct {
	gorm.Model
	Season string `json:"season"`
}

// Question questionテーブルのstruct
type OnlyQuestion struct {
	gorm.Model
	Question    string `json:"question" gorm:"type:text"`
	QimgPath    string `json:"qimg_path"`
	Mistake1    string `json:"mistake1"`
	Mistake2    string `json:"mistake2"`
	Mistake3    string `json:"mistake3"`
	Ans         string `json:"ans"`
	MimgPath1   string `json:"mimg_path1"`
	MimgPath2   string `json:"mimg_path2"`
	MimgPath3   string `json:"mimg_path3"`
	AimgPath    string `json:"aimg_path"`
	SeasonID    uint   `json:"season_id"`
	Season      Season `gorm:"foreignKey:SeasonID"`
	QuestionNum string `json:"question_num"`
	GenreID     uint   `json:"genre_id"`
	Genre       Genre  `gorm:"foreignKey:GenreID"`
}

type SelectQuestion struct {
	gorm.Model
	SelectQuestionName string `json:"select_question_name"`
	SelectQuestionIDs  string `json:"select_question_ids"`
}

// Results answer_result_idsをnosqlに保存
type AnswerResultIDs struct {
	AnswerResultIDs []uint64 `json:"answer_result_ids"`
}

// SelectQuestionIDs select_question_idsをmongoに保存
type SelectQuestionIDs struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SelectQuestionIDs []uint             `json:"select_question_ids"`
}

type GetSelectQuestionRes struct {
	SelectQuestion []SelectQuestion `json:"select_question"`
}

type GetQuestionGymBind struct {
	NowLevel int `json:"now_level"`
}
type GetQuestionGymRes struct {
	QuestionID uint     `json:"question_id"`
	Question   string   `json:"question"`
	AnsList    []string `json:"ans_list"`
}

type QuestionsSub struct {
	ID       uint   `json:"id"`
	Question string `json:"question"`
	Mistake1 string `json:"mistake1"`
	Mistake2 string `json:"mistake2"`
	Mistake3 string `json:"mistake3"`
	Ans      string `json:"ans"`
}
