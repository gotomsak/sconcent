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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AdminGetRecUserDate(c echo.Context) error {

	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	if b, _ := sess.Values["admin_user_id"]; b == nil {
		return c.String(http.StatusUnauthorized, "401")
	}
	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	res := new(models.GetRecUserDateRes)

	concID := c.Param("conc_id")

	filter, err := primitive.ObjectIDFromHex(concID)
	fmt.Println(filter)
	if err != nil {
		return c.String(http.StatusBadRequest, "noconc")
	}

	dbColl := mc.Database("gotoSys").Collection("gotoConc")

	err = dbColl.FindOne(context.Background(), bson.D{{"_id", filter}}).Decode(&res.Concentration)

	if err != nil {
		return c.String(http.StatusBadRequest, "bat")
	}

	userIDFilter := bson.D{{Key: "userid", Value: res.Concentration.UserID}}

	dbColl = mc.Database("gotoSys").Collection("maxFrequency")

	cur, err := dbColl.Find(context.Background(), userIDFilter)

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

	dbColl = mc.Database("gotoSys").Collection("minFrequency")

	cur, err = dbColl.Find(context.Background(), userIDFilter)
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

	dbColl = mc.Database("gotoSys").Collection("facePoint")

	faceAllPointFilter := res.Concentration.Concentration.FacePointAll
	fmt.Println(faceAllPointFilter)
	err = dbColl.FindOne(context.Background(), bson.D{{"_id", faceAllPointFilter}}).Decode(&res.FacePointAll)

	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	fmt.Println(res.Concentration)

	return c.JSON(200, res)

}
