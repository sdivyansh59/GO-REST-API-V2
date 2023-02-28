package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrFetchingComment = errors.New("failed to fetch comment by id")
)

// Comment - a representation of the comment
// structure for our services
type Comment struct {
	ID     string
	Slug   string
	Author string
	Body   string
}

// Stroe - this interface define all of the methods 
//that our service  need in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment,error)
	PostComment(context.Context, Comment) (Comment,error)
	UpdateComment(context.Context, string, Comment) (Comment,error)
	DeleteComment(context.Context, string) (error)
}

// Service - is the struct on which all our
// logic will be build on top of
type Service struct {
	Store Store
}

// NewService - (constructor) return a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriving a comment")
	cmt, err := s.Store.GetComment(ctx, id)// from repository layer
	if err != nil {
		fmt.Println(err)
		return Comment{},err
	}

	return cmt,nil
}


func (s *Service) UpdateComment(ctx context.Context, id string ,cmt Comment) (Comment,error) {
	return Comment{},ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment ,error){
	insertedCmt, err := s.Store.PostComment(ctx,cmt)
    if err != nil {
		return Comment{}, err
	}	
	return insertedCmt,nil
}


/* 
Why we are passing ctx context.Context every where ?
1. ctx helps us to pass information from 1 layer to other 
ex:
--> Put into ctx
ctx = context.WithValue(ctx,"request_id", "unique-string")

-->Get from ctx
ctx.Value("request_id") // return unique-string


2. Passing Time-out using ctx in all layers of applications


*/