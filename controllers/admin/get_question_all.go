package admin

import (
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminGetQuestionAll(c echo.Context) error {

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
	db := utils.SqlConnect()
	defer db.Close()
	res := new(models.AdminGetQuestionAllRes)
	rows, err := db.Model(&models.Question{}).Rows()
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var question models.Question
		db.ScanRows(rows, &question)
		res.QuestionAll = append(res.QuestionAll, question)
	}

	return c.JSON(200, res)

}
