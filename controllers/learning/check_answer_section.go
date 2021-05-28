package learning

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAnswerSection(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Please sign in")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	cas := new(models.CheckAnswerSectionBind)
	if err = c.Bind(cas); err != nil {
		return c.String(http.StatusInternalServerError, "The format is different")
	}
	fmt.Println(cas)

	db := utils.SqlConnect()
	defer db.Close()

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	results := mc.Database("sconcent").Collection("answer_result_sectoin_ids")
	res, err := results.InsertOne(context.Background(), models.Results{ResultIDs: cas.AnswerResultIDs})
	var resID string
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		resID = oid.Hex()
	} else {
		return c.JSON(http.StatusInternalServerError, "Not objectid.ObjectID, do what you want")
	}

	answerResultSection := models.AnswerResultSection{
		UserID:              cas.UserID,
		AnswerResultIDs:     resID,
		CorrectAnswerNumber: cas.CorrectAnswerNumber,
		StartTime:           utils.StringToTime(cas.StartTime),
		EndTime:             utils.StringToTime(cas.EndTime),
	}

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err := tx.Create(&answerResultSection)
		tx.Rollback()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, "testok")
	}
	err = db.Create(&answerResultSection).Error
	if err != nil {
		return err
	}

	answerResultSectionIDSend := models.AnswerResultSectionIDSend{}
	answerResultSectionIDSend.AnswerResultSectionID = answerResultSection.ID

	return c.JSON(http.StatusOK, answerResultSectionIDSend)
}
