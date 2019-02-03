package pods

import (
	"context"
	"log"
	podssvc "ninepod/gen/pods"
)

// pods service example implementation.
// The example methods log the requests and return zero values.
type podssrvc struct {
	logger *log.Logger
}

// NewPods returns the pods service implementation.
func NewPods(logger *log.Logger) podssvc.Service {
	return &podssrvc{logger}
}

// Pods implements pods.
func (s *podssrvc) Pods(ctx context.Context, p *podssvc.PodsPayload) (res int, err error) {
	s.logger.Print("pods.pods")
	return
}
