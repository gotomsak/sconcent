package jinsmeme

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/gotomsak/sconcent/models"
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
	req := models.GetJinsMemeTokenReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(500, "concentratioon not found")
	}

	req.ClientSecret = os.Getenv("JINS_MEME_SECRET")
	req.ClientID = os.Getenv("JINS_MEME_ID")
	req.GrantType = "authorization_code"
	req.RedirectUri = "https://fland.kait-matsulab.com/callback"

	queryValue := url.Values{}
	rv := reflect.ValueOf(req)
	rt := rv.Type()
	//error pointerだから？
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		// kind:= field.Type.Kind()
		value := rv.FieldByName(field.Name)
		queryValue.Add(field.Name, value.String())
		println(value.String())
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
	fmt.Println(resp)

	return err
}
