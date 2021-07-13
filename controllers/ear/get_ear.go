package ear

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEar(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "session error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	fmt.Println(sess.Values["user_id"])

	if b, _ := sess.Values["user_id"]; b == nil {
		return c.String(http.StatusUnauthorized, "401")
	}

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	result := new(models.GetEarDataRes)

	dbColl := mc.Database("gotoSys").Collection("ear")

	cur, err := dbColl.Find(context.Background(), bson.D{{Key: "userid", Value: sess.Values["user_id"]}})
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not max frequency")
	}

	for cur.Next(ctx) {
		resear := models.EarData{}
		err := cur.Decode(&resear)
		if err != nil {
			log.Fatal(err)
		}
		result.EarData = append(result.EarData, resear)
	}

	return c.JSON(http.StatusOK, result)

}
