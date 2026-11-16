package requests_test

import (
	"net/http"

	"github.com/gophertools/requests/v2/reqtest"
)

func init() {
	http.DefaultClient.Transport = reqtest.Replay("testdata")
}
