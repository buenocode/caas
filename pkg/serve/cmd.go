package serve

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/buenocode/caas/pkg/pdf"
	"github.com/buenocode/caas/pkg/screenshot"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var listen *string

func getInt(values url.Values, name string, defaultValue int) int {
	if !values.Has(name) {
		return defaultValue
	}
	value, _ := strconv.Atoi(values.Get(name))
	return value
}

func getBool(values url.Values, name string, defaultValue bool) bool {
	if !values.Has(name) {
		return defaultValue
	}
	if defaultValue {
		return values.Get(name) != "false"
	} else {
		return values.Get(name) == "false"
	}
}

func getDuration(values url.Values, name string, defaultValue time.Duration) time.Duration {
	if !values.Has(name) {
		return defaultValue
	}
	time, err := time.ParseDuration(values.Get(name))
	if err != nil {
		log.Printf("getDuration for %s error %v\n", name, err)
		return defaultValue
	} else {
		return time
	}
}

var Cmd = &cobra.Command{
	Use:   "serve",
	Args:  cobra.NoArgs,
	Short: "Listen for http requests and make screenshots on demand...",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := gin.Default()
		r.GET("/screenshot", func(c *gin.Context) {
			params := c.Request.URL.Query()
			options := screenshot.Options{
				Url: params.Get("url"),
				Viewport: &screenshot.Viewport{
					Width:  getInt(params, "width", 1920),
					Height: getInt(params, "height", 1080),
				},
				DarkMode:   getBool(params, "darkmode", false),
				Insecure:   getBool(params, "insecure", false),
				JavaScript: getBool(params, "javascript", true),
				Wait:       getDuration(params, "wait", 0),
			}

			log.Printf("options %v\n", options)
			log.Printf("viewport %v\n", options.Viewport)

			// create context
			ctx, cancel := chromedp.NewContext(
				context.Background(),
				// chromedp.WithDebugf(log.Printf),
			)
			defer cancel()

			var res []byte
			if err := chromedp.Run(ctx, screenshot.MakeScreenshot(options, &res)); err != nil {
				log.Printf("Error: %v\n", err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.Data(http.StatusOK, "image/png", res)
		})
		r.POST("/screenshot", func(c *gin.Context) {
			var options screenshot.Options
			if err := c.BindJSON(&options); err != nil {
				log.Printf("Error: %v\n", err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			log.Printf("options %v\n", options)
			log.Printf("viewport %v\n", options.Viewport)

			// create context
			ctx, cancel := chromedp.NewContext(
				context.Background(),
				// chromedp.WithDebugf(log.Printf),
			)
			defer cancel()

			var res []byte
			if err := chromedp.Run(ctx, screenshot.MakeScreenshot(options, &res)); err != nil {
				log.Printf("Error: %v\n", err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.Data(http.StatusOK, "image/png", res)
		})
		r.GET("/pdf", func(c *gin.Context) {
			params := c.Request.URL.Query()
			options := pdf.Options{
				Url: params.Get("url"),
				Viewport: &pdf.Viewport{
					Width:  getInt(params, "width", 1920),
					Height: getInt(params, "height", 1080),
				},
				// DarkMode:   getBool(params, "darkmode", false),
				Insecure:   getBool(params, "insecure", false),
				JavaScript: getBool(params, "javascript", true),
				Wait:       getDuration(params, "wait", 0),
			}

			log.Printf("options %v\n", options)
			log.Printf("viewport %v\n", options.Viewport)

			// create context
			ctx, cancel := chromedp.NewContext(
				context.Background(),
				// chromedp.WithDebugf(log.Printf),
			)
			defer cancel()

			var res []byte
			if err := chromedp.Run(ctx, pdf.MakePdf(options, &res)); err != nil {
				log.Printf("Error: %v\n", err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.Data(http.StatusOK, "application/pdf", res)
		})
		r.POST("/pdf", func(c *gin.Context) {
			var options pdf.Options
			if err := c.BindJSON(&options); err != nil {
				log.Printf("Error: %v\n", err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			log.Printf("options %v\n", options)
			log.Printf("viewport %v\n", options.Viewport)

			// create context
			ctx, cancel := chromedp.NewContext(
				context.Background(),
				// chromedp.WithDebugf(log.Printf),
			)
			defer cancel()

			var res []byte
			if err := chromedp.Run(ctx, pdf.MakePdf(options, &res)); err != nil {
				log.Printf("Error: %v\n", err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.Data(http.StatusOK, "application/pdf", res)
		})
		return r.Run(*listen)
	},
}

func init() {
	listen = Cmd.Flags().String("listen", "localhost:8080", "")
}
