package dtos

import "labs-stk-go/internal/application/domain"

type GetAllBooksResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (g *GetAllBooksResponse) ToDomain() *domain.BookDomain {
	return &domain.BookDomain{
		ID:   g.ID,
		Name: g.Name,
	}
}
