package learning

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func InitMaxFrequency(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	userID := c.FormValue("user_id")
	maxBlinkNumber := c.FormValue("max_blink_number")
	maxFaceMoveNumber := c.FormValue("max_face_move_number")

	maxBlinkNumberFloat, _ := strconv.ParseFloat(maxBlinkNumber, 64)
	maxFaceMoveNumberFloat, _ := strconv.ParseFloat(maxFaceMoveNumber, 64)
	var maxBlinkFrequency float64 = (maxBlinkNumberFloat / 60) * 5
	var maxFaceMoveFrequency float64 = (maxFaceMoveNumberFloat / 60) * 5

	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	db := utils.SqlConnect()
	defer db.Close()

	var frequency Frequency
	err = db.Where("user_id = ?", userID).First(&frequency).Error
	if err != nil {
		frequency := Frequency{
			UserID: utils.StringToUint(userID),

			MaxFaceMoveNumber:    maxFaceMoveNumberFloat,
			MaxFaceMoveFrequency: maxFaceMoveFrequency,
			MaxBlinkNumber:       maxBlinkNumberFloat,
			MaxBlinkFrequency:    maxBlinkFrequency,
		}
		err = db.Create(&frequency).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	err = db.Model(&frequency).Updates(Frequency{

		MaxFaceMoveNumber:    maxFaceMoveNumberFloat,
		MaxFaceMoveFrequency: maxFaceMoveFrequency,
		MaxBlinkNumber:       maxBlinkNumberFloat,
		MaxBlinkFrequency:    maxBlinkFrequency,
	}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(frequency)

	return c.JSON(http.StatusOK, "ok")
}
