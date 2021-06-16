package admin

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"golang.org/x/crypto/bcrypt"
)

func AdminSignup(c echo.Context) error {
	db := utils.SqlConnect()
	defer db.Close()

	us := new(models.AdminUserSignup)
	u := models.AdminUser{}
	if err := c.Bind(us); err != nil {
		return c.JSON(http.StatusInternalServerError, "Faild Bind")
	}
	fmt.Println(u)
	check := os.Getenv("TOKEN")
	if check != us.Token {
		return c.JSON(500, "access token mistaken")
	}
	u.PasswordDigest = AdminPasswordHash(us.Password)
	u.Username = us.Username
	u.Email = us.Email

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err := tx.Create(&u).Error
		tx.Rollback()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, "testok")
	}
	err := db.Create(&u).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func AdminSignin(c echo.Context) error {
	db := utils.SqlConnect()
	defer db.Close()
	u := models.AdminUser{}
	us := new(models.AdminUserSignin)
	if err := c.Bind(us); err != nil {
		return c.JSON(http.StatusInternalServerError, "Faild Bind")
	}
	check := os.Getenv("TOKEN")
	if check != us.Token {
		return c.JSON(500, "access token mistaken")
	}

	db.Where("email = ?", us.Email).Find(&u)

	passcheck := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(us.Password))

	if passcheck == nil {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		sess.Values["authenticated"] = true
		sess.Values["admin_user_id"] = u.ID
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, models.UserSend{UserID: u.ID, Username: u.Username})
	}
	return passcheck
}

func AdminSignout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func AdminCheckSession(c echo.Context) error {
	sess, _ := session.Get("session", c)
	log.Print(sess.Values["authenticated"])
	if b, _ := sess.Values["admin_user_id"]; b == nil {
		return c.String(http.StatusUnauthorized, "401")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	return c.JSON(http.StatusOK, "200")

}

func AdminPasswordHash(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
