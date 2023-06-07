package main

import (
	"weavestore/api"

	"github.com/labstack/echo/v4"
)

func main() {
	// initialize the KV store with options
	// init http handlers
	e := echo.New()
	apiHandler := api.NewAPIInstance()

	// Store HTTP API
	e.POST("/insert/", apiHandler.InsertObject)
	e.POST("/read/", apiHandler.Read)
	e.POST("/delete/", apiHandler.DeleteObject)
	e.POST("/update/", apiHandler.UpdateObject)
	e.POST("/updateBulk/", apiHandler.UpdateBulk)
	e.Logger.Fatal(e.Start(":8000"))
}
