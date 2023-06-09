package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Message string `json:"message"`
}

func greet(c echo.Context) error {
	message := fmt.Sprintf("Hello, World!, Current Time:%s", time.Now().Format("2006/01/02 15:04:05"))
	response := Response{Message: message}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", greet)

	svr := &http.Server{Addr: ":8080", Handler: e}
	svr.SetKeepAlivesEnabled(false) // Keep Aliveの無効

	log.Fatal(svr.ListenAndServe())
}
