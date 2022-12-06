package main

import (
	"net/http"

	"github.com/labstack/echo"
	"fmt"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, ping())
	})

	e.Logger.Fatal(e.Start(":1323"))
	fmt.Printf("Hello World\n")
}

func ping() string {
	return "pong"
}
