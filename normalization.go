package main

import (
	"fmt"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
)

func Normalization() {
	db := utils.SqlConnect()
	defer db.Close()

	rows, err := db.Model(&models.Question{}).Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var question models.Question

		genre := models.Genre{}
		season := models.Season{}

		db.ScanRows(rows, &question)

		resGenre := db.First(&genre, "genre = ?", question.Genre).Error
		if resGenre != nil {
			genre.Genre = question.Genre
			resGenreCreate := db.Create(&genre).Error
			if resGenreCreate != nil {
				fmt.Println(resGenreCreate)
			}
		}

		resSeason := db.First(&season, "season = ?", question.Season).Error
		if resSeason != nil {
			season.Season = question.Season
			resSeasonCreate := db.Create(&season).Error
			if resSeasonCreate != nil {
				fmt.Println(resSeasonCreate)
			}
		}

		resOnlyQuestion := db.Create(&models.OnlyQuestion{
			Question:    question.Question,
			Ans:         question.Ans,
			Mistake1:    question.Mistake1,
			Mistake2:    question.Mistake2,
			Mistake3:    question.Mistake3,
			QimgPath:    question.QimgPath,
			AimgPath:    question.AimgPath,
			MimgPath1:   question.MimgPath1,
			MimgPath2:   question.MimgPath2,
			MimgPath3:   question.MimgPath3,
			QuestionNum: question.QuestionNum,
			Genre:       genre,
			Season:      season,
		}).Error

		if resOnlyQuestion != nil {
			fmt.Println(resOnlyQuestion)
		}
		fmt.Println(question)

	}

}
