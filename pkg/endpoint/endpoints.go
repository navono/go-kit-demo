package endpoint

import (
	"context"
	"errors"
	"strings"

	lorem "navono/go-kit-demo/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

// ErrRequestTypeNotFound represents a invalid request type error
var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// LoremRequest represents http request
type LoremRequest struct {
	RequestType string
	Min         int
	Max         int
}

// LoremResponse represents http response
type LoremResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

// Endpoints wrapper
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// MakeLoremEndpoint creating Lorem Ipsum endpoint
func MakeLoremEndpoint(svc lorem.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		var (
			txt      string
			min, max int
		)

		min = req.Min
		max = req.Max

		if strings.EqualFold(req.RequestType, "Word") {
			txt = svc.Word(min, max)
		} else if strings.EqualFold(req.RequestType, "Sentence") {
			txt = svc.Sentence(min, max)
		} else if strings.EqualFold(req.RequestType, "Paragraph") {
			txt = svc.Paragraph(min, max)
		} else {
			return nil, ErrRequestTypeNotFound
		}

		return LoremResponse{Message: txt}, nil
	}
}
