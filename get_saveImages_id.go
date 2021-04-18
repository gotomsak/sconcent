package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gotomsak/sconcent/utils"
	"github.com/labstack/echo/v4"
)

func GetSaveImagesID(c echo.Context) error {
	var token string = ""
	if token = c.Request().Header["Access-Token"][0]; token == "" {
		return c.JSON(500, "access token not found")
	}
	check := os.Getenv("TOKEN")
	if check != token {
		return c.JSON(500, "access token mistaken")
	}

	gs := c.QueryParam("type")
	dir := "./images/" + gs
	idFile := "./images/" + gs + "/id"
	if err := os.MkdirAll(dir, 0777); err != nil {
		fmt.Println(err)
	}

	lastID, err := ioutil.ReadFile(idFile)
	var newID uint64
	if err != nil {
		os.Create(idFile)
		bnewID := utils.Uint2bytes(0, 5)
		ioutil.WriteFile(idFile, bnewID, os.ModePerm)
	} else {
		newID = utils.Bytes2uint(lastID)
		newID++
		bnewID := utils.Uint2bytes(newID, 8)
		ioutil.WriteFile(idFile, bnewID, os.ModePerm)
	}

	return c.JSON(200, &GetSaveImagesIDRes{ID: newID})
}
