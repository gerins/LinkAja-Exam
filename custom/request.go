package custom

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Middleware biasa2 ini biasanya jalan di paling awal
func MiddlewarePrintIP(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.RealIP())
		return next(c)
	}
}

// Middleware Body Dump adalah middleware yang jalan terakhir
// Biasanya untuk logging hasil response
func MiddlewareBodyDump() echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, req, resp []byte) {
			fmt.Println("Body Dump Middleware Running", string(resp))
		},
	})
}
