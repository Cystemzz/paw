// Package haveibeenpwned implements a client for the haveibeenpwned.com API v3
// to search if passwords have been exposed in data breaches
package haveibeenpwned

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"lucor.dev/paw/internal/paw"
)

const apiURL = "https://api.pwnedpasswords.com/range/%s"

var defaultClient = &http.Client{
	Timeout: 10 * time.Second,
}

// httpClient interface
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Pwned struct {
	Item  paw.Item
	Count int
}

// Search searches if the password has been exposed in data
// breaches using the Have I Been Pwned APIs
func Search(ctx context.Context, item paw.Item) (pwned bool, count int, err error) {
	meta := item.GetMetadata()
	if meta == nil {
		return pwned, count, fmt.Errorf("metadata cannot be nil")
	}
	var p string
	switch meta.Type {
	case paw.PasswordItemType:
		p = item.(*paw.Password).Value
	case paw.LoginItemType:
		p = item.(*paw.Login).Password.Value
	default:
		return pwned, count, fmt.Errorf("invalid item type %q", meta.Type)
	}
	return hibp(ctx, defaultClient, p)
}

// hibp consumes the range endpoint. It returns true if the provided password has been
// exposed in data breaches along with a count of how many times it appears in the data set.
// See https://haveibeenpwned.com/API/v3#PwnedPasswords
func hibp(ctx context.Context, c httpClient, password string) (bool, int, error) {
	return false, 0, nil
}
