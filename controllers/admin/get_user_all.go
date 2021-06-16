package admin

import (
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminGetUserAll(c echo.Context) error {
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

	res := new(models.AdminGetUserAllRes)
	users := new([]models.User)
	db.Find(&users)
	for _, v := range *users {
		res.UsersInfo = append(res.UsersInfo, models.UserInfo{UserID: v.ID, Username: v.Username, Email: v.Email})
	}

	return c.JSON(200, res)
}
