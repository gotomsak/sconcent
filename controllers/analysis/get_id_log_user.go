package analysis

import (
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetIDLogUser(c echo.Context) error {
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

	res := new(models.AdminGetIDLogUserRes)
	db := utils.SqlConnect()
	defer db.Close()

	db.Where("user_id = ?", sess.Values["user_id"]).Find(&res.GetIDLogUser)
	fmt.Println(res.GetIDLogUser)

	return c.JSON(http.StatusOK, res)
}
