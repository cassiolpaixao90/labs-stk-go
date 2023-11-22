package helpers

import (
	"context"
	"github.com/avast/retry-go"
	"net/http"
)

type HttpClientRequest struct {
	http.Client
}

func (h *HttpClientRequest) Retry(ctx context.Context, req *http.Request) (error, *http.Response) {
	var response *http.Response
	err := retry.Do(func() error {
		res, err := h.Client.Do(req)
		response = res
		if err != nil {
			return err
		}
		return nil
	},
		retry.Attempts(1),
		retry.Delay(1),
		retry.OnRetry(func(n uint, err error) {
			// TODO add logger
		}),
	)
	return err, response
}
