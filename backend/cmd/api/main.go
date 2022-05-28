package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

    app.GET("/", func(ctx echo.Context) error {
        return ctx.String(http.StatusOK, "Hello, World!!!")
    })

    err := app.Start(":8080")
    app.Logger.Fatal(err)
}
