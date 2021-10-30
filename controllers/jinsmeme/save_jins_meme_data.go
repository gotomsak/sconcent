package jinsmeme

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SaveJinsMemeData(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	bind := models.SaveJinsMemeDataBind{}

	if err := c.Bind(&bind); err != nil {
		return c.JSON(500, "concentratioon not found")
	}

	req := models.SaveJinsMemeDataReq{}
	req.StartTime = bind.StartTime
	req.EndTime = bind.EndTime

	queryValue := url.Values{}
	rv := reflect.ValueOf(req)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		value := rv.FieldByName(field.Name)
		queryValue.Add(field.Tag.Get("json"), value.String())

	}

	reqJ, err := http.NewRequest("GET", "https://apis.jins.com/meme/v1/users/me/official/computed_data", strings.NewReader(queryValue.Encode()))
	if err != nil {
		return err
	}
	reqJ.Header.Set("Accept", "application/json")
	reqJ.Header.Set("Authorization", "Bearer "+req.AccessToken)

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

	var resRoot models.SaveJinsMemeDataRes

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &resRoot)

	save := models.SaveJinsMemeDataSave{}
	save.ConcDataID = bind.ConcDataID
	save.SaveJinsMemeDataRes = resRoot

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)
	dbColl := mc.Database("gotoSys").Collection("JinsMemeData")

	res, err := dbColl.InsertOne(context.Background(), req)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
