package recording

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveConcentration(c echo.Context) error {
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

	oldData := new(models.SaveConcentrationBind)
	newData := new(models.SaveConcentrationBind)
	if err := c.Bind(newData); err != nil {
		return c.JSON(500, "concentratioon not found")
	}
	fmt.Println(newData)

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	filter := bson.M{"_id": newData.ID}

	dbColl := mc.Database(newData.Type).Collection(newData.Measurement)

	err = dbColl.FindOne(context.Background(), filter).Decode(&oldData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "find error")
	}

	// oldData.Concentration = append(oldData.Concentration, newData.Concentration)
	// oldData.Concentration = newData.Concentration
	// bson.D{{"$set", bson.D{{"concentration", newData.Concentration}}}}
	res, err := dbColl.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"concentration": newData.Concentration, "memo": newData.Memo}})
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "update error")
	}
	fmt.Println(res)
	return c.JSON(200, "ok")
}
