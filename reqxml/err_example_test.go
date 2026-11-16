package reqxml_test

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/gophertools/requests/v2"
	"github.com/gophertools/requests/v2/reqxml"
)

func ExampleError() {
	type ErrorXML struct {
		XMLName   xml.Name `xml:"Error"`
		Code      string   `xml:"Code"`
		Message   string   `xml:"Message"`
		RequestID string   `xml:"RequestId"`
		HostID    string   `xml:"HostId"`
	}

	var errObj ErrorXML
	err := requests.
		URL("http://example.s3.us-east-1.amazonaws.com").
		AddValidator(reqxml.Error(&errObj)).
		Fetch(context.Background())
	switch {
	case errors.Is(err, requests.ErrInvalidHandled):
		fmt.Println(errObj.Message)
	case err != nil:
		fmt.Println("Error!", err)
	case err == nil:
		fmt.Println("unexpected success")
	}

	// Output:
	// Access Denied
}
