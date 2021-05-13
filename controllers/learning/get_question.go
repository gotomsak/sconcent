package learning

import (
	"net/http"
	"regexp"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetQuestion(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	db := utils.SqlConnect()
	defer db.Close()
	question := models.Question{}
	questionSend := models.QuestionSend{}
	questionID := c.QueryParam("id")
	uquestionID := utils.StringToUint(questionID)
	db.First(&question, uquestionID)
	qimg := regexp.MustCompile(",").Split(question.QimgPath, -1)
	ans := []string{question.Ans, question.Mistake1, question.Mistake2, question.Mistake3}
	aimg := []string{question.AimgPath, question.MimgPath1, question.MimgPath2, question.MimgPath3}
	utils.Shuffle(ans)
	utils.Shuffle(aimg)
	questionSend.QuestionID = question.ID
	questionSend.AnsList = ans
	questionSend.AimgList = aimg
	questionSend.QimgPath = qimg
	questionSend.QuestionNum = question.QuestionNum
	questionSend.Question = question.Question
	questionSend.Season = question.Season
	questionSend.Genre = question.Genre
	return c.JSON(http.StatusOK, questionSend)
}
