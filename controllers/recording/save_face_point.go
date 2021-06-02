package recording

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveFacePoint(c echo.Context) error {
	start := time.Now()
	var token string = ""
	sess, err := session.Get("session", c)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	if token = c.Request().Header["Access-Token"][0]; token == "" {
		return c.JSON(500, "access token not found")
	}
	check := os.Getenv("TOKEN")
	if check != token {
		return c.JSON(500, "access token mistaken")
	}

	bindData := new(models.PostFacePointSave)

	oldData := new(models.PostFacePointSave)

	if err := c.Bind(bindData); err != nil {
		return c.JSON(500, "concentratioon not found")
	}

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	filter := bson.M{"_id": bindData.ID}

	dbColl := mc.Database("gotoSys").Collection("facePoint")

	err = dbColl.FindOne(context.Background(), filter).Decode(&oldData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	oldData.FacePointAll = append(oldData.FacePointAll, bindData.FacePointAll...)

	res, err := dbColl.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"face_point_all": oldData.FacePointAll}})
	fmt.Println(res)
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	return c.JSON(200, "ok")
}
