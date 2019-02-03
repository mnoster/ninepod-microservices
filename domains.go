package ninepod

import (
	"context"
	"log"
	domains "ninepod/gen/domains"
)

// domains service example implementation.
// The example methods log the requests and return zero values.
type domainssrvc struct {
	logger *log.Logger
}

// NewDomains returns the domains service implementation.
func NewDomains(logger *log.Logger) domains.Service {
	return &domainssrvc{logger}
}

// Domains implements domains.
func (s *domainssrvc) Domains(ctx context.Context, p *domains.DomainsPayload) (res int, err error) {
	s.logger.Print("------ domains.domains ---------")
	s.logger.Print("\n--- A --- \n", p.A)
	s.logger.Print("\n--- B --- \n", p.B)
	return
}
