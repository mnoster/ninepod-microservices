// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// domains HTTP server types
//
// Command:
// $ goa gen ninepod/design

package server

import (
	domains "ninepod/gen/domains"
)

// NewDomainsPayload builds a domains service domains endpoint payload.
func NewDomainsPayload(a int, b string) *domains.DomainsPayload {
	return &domains.DomainsPayload{
		A: a,
		B: b,
	}
}
