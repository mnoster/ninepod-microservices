package design

import . "goa.design/goa/dsl"

// API describes the global properties of the API server.
var _ = API("pods", func() {
	Title("API for working with pods (domains, categories, etc.)")
	Description("HTTP service for ninepod")
	Server("pods", func() {
		Description("Production host")
		Services("domains", "openapi")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000")
		})

	})
	License(func() {
		Name("MIT")
	})
})
