package repository

import (
	"context"

	"github.com/JakubPluta/gocrud/model"
)

type PostRepository interface {
	Save(ctx context.Context, post model.Post)
	Update(ctx context.Context, post model.Post)
	Delete(ctx context.Context, postId string)
	GetById(ctx context.Context, postId string) (model.Post, error) // returns model.Post
	GetAll(ctx context.Context) []model.Post
}
