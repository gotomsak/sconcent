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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetID(c echo.Context) error {
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

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)
	getID := new(models.GetIDBind)
	if err := c.Bind(getID); err != nil {
		return c.JSON(500, "concentration not found")
	}
	fmt.Println(getID)

	dbColl := mc.Database(getID.Type).Collection(getID.Measurement)
	newID := primitive.NewObjectID()
	facePointNewID := primitive.NewObjectID()
	request := models.GetIDSave{
		ID:            newID,
		UserID:        getID.UserID,
		Type:          getID.Type,
		Work:          getID.Work,
		Memo:          getID.Memo,
		Measurement:   getID.Measurement,
		Concentration: getID.Concentration,
		// FacePointAll:  facePointNewID,
	}
	_, err = dbColl.InsertOne(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "insert error")
	}

	dbColl = mc.Database("gotoSys").Collection("facePoint")

	facePointRequest := models.GetFacePointIDSave{
		ID: facePointNewID,
		// FacePointAll: []interface{}{},
	}

	_, err = dbColl.InsertOne(context.Background(), facePointRequest)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "insert error")
	}

	db := utils.SqlConnect()
	defer db.Close()

	getIDLog := models.GetIDLog{
		ConcDataID: newID.Hex(),
		UserID:     getID.UserID,
	}
	err = db.Create(&getIDLog).Error

	return c.JSON(200, &models.GetIDRes{ConcDataID: newID, FacePointID: facePointNewID})
}
