package admin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotomsak/sconcent/models"
	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminGetFacePoint(c echo.Context) error {
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

	res := new(models.GetFacePointRes)
	facePointID := c.Param("face_point_id")

	filter, err := primitive.ObjectIDFromHex(facePointID)
	fmt.Println(filter)
	if err != nil {
		return c.String(http.StatusBadRequest, "nofacePointID")
	}
	dbColl := mc.Database("gotoSys").Collection("facePoint")
	err = dbColl.FindOne(context.Background(), bson.D{{"_id", filter}}).Decode(&res.FacePointAll)
	if err != nil {
		return c.String(http.StatusNotFound, "noFacePointData")
	}
	return c.JSON(200, res)
}
