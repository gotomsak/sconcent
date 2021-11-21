package learning

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetQuestionIds(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	sq := c.Param("select_question_id")

	db := utils.SqlConnect()
	defer db.Close()
	questions := models.Question{}

	if sq != "none" {
		sqdb := new(models.SelectQuestion)
		err = db.First(&sqdb, "id = ?", sq).Error
		if err != nil {
			fmt.Println("sqdata not found")
		}
		res := new(models.SelectQuestionIDs)
		mc, ctx := utils.MongoConnect()
		defer mc.Disconnect(ctx)
		filter, err := primitive.ObjectIDFromHex(sq)
		if err != nil {
			fmt.Println("can't convert")
		}

		dbColl := mc.Database("learning").Collection("select_question_ids")
		fmt.Println(sq)
		fmt.Println(res)
		fmt.Println(filter)
		err = dbColl.FindOne(context.Background(), bson.D{{"_id", filter}}).Decode(&res)
		gqi := models.GetQuestionIDs{
			QuestionIDs: res.SelectQuestionIDs,
		}
		return c.JSON(http.StatusOK, gqi)
	}

	db.Last(&questions)
	var count int
	db.Table("questions").Count(&count)
	firstID := int(questions.ID) - int(count) + 1
	firstIDU := uint(firstID)
	sIdsStr := c.FormValue("solved_ids")
	qIdsStr := c.FormValue("question_ids")
	newQuestionList := []uint{}
	solveList := utils.StrToUIntList(sIdsStr)
	questionList := utils.StrToUIntList(qIdsStr)
	solveList = append(solveList, questionList...)

	for {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(count)
		urandom := uint(random)
		if true == utils.SearchIDs(solveList, urandom) {
			continue
		}
		if true == utils.SearchIDs(newQuestionList, urandom) {
			continue
		}
		newQuestionList = append(newQuestionList, firstIDU+urandom)
		if len(newQuestionList) == 10 {
			break
		}
	}
	gqi := models.GetQuestionIDs{
		QuestionIDs: newQuestionList,
		SolvedIDs:   solveList,
	}
	return c.JSON(http.StatusOK, gqi)
}
