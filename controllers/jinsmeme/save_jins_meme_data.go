package jinsmeme

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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
	token := models.GetJinsMemeTokenSave{}

	db := utils.SqlConnect()
	defer db.Close()
	// fmt.Println(bind.UserID)
	db.First(&token, "user_id = ?", bind.UserID)
	// fmt.Println(token.AccessToken)
	req := models.SaveJinsMemeDataReq{}
	start := url.QueryEscape(bind.StartTime.Format("2006-01-02T15:04:05") + "+09:00")
	end := url.QueryEscape(bind.EndTime.Format("2006-01-02T15:04:05") + "+09:00")
	req.StartTime = start
	req.EndTime = end
	req.AccessToken = token.AccessToken
	fmt.Println(req.StartTime)
	fmt.Println(req.EndTime)

	// queryValue := url.Values{}
	// rv := reflect.ValueOf(req)
	// rt := rv.Type()

	// for i := 0; i < rt.NumField(); i++ {
	// 	field := rt.Field(i)

	// 	value := rv.FieldByName(field.Name)
	// 	queryValue.Add(field.Tag.Get("json"), value.String())

	// }
	u, err := url.Parse("https://apis.jins.com/meme/v1/users/me/official/computed_data?date_from=" + req.StartTime + "&date_to=" + req.EndTime)
	fmt.Println(u.String())
	reqJ, err := http.NewRequest("GET", u.String(), nil)
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

	var resRoot models.SaveJinsMemeDataRes

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	json.Unmarshal(body, &resRoot)
	fmt.Println(body)
	fmt.Println(resRoot)

	save := models.SaveJinsMemeDataSave{}
	save.ConcDataID = bind.ConcDataID
	save.SaveJinsMemeDataRes = resRoot

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)
	dbColl := mc.Database("gotoSys").Collection("JinsMemeData")

	res, err := dbColl.InsertOne(context.Background(), save)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
