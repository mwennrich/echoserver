package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {

	e := echo.New()
	e.GET("/stream", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		intervalS := c.QueryParam("interval")
		if intervalS == "" {
			intervalS = "1s"
		}
		interval, err := time.ParseDuration(intervalS)
		if err != nil {
			interval = time.Second
		}
		enc := json.NewEncoder(c.Response())
		for {
			if err := enc.Encode(fmt.Sprintf("%s / %s", time.Now(), os.Getenv("KUBE_NODE_NAME"))); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(interval)
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

	e.POST("/echo", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		data := c.FormValue("data")

		enc := json.NewEncoder(c.Response())
		if err := enc.Encode(fmt.Sprintf("%s / %s / %v", time.Now(), os.Getenv("KUBE_NODE_NAME"), data)); err != nil {
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

	e.GET("/speed", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		sizeS := c.QueryParam("size")
		if sizeS == "" {
			sizeS = "10Mi"
		}
		quantity, err := resource.ParseQuantity(sizeS)
		if err != nil {
			// return fmt.Errorf("failed to parse quantity: %v", err)
			c.Response().WriteHeader(http.StatusInternalServerError)
			c.Response().Write([]byte(fmt.Sprintf("failed to parse quantity: %v\n", err)))
			c.Response().Flush()
			return nil
		}

		data := make([]byte, quantity.Value())

		c.Response().Write(data)
		c.Response().Flush()
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		help := `Usage:
> curl http://example.com/headers
> curl http://example.com/hello
> curl http://example.com/stream
> curl http://example.com/stream?interval=0.5s
> curl http://example.com/speed?size=10Gi -o /dev/null
> echo -n "blafasel"| curl  --data-urlencode data@- http://example.com/echo
`

		c.Response().Write([]byte(help))
		c.Response().Flush()
		return nil

	})
	e.Logger.Fatal(e.Start(":8090"))
}
