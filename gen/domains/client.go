// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// domains client
//
// Command:
// $ goa gen ninepod/design

package domains

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "domains" service client.
type Client struct {
	DomainsEndpoint goa.Endpoint
}

// NewClient initializes a "domains" service client given the endpoints.
func NewClient(domains goa.Endpoint) *Client {
	return &Client{
		DomainsEndpoint: domains,
	}
}

// Domains calls the "domains" endpoint of the "domains" service.
func (c *Client) Domains(ctx context.Context, p *DomainsPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.DomainsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
