package admin

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

func AdminSaveSelectQuestion(c echo.Context) error {

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
	sqb := new(models.AdminSaveSelectQuestionBind)

	if err := c.Bind(sqb); err != nil {
		return c.JSON(500, "initData not found")
	}

	db := utils.SqlConnect()
	defer db.Close()

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	dbColl := mc.Database("learning").Collection("select_question_ids")

	sqID := primitive.NewObjectID()
	_, err = dbColl.InsertOne(context.Background(), models.SelectQuestionIDs{ID: sqID, SelectQuestionIDs: sqb.SelectQuestionIDs})

	sq := new(models.SelectQuestion)
	sq.SelectQuestionIDs = sqID.Hex()
	sq.SelectQuestionName = sqb.SelectQuestionName

	err = db.Create(&sq).Error

	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "data is missing")
	}
	fmt.Println(sq)

	return c.JSON(200, sq)

}
