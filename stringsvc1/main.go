package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type ctxManager struct {
	name string
}

// type Middleware func(endpoint.Endpoint) endpoint.Endpoint

// func loggingMiddleware(logger kitlog.Logger) Middleware {
// 	return func(next endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (interface{}, error) {
// 			logger.Log("msg", "calling endpoint")
// 			defer logger.Log("msg", "called endpoint")
// 			return next(ctx, request)
// 		}
// 	}
// }

type loggingMiddleware struct {
	logger kitlog.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(ctx, s)
	return
}

func (mw loggingMiddleware) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Count(ctx, s)
	return
}

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// svc := stringService{}

	// var uppercase endpoint.Endpoint
	// uppercase = makeUppercaseEndpoint(svc)
	// uppercase = loggingMiddleware(kitlog.With(logger, "method", "uppercase"))(uppercase)

	// var count endpoint.Endpoint
	// count = makeCountEndpoint(svc)
	// count = loggingMiddleware(kitlog.With(logger, "method", "count"))(count)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

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
