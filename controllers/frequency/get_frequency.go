package frequency

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

func GetFrequency(c echo.Context) error {
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

	result := new(models.GetFrequencyResData)

	dbColl := mc.Database("gotoSys").Collection("maxFrequency")

	cur, err := dbColl.Find(context.Background(), bson.D{{Key: "userid", Value: sess.Values["user_id"]}})
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not max frequency")
	}

	for cur.Next(ctx) {
		resmax := models.MaxFrequency{}
		err := cur.Decode(&resmax)
		if err != nil {
			log.Fatal(err)
		}
		result.MaxFrequency = append(result.MaxFrequency, resmax)
	}

	dbColl = mc.Database("gotoSys").Collection("minFrequency")

	cur, err = dbColl.Find(context.Background(), bson.D{{Key: "userid", Value: sess.Values["user_id"]}})
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not min frequency")
	}
	for cur.Next(ctx) {
		resmin := models.MinFrequency{}
		err := cur.Decode(&resmin)
		if err != nil {
			log.Fatal(err)
		}
		result.MinFrequency = append(result.MinFrequency, resmin)
	}

	return c.JSON(http.StatusOK, result)

}
