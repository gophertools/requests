package reqxml

import "github.com/gophertools/requests/v2"

// Error is a [ValidatorHandler] that applies DefaultValidator
// and decodes the response as an XML object
// if the DefaultValidator check fails.
func Error(v any) requests.ResponseHandler {
	return requests.ValidatorHandler(requests.DefaultValidator, To(v))
}
