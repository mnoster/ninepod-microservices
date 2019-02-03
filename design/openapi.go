package design

import . "goa.design/goa/dsl"

var _ = Service("openapi", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/openapi")
	})
	Files("/swagger.json", "../../gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})
