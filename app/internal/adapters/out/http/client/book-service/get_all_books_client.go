package book_service

import (
	"context"
	"encoding/json"
	"io"
	"labs-stk-go/internal/adapters/out/http/client/book-service/dtos"
	"labs-stk-go/internal/adapters/out/http/helpers"
	"labs-stk-go/internal/application/domain"
	"labs-stk-go/internal/application/ports/out"
	"net/http"
)

type GetAllBooksClient struct {
	client helpers.HttpClientRequest
}

func NewGetAllBooksClient() out.GetAllBooksClientPort {
	return &GetAllBooksClient{}
}

func (g *GetAllBooksClient) GetAllBooks(ctx context.Context, id string) (*domain.BookDomain, error) {

	req, _ := http.NewRequest(http.MethodGet, "url", nil)

	err, resp := g.client.Retry(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO add logger
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	var response *dtos.GetAllBooksResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.ToDomain(), nil
}
