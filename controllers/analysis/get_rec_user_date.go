package analysis

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

func GetRecUserDate(c echo.Context) error {

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

	res := new(models.GetRecUserDateRes)

	concID := c.Param("conc_id")

	filter, err := primitive.ObjectIDFromHex(concID)
	fmt.Println(filter)
	if err != nil {
		return c.String(http.StatusBadRequest, "noconc")
	}

	dbColl := mc.Database("gotoSys").Collection("gotoConc")

	err = dbColl.FindOne(context.Background(), bson.D{{"_id", filter}}).Decode(&res.GetConcentrationRes)

	if err != nil {
		return c.String(http.StatusBadRequest, "bat")
	}

	userIDFilter := bson.D{{Key: "userid", Value: res.GetConcentrationRes.UserID}}

	dbColl = mc.Database("gotoSys").Collection("environment")

	cur, err := dbColl.Find(context.Background(), userIDFilter)
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "not max frequency")
	}

	for cur.Next(ctx) {
		resenv := models.EnvironmentRes{}
		enviro := models.Environment{}
		maxF := models.MaxFrequency{}
		minF := models.MinFrequency{}
		ear := models.EarData{}
		err := cur.Decode(&enviro)

		dbColl = mc.Database("gotoSys").Collection("maxFrequency")
		err = dbColl.FindOne(context.TODO(), bson.M{"_id": enviro.MaxFreqID}).Decode(&maxF)
		resenv.MaxFreq = maxF
		dbColl = mc.Database("gotoSys").Collection("minFrequency")
		err = dbColl.FindOne(context.TODO(), bson.M{"_id": enviro.MinFreqID}).Decode(&minF)
		resenv.MinFreq = minF
		dbColl = mc.Database("gotoSys").Collection("ear")
		err = dbColl.FindOne(context.TODO(), bson.M{"_id": enviro.EarID}).Decode(&ear)
		resenv.Ear = ear
		resenv.Date = enviro.Date
		resenv.Name = enviro.Name
		resenv.ID = enviro.ID

		if err != nil {
			log.Fatal(err)
		}

		res.GetEnvironmentRes = append(res.GetEnvironmentRes, resenv)
	}

	dbColl = mc.Database("gotoSys").Collection("facePoint")

	faceAllPointFilter := res.GetConcentrationRes.Concentration.FacePointAll
	fmt.Println(faceAllPointFilter)
	err = dbColl.FindOne(context.Background(), bson.D{{"_id", faceAllPointFilter}}).Decode(&res.FacePointAll)

	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	// fmt.Println(res.Concentration)

	return c.JSON(200, res)

}
