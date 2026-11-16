package requests_test

import (
	"context"
	"fmt"

	"github.com/gophertools/requests/v2"
)

func ExampleHasStatusErr() {
	err := requests.
		URL("http://example.com/404").
		CheckStatus(200).
		Fetch(context.Background())
	if requests.HasStatusErr(err, 404) {
		fmt.Println("got a 404")
	}
	// Output:
	// got a 404
}
