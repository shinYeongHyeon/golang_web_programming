package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shinYeongHyeon/golang_web_programming/membership"
)

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(membershipLogger()))

	membership.CreateSubRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func membershipLogger() func(c echo.Context, reqBody []byte, resBody []byte) {
	return func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println("Request " + c.Request().Method + " " + c.Request().RequestURI)
		printIfBodyExist(reqBody)

		fmt.Printf("Response %d\n", c.Response().Status)
		printIfBodyExist(resBody)
	}
}

func printIfBodyExist(body []byte) {
	if bodyString := string(body[:]); bodyString != "" {
		fmt.Println(bodyString + "\n")
	}
}
