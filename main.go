package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/gotomsak/sconcent/controllers/admin"
	"github.com/gotomsak/sconcent/controllers/analysis"
	"github.com/gotomsak/sconcent/controllers/ear"
	"github.com/gotomsak/sconcent/controllers/environment"
	"github.com/gotomsak/sconcent/controllers/frequency"
	"github.com/gotomsak/sconcent/controllers/jinsmeme"
	"github.com/gotomsak/sconcent/controllers/learning"
	"github.com/gotomsak/sconcent/controllers/recording"
	"github.com/gotomsak/sconcent/controllers/user"
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

	e.POST("/get_id", recording.GetID)
	e.POST("/init_ear", ear.InitEAR)
	e.POST("/init_max", frequency.InitMaxFrequency)
	e.POST("/init_min", frequency.InitMinFrequency)
	e.GET("/get_frequency", frequency.GetFrequency)
	e.GET("/get_ear", ear.GetEar)
	e.GET("/check_session", user.CheckSession)
	e.POST("/signin", user.Signin)
	e.POST("/signup", user.Signup)
	e.GET("/signout", user.Signout)
	e.POST("/save_concent", recording.SaveConcentration)
	e.POST("/save_face_point", recording.SaveFacePoint)
	e.GET("/get_rec_all", analysis.GetRecAll)
	e.POST("/save_concent_split", recording.SaveConcentSplit)
	e.POST("/save_environment", environment.SaveEnvironment)
	e.GET("/get_environment", environment.GetEnvironment)

	e.POST("/admin_signin", admin.AdminSignin)
	e.POST("/admin_signup", admin.AdminSignup)
	e.GET("/admin_signout", admin.AdminSignout)
	e.GET("/admin_check_session", admin.AdminCheckSession)
	e.GET("/admin_get_id_log_all", admin.AdminGetIDLogAll)
	e.GET("/admin_get_user_all", admin.AdminGetUserAll)
	e.GET("/admin_get_rec_all/:user_id", admin.AdminGetRecAll)
	e.GET("/admin_get_id_log_user/:user_id", admin.AdminGetIDLogUser)
	e.GET("/admin_get_rec_user_date/:conc_id", admin.AdminGetRecUserDate)
	e.GET("/admin_get_face_point/:face_point_id", admin.AdminGetFacePoint)

	e.POST("/question_ids", learning.GetQuestionIds)
	e.GET("/question", learning.GetQuestion)
	e.POST("/check_answer", learning.CheckAnswer)
	e.POST("/check_answer_section", learning.CheckAnswerSection)
	e.POST("/save_questionnaire", learning.SaveQuestionnaire)

	e.POST("/get_jins_meme_token", jinsmeme.GetJinsMemeToken)

	return e
}

func main() {
	utils.EnvLoad()
	e := router()

	e.Logger.Fatal(e.Start(":1323"))
	// e.Logger.Fatal(e.StartTLS(":1323", "./fullchain.pem", "./privkey.pem"))
}
