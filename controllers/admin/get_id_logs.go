package admin

import (
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminGetIDLogs(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	if b, _ := sess.Values["user_id"]; b == nil {
		return c.String(http.StatusUnauthorized, "401")
	}

	db := utils.SqlConnect()
	defer db.Close()
	res := models.AdminGetIDLogsRes{}
	db.Find(&res.GetIDLogs)

	// mc, ctx := utils.MongoConnect()
	// defer mc.Disconnect(ctx)
	// res := new(models.AdminGetIDLogsRes)

	// dbColl := mc.Database("gotoSys").Collection("gotoConc")
	// cur, err := dbColl.Find(context.Background(), bson.D{})
	// if err == mongo.ErrNoDocuments {
	// 	return c.String(http.StatusNotFound, "not found")
	// }

	// for cur.Next(ctx) {
	// 	curN := models.GetIDSave{}
	// 	err := cur.Decode(&curN)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	res.GetIDLogs = append(res.GetIDLogs, curN)
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.JSON(500, "find error")
	// }
	return c.JSON(http.StatusOK, res)
}
