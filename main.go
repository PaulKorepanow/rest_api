package main

import (
	"github.com/PaulKorepanow/rest_api/server"
	"github.com/labstack/echo/v4"
)

func main() {
	//fmt.Println("hallo world!")
	e := echo.New()

	e.GET("/data", server.HandleUsers())
	e.POST("/data", server.HandlePost())
	e.Logger.Fatal(e.Start(":8080"))
}
