package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toumakido/godoc/echoxml/def"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.POST("/", xmlRes)

	e.Logger.Fatal(e.Start(":8000"))
}

func xmlRes(c echo.Context) error {
	erresp := def.Response{
		Errinf: &def.ErrorInfo{
			Errcd:  "11",
			Errmsg: "XML Error",
		},
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXML)
	return c.XML(http.StatusBadRequest, erresp)
}
