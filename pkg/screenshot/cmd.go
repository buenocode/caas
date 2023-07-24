package screenshot

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

var width *int
var height *int
var darkMode *bool
var insecure *bool
var javascript *bool
var wait *time.Duration

var Cmd = &cobra.Command{
	Use:   "screenshot [url] [filename]",
	Args:  cobra.ExactArgs(2),
	Short: "Opens a website, make a screenshot and save it into a file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		options := Options{
			Url: args[0],
			Viewport: &Viewport{
				Width:  *width,
				Height: *height,
			},
			DarkMode:   *darkMode,
			Insecure:   *insecure,
			JavaScript: *javascript,
			Wait:       *wait,
		}
		filename := args[1]

		log.Printf("options %v\n", options)
		log.Printf("viewport %v\n", options.Viewport)

		// create context
		ctx, cancel := chromedp.NewContext(
			context.Background(),
			// chromedp.WithDebugf(log.Printf),
		)
		defer cancel()

		var res []byte

		if err := chromedp.Run(ctx, MakeScreenshot(options, &res)); err != nil {
			log.Fatal(err)
			return err
		}
		if err := os.WriteFile(filename, res, 0o644); err != nil {
			log.Fatal(err)
			return err
		}

		log.Printf("wrote %s as %s\n", options.Url, filename)
		return nil
	},
}

func init() {
	width = Cmd.Flags().Int("width", 1920, "width")
	height = Cmd.Flags().Int("height", 1080, "height")
	darkMode = Cmd.Flags().Bool("darkmode", false, "darkmode")
	insecure = Cmd.Flags().Bool("insecure", false, "insecure")
	javascript = Cmd.Flags().Bool("javascript", true, "javascript")
	wait = Cmd.Flags().Duration("wait", 0, "wait")
}
