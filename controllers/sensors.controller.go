package controllers

import (
	"golangsensors/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func FetchAllSensors(c echo.Context) error {
	result, err := models.FetchAllSensors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSensors(c echo.Context) error {
	sensor_value := c.FormValue("sensor_value")
	id1 := c.FormValue("id1")
	id2 := c.FormValue("id2")
	timestamp := c.FormValue("timestamp")

	conv_sensor_value, err := strconv.Atoi(sensor_value)
	conv_id1, err := strconv.Atoi(id1)
	conv_timestamp, err := strconv.Atoi(timestamp)

	result, err := models.StoreSensors(conv_sensor_value, conv_id1, id2, conv_timestamp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSensors(c echo.Context) error {
	id := c.FormValue("id")
	sensor_value := c.FormValue("sensor_value")
	id1 := c.FormValue("id1")
	id2 := c.FormValue("id2")
	timestamp := c.FormValue("timestamp")

	conv_id, err := strconv.Atoi(id)
	conv_sensor_value, err := strconv.Atoi(sensor_value)
	conv_id1, err := strconv.Atoi(id1)
	conv_timestamp, err := strconv.Atoi(timestamp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSensors(conv_id, conv_sensor_value, conv_id1, id2, conv_timestamp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSensors(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteSensors(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
