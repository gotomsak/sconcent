package learning

import (
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetSelectQuestion(c echo.Context) error {

	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	if b, _ := sess.Values["user_id"]; b == nil {
		return c.String(http.StatusUnauthorized, "401")
	}

	db := utils.SqlConnect()
	defer db.Close()

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	res := new(models.GetSelectQuestionRes)
	rows, err := db.Model(models.SelectQuestion{}).Rows()

	for rows.Next() {
		var sq models.SelectQuestion
		db.ScanRows(rows, &sq)
		res.SelectQuestion = append(res.SelectQuestion, sq)

	}

	if err != nil {
		fmt.Println(err)
		return c.JSON(200, nil)
	}
	fmt.Println(res)

	// fmt.Println(res.Concentration)

	return c.JSON(200, res)

}
