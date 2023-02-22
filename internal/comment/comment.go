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
	Body   string
	Author string
}

// Stroe - this interface define all of the methods 
//that our service  need in order to operate
type Store interface {
	GetComment (context.Context, string) (Comment,error)
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


func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment ,error){
	return Comment{}, ErrNotImplemented
}