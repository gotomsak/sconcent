package environment

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

func GetEnvironment(c echo.Context) error {
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

	result := new(models.GetEnvironmentRes)

	dbColl := mc.Database("gotoSys").Collection("environment")

	cur, err := dbColl.Find(context.Background(), bson.D{{Key: "userid", Value: sess.Values["user_id"]}})
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

		result.Environments = append(result.Environments, resenv)
	}

	return c.JSON(http.StatusOK, result)

}
