// Package favicon provides a favicon downloader
package favicon

import (
	"context"
	"fmt"
	"image"
	"net/http"
	// Imports required to decode favicon images
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "lucor.dev/paw/internal/ico"
)

const (
	defaultMinSize = 32
)

type Service func(host string) string

type Options struct {
	// Client is http.Client used to download the favicon. Leave nil to use the
	// http.Default with a timeout of 10 seconds
	Client *http.Client
	// MinSize is the min size of the favicon to be considered valid.
	// favicon smaller than minSize will be ignored unless ForceMinSize is true
	MinSize int
	// ForceMinSize when true will force to return a favicon even if its size is
	// smaller than MinSize
	ForceMinSize bool
	Service      Service
}

// Download tries to download the favicon with highest resolution for the specified host
// By default it looks into the standard locations
// - http://<host>/apple-touch-icon.png
// - http://<host>/favicon.ico
// - http://<host>/favicon.png
// Alternatively a third-party service can be used via the Service option.
// Example:
// // Use the DuckDuckGo service
// ddg := func(host string) string  {
// 	return fmt.Sprintf("https://icons.duckduckgo.com/ip3/%s.ico", host)
// }
// img, err := favicon.Download(e.ctx, host, favicon.Options{
// 	ForceMinSize: true,
// 	Service: ddg,
// })
func Download(ctx context.Context, host string, opts Options) (image.Image, error) {
	return nil, fmt.Errorf("could not found any favicon at default locations")
}
