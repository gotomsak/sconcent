package admin

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminGetIDLogUser(c echo.Context) error {
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

	userID, _ := strconv.Atoi(c.Param("user_id"))
	fmt.Println(reflect.TypeOf(c.Param("user_id")))

	res := new(models.AdminGetIDLogUserRes)
	db := utils.SqlConnect()
	defer db.Close()
	// getIDLog := new(models.GetIDLog)
	db.Where("user_id = ?", userID).Find(&res.GetIDLogUser)
	fmt.Println(res.GetIDLogUser)

	return c.JSON(http.StatusOK, res)
}
