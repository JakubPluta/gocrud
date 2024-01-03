package service

import (
	"context"

	"github.com/JakubPluta/gocrud/data/request"
	"github.com/JakubPluta/gocrud/data/response"
)

type PostService interface {
	Create(ctx context.Context, request request.PostCreateRequest)
	Update(ctx context.Context, request request.PostUpdateRequest)
	Delete(ctx context.Context, postId string)
	GetAll(ctx context.Context) []response.PostResponse
	GetById(ctx context.Context, postId string) response.PostResponse
}
