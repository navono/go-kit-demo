package main

import (
	"context"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

type ctxManager struct {
	name string
}

func main() {
	svc := stringService{}
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			ctx = context.WithValue(ctx, ctxManager{"request"}, r)
			return ctx
		}),
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
