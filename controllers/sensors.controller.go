package controllers

import (
	"golangsensors/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func FetchAllSensors(c echo.Context) error {

	id1 := c.FormValue("id1")
	id2 := c.FormValue("id2")

	start_timestamp := c.FormValue("start_timestamp")
	end_timestamp := c.FormValue("end_timestamp")

	conv_id1, err := strconv.Atoi(id1)
	conv_start_timestamp, err := strconv.Atoi(start_timestamp)
	conv_end_timestamp, err := strconv.Atoi(end_timestamp)

	result, err := models.FetchAllSensors(conv_id1, id2, conv_start_timestamp, conv_end_timestamp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSensors(c echo.Context) error {
	sensorvalue := c.FormValue("sensorvalue")
	id1 := c.FormValue("id1")
	id2 := c.FormValue("id2")

	conv_sensorvalue, err := strconv.Atoi(sensorvalue)
	conv_id1, err := strconv.Atoi(id1)

	result, err := models.StoreSensors(conv_sensorvalue, conv_id1, id2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSensors(c echo.Context) error {
	id := c.FormValue("id")
	sensorvalue := c.FormValue("sensorvalue")
	id1 := c.FormValue("id1")
	id2 := c.FormValue("id2")

	conv_id, err := strconv.Atoi(id)
	conv_sensorvalue, err := strconv.Atoi(sensorvalue)
	conv_id1, err := strconv.Atoi(id1)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSensors(conv_id, conv_sensorvalue, conv_id1, id2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSensors(c echo.Context) error {
	idx := c.FormValue("idx")

	conv_id, err := strconv.Atoi(strings.TrimSpace(idx))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteSensors(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
