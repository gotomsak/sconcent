package environment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveEnvironment(c echo.Context) error {

	sess, err := session.Get("session", c)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	req := new(models.Environment)

	if err := c.Bind(req); err != nil {
		return c.JSON(500, "initData not found")
	}
	req.ID = primitive.NewObjectID()
	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)

	dbColl := mc.Database("gotoSys").Collection("environment")
	res, err := dbColl.InsertOne(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "insert error")
	}
	return c.JSON(http.StatusOK, res)
}
