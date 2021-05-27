package server

import (
	"fmt"
	"net/http"

	apihandlers "github.com/taalhach/safepay/internal/server/api_handlers"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

const port = 4000

var serveApi = cobra.Command{
	Use:   "serve_api",
	Short: "serves penny pincher api",
	Long:  fmt.Sprintf("serves penny pincher api on localhost %v port", port),
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.HTTPErrorHandler = customErrHandler

		// add request logger middleware
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, in=${bytes_in}, out=${bytes_out}\n",
		}))

		// in case panic occurs middle for recovering from panics
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize:         1 << 10,
			DisableStackAll:   true,
			DisablePrintStack: true,
		}))

		// CORS middleware
		e.Use(middleware.CORS())

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "awesome")
		})
		e.GET("/exchange-routing", apihandlers.ExchangeRouting)

		listenAddress := fmt.Sprintf(":%d", port)
		e.Logger.Fatal(e.Start(listenAddress))
	},
}

func init() {
	rootCommand.AddCommand(&serveApi)
}
