package main

import (
	"net/http"
	cli "ninepod/gen/http/cli/pods"
	"os"
	"time"

	"goa.design/goa"
	goahttp "goa.design/goa/http"
)

func doHTTP(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	var (
		doer goahttp.Doer
	)
	{
		doer = &http.Client{Timeout: time.Duration(timeout) * time.Second}
		if debug {
			doer = goahttp.NewDebugDoer(doer)
			doer.(goahttp.DebugDoer).Fprint(os.Stderr)
		}
	}

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
	)
}
func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}
