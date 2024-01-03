package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/JakubPluta/gocrud/helpers"
	"github.com/JakubPluta/gocrud/model"
	"github.com/JakubPluta/gocrud/prisma/db"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewPostRepository(db *db.PrismaClient) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		Db: db,
	}
}

func (p *PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	result, err := p.Db.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helpers.ErrorPanic(err)
	fmt.Println("Row inserted: ", result)
}

func (p *PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	result, err := p.Db.Post.FindMany(db.Post.ID.Equals(post.Id)).Update(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helpers.ErrorPanic(err)
	fmt.Println("Row updated: ", result)
}

func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helpers.ErrorPanic(err)
	fmt.Println("Rows deleted: ", result)
}

func (p *PostRepositoryImpl) GetById(ctx context.Context, postId string) (model.Post, error) {
	post, err := p.Db.Post.FindFirst(db.Post.ID.Equals(postId)).Exec(ctx)
	helpers.ErrorPanic(err)
	published, _ := post.Published()
	description, _ := post.Description()
	postData := model.Post{
		Id:          post.ID,
		Title:       post.Title,
		Published:   published,
		Description: description,
	}
	if post != nil {
		return postData, nil
	}
	return postData, errors.New("post not found")
}

func (p *PostRepositoryImpl) GetAll(ctx context.Context) []model.Post {
	allPosts, err := p.Db.Post.FindMany().Exec(ctx)
	helpers.ErrorPanic(err)

	var posts []model.Post
	for _, post := range allPosts {
		published, _ := post.Published()
		description, _ := post.Description()
		postData := model.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: description,
		}
		posts = append(posts, postData)
	}
	return posts
}
