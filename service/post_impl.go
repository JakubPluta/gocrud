package service

import (
	"context"

	"github.com/JakubPluta/gocrud/data/request"
	"github.com/JakubPluta/gocrud/data/response"
	"github.com/JakubPluta/gocrud/helpers"
	"github.com/JakubPluta/gocrud/model"
	"github.com/JakubPluta/gocrud/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

// Create implements PostService.
func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	postData := model.Post{
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	post, err := p.PostRepository.GetById(ctx, postId)
	helpers.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)

}

// GetAll implements PostService.
func (p *PostServiceImpl) GetAll(ctx context.Context) []response.PostResponse {
	posts := p.PostRepository.GetAll(ctx)
	var postResponses []response.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, response.PostResponse{
			Id:          post.Id,
			Title:       post.Title,
			Published:   post.Published,
			Description: post.Description,
		})
	}
	return postResponses
}

// GetById implements PostService.
func (p *PostServiceImpl) GetById(ctx context.Context, postId string) response.PostResponse {
	post, err := p.PostRepository.GetById(ctx, postId)
	helpers.ErrorPanic(err)
	postResponse := response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}
	return postResponse
}

// Update implements PostService.
func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	postData := model.Post{
		Id:          request.Id,
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Update(ctx, postData)
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository}
}
