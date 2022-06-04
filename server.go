package main

import (
	"golangsensors/db"

	"golangsensors/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
