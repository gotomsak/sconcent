package jinsmeme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetJinsMemeToken(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	bind := models.GetJinsMemeTokenBind{}

	if err := c.Bind(&bind); err != nil {
		return c.JSON(500, "concentratioon not found")
	}
	req := models.GetJinsMemeTokenReq{}
	req.Code = bind.Code
	req.ClientSecret = os.Getenv("JINS_MEME_SECRET")
	req.ClientID = os.Getenv("JINS_MEME_ID")
	req.GrantType = "authorization_code"
	req.RedirectUri = "https://fland.kait-matsulab.com/callback"

	queryValue := url.Values{}
	rv := reflect.ValueOf(req)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		value := rv.FieldByName(field.Name)
		queryValue.Add(field.Tag.Get("json"), value.String())

	}

	reqJ, err := http.NewRequest("POST", "https://apis.jins.com/meme/v1/oauth/token", strings.NewReader(queryValue.Encode()))
	if err != nil {
		return err
	}
	reqJ.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(reqJ)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	accessTokenSave := models.GetJinsMemeTokenSave{}
	accessTokenSave.UserID = bind.UserID

	db := utils.SqlConnect()
	defer db.Close()

	var resRoot models.GetJinsMemeTokenRes

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &resRoot)
	fmt.Println(body)
	fmt.Println(resRoot.Scope)

	accessTokenSave.AccessToken = resRoot.AccessToken

	old := models.GetJinsMemeTokenSave{}

	err = db.Where("user_id = ?", accessTokenSave.UserID).First(&old).Error
	if err != nil {
		err = db.Create(&accessTokenSave).Error
	} else {
		err = db.Model(&accessTokenSave).Where("user_id = ?", accessTokenSave.UserID).Update("access_token", accessTokenSave.AccessToken).Error
	}

	if err != nil {
		return err
	}

	return c.JSON(200, resRoot)
}
