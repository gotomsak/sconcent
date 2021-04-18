package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetID(c echo.Context) error {
	var token string = ""
	if token = c.Request().Header["Access-Token"][0]; token == "" {
		return c.JSON(500, "access token not found")
	}
	check := os.Getenv("TOKEN")
	if check != token {
		return c.JSON(500, "access token mistaken")
	}

	mc, ctx := utils.MongoConnect()
	defer mc.Disconnect(ctx)
	getID := new(GetIDBind)
	if err := c.Bind(getID); err != nil {
		return c.JSON(500, "concentration not found")
	}

	dbColl := mc.Database(getID.Type).Collection(getID.Measurement)
	newID := primitive.NewObjectID()
	request := GetIDSave{
		ID:            newID,
		Type:          getID.Type,
		Measurement:   getID.Measurement,
		Concentration: getID.Concentration,
	}
	res, err := dbColl.InsertOne(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, "insert error")
	}
	fmt.Println(res.InsertedID)

	return c.JSON(200, &GetIDRes{ID: newID})
}
