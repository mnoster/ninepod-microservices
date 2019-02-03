package ninepod

import (
	"context"
	"log"
	openapi "ninepod/gen/openapi"
)

// openapi service example implementation.
// The example methods log the requests and return zero values.
type openapisrvc struct {
	logger *log.Logger
}

// NewOpenapi returns the openapi service implementation.
func NewOpenapi(logger *log.Logger) openapi.Service {
	return &openapisrvc{logger}
}

// Domains implements domains.
func (s *openapisrvc) OpenAPI(ctx context.Context) (res int, err error) {
	s.logger.Print("------ openapi.openapi ---------")
	return
}
