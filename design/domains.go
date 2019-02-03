package design

import . "goa.design/goa/dsl"

// Service describes a service
var _ = Service("domains", func() {
	Description("Get pods, categories, etc.")
	// Method describes a service method (endpoint)
	Method("domains", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "number")
			Attribute("b", String, "category")
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)

		// Error which applies to all methods.
		// Error(ErrUnauthorized, Unauthorized)

		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/domains/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})
