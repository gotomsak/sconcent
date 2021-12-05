package normalization

import (
	"context"
	"fmt"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CorrectRefacter() {

	db := utils.SqlConnect()
	defer db.Close()

	rows, err := db.Model(&models.AnswerResultSection{}).Rows()

	if err != nil {
		fmt.Println("not model")
	}

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	dbColl := mc.Database("learning").Collection("answer_result_ids")
	for rows.Next() {
		var ars models.AnswerResultSection
		var ari models.AnswerResultIDs
		var ar models.AnswerResult
		var cn uint = 0
		db.ScanRows(rows, &ars)
		// fmt.Println(ars)
		filter, err := primitive.ObjectIDFromHex(ars.AnswerResultIDs)
		if err != nil {
			fmt.Println("not Filter")
		}
		err = dbColl.FindOne(context.Background(), bson.D{{"_id", filter}}).Decode(&ari)
		fmt.Println(ari)
		for _, v := range ari.AnswerResultIDs {
			db.Where("id = ?", v).First(&ar)
			fmt.Println(ar)
			if ar.AnswerResult == "correct" {
				cn += 1
			}
		}
		ars.CorrectAnswerNumber = cn
		db.Save(&ars)
	}
}
