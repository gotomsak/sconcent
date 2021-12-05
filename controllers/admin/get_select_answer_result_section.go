package admin

import (
	"context"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminGetSelectAnswerResultSection(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	if b, _ := sess.Values["admin_user_id"]; b == nil {
		return c.String(http.StatusUnauthorized, "401")
	}

	sqi := c.Param("select_question_id")
	res := new(models.AdminGetSelectAnswerResultSectionRes)

	db := utils.SqlConnect()
	defer db.Close()

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	rows, err := db.Where("select_question_id ?", sqi).Find(&models.AnswerResultSection{}).Rows()
	dbColl := mc.Database("gotoSys").Collection("gotoConc")
	for rows.Next() {
		var ars models.AnswerResultSection

		var gsars models.GetSelectAnswerResultSection
		db.ScanRows(rows, &ars)
		filter, err := primitive.ObjectIDFromHex(ars.ConcID)

		err = dbColl.FindOne(context.Background(), bson.D{{"_id", filter}}).Decode(&gsars.Concentration)
		gsars.AnswerResultSection = ars
		res.SelectAnswerResultSection = append(res.SelectAnswerResultSection, gsars)
		if err != nil {
			return c.String(http.StatusBadRequest, "bat")
		}
	}

	return c.JSON(200, res)
}
