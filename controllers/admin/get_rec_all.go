package admin

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

func AdminGetRecAll(c echo.Context) error {
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

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	res := new(models.GetRecAllRes)

	filter := bson.D{{Key: "userid", Value: sess.Values["user_id"]}}

	dbColl := mc.Database("gotoSys").Collection("gotoConc")

	cur, err := dbColl.Find(context.Background(), filter)

	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not min frequency")
	}

	for cur.Next(ctx) {
		curN := models.GetConcentrationRes{}
		err := cur.Decode(&curN)
		if err != nil {
			log.Fatal(err)
		}
		res.Concentration = append(res.Concentration, curN)
	}

	dbColl = mc.Database("gotoSys").Collection("maxFrequency")

	cur, err = dbColl.Find(context.Background(), filter)
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not max frequency")
	}

	for cur.Next(ctx) {
		curN := models.MaxFrequency{}
		err := cur.Decode(&curN)
		if err != nil {
			log.Fatal(err)
		}
		res.MaxFrequency = append(res.MaxFrequency, curN)
	}
	dbColl = mc.Database("gotoSys").Collection("maxFrequency")

	cur, err = dbColl.Find(context.Background(), filter)
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not min frequency")
	}

	for cur.Next(ctx) {
		curN := models.MinFrequency{}
		err := cur.Decode(&curN)
		if err != nil {
			log.Fatal(err)
		}
		res.MinFrequency = append(res.MinFrequency, curN)
	}

	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	return c.JSON(http.StatusOK, res)

}
