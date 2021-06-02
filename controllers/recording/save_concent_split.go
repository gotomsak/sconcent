package recording

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveConcentSplit(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	oldData := new(models.SaveConcentrationBind)
	newData := new(models.PostConcentSplitBind)

	if err := c.Bind(newData); err != nil {
		return c.JSON(500, "concentration not found")
	}

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)
	dbColl := mc.Database(newData.Type).Collection(newData.Measurement)

	filter := bson.M{"_id": newData.ID}

	err = dbColl.FindOne(context.Background(), filter).Decode(&oldData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	oldData.Concentration.C1 = append(oldData.Concentration.C1, newData.Concentration.C1...)
	oldData.Concentration.C2 = append(oldData.Concentration.C2, newData.Concentration.C2...)
	oldData.Concentration.C3 = append(oldData.Concentration.C3, newData.Concentration.C3...)
	oldData.Concentration.Date = append(oldData.Concentration.Date, newData.Concentration.Date...)
	oldData.Concentration.W = append(oldData.Concentration.W, newData.Concentration.W...)

	res, err := dbColl.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"concentration": oldData.Concentration}})
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "update error")
	}
	fmt.Println(res)
	return c.JSON(200, "ok")
}