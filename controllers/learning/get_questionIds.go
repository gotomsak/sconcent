package learning

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetQuestionIds(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	db := utils.SqlConnect()
	defer db.Close()
	questions := models.Question{}
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
