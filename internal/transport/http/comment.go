package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sdivyansh59/GO-REST-API-V2/internal/comment"
)


type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment,error)
	GetComment(context.Context, string) (comment.Comment,error)
	UpdateComment(context.Context,  string, comment.Comment) (comment.Comment,error)
	DeleteComment(context.Context, string) (error)
}


func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return 
	}

	cmt, err := h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}


func  (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

}


func  (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

}


func  (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	
}