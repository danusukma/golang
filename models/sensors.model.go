package models

import (
	"fmt"
	"golangsensors/db"
	"net/http"

	"time"

	"github.com/go-playground/validator"
)

type Sensors struct {
	Id          int       `json:"id"`
	Sensorvalue int       `json:"sensor_value"`
	Id1         int       `json:"id1"`
	Id2         string    `json:"id2"`
	Timestamp   time.Time `json:"timestamp"`
}

type SensorsStore struct {
	Sensorvalue int    `json:"sensor_value"  validate:"required"`
	Id1         int    `json:"id1"  validate:"required"`
	Id2         string `json:"id2"  validate:"required"`
}

func FetchAllSensors() (Response, error) {
	var obj Sensors
	var arrobj []Sensors
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM sensors"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Sensorvalue, &obj.Id1, &obj.Id2, &obj.Timestamp)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	fmt.Println(res)

	return res, nil
}

func StoreSensors(sensorvalue int, id1 int, id2 string) (Response, error) {
	var res Response

	v := validator.New()

	var_sen := SensorsStore{
		Sensorvalue: sensorvalue,
		Id1:         id1,
		Id2:         id2,
	}

	err := v.Struct(var_sen)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT sensors (sensorvalue, id1, id2) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(sensorvalue, id1, id2)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateSensors(id int, sensorvalue int, id1 int, id2 string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE sensors SET sensorvalue = ?, id1 = ?, id2 = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(sensorvalue, id1, id2, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteSensors(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM sensors WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
