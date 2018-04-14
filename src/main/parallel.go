package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", welcomePage)
	e.POST("/upload", DownloadFiles)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the matrix!")
}

func handleError(err error, context echo.Context) error {
	if err != nil {
		return context.String(http.StatusNotFound, err.Error())
	}
	return nil
}
