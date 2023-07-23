package pdf

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
var wait *time.Duration

var Cmd = &cobra.Command{
	Use:   "pdf [url] [filename]",
	Args:  cobra.ExactArgs(2),
	Short: "Opens a website and create a PDF file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		options := Options{
			Url: args[0],
			Viewport: &Viewport{
				Width:  *width,
				Height: *height,
			},
			Wait: *wait,
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

		if err := chromedp.Run(ctx, makePdf(options, &res)); err != nil {
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
	wait = Cmd.Flags().Duration("wait", 0, "wait")
}
