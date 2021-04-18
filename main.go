package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gotomsak/sconcent/testcheck"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://localhost:3000", "https://192.168.1.10:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/get_id", GetID)
	// e.GET("/check_answer", le)

	e.POST("/save_concent", SaveConcentration)
	return e
}

func main() {
	utils.EnvLoad()
	e := router()

	e.Logger.Fatal(e.Start(":1323"))
	// e.Logger.Fatal(e.StartTLS(":1323", "./fullchain.pem", "./privkey.pem"))
}
