package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/echo", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		intervalS := c.QueryParam("interval")
		if intervalS == "" {
			intervalS = "1"
		}
		interval, err := strconv.ParseFloat(intervalS, 32)
		if err != nil {
			interval = 1
		}
		enc := json.NewEncoder(c.Response())
		for {
			if err := enc.Encode(fmt.Sprintf("%s / %s", time.Now(), os.Getenv("KUBE_NODE_NAME"))); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(time.Duration(interval) * time.Second)
		}
	})

	e.GET("/hello", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())
		if err := enc.Encode(fmt.Sprintf("%s / %s", time.Now(), os.Getenv("KUBE_NODE_NAME"))); err != nil {
			return err
		}
		c.Response().Flush()
		return nil
	})
	e.GET("/headers", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())
		for name, headers := range c.Request().Header {
			for _, h := range headers {
				if err := enc.Encode(fmt.Sprintf("%s: %s", name, h)); err != nil {
					return err
				}
			}
		}
		c.Response().Flush()
		return nil
	})

	e.Logger.Fatal(e.Start(":8090"))
}
