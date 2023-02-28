package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sdivyansh59/GO-REST-API-V2/internal/comment"

	"github.com/gorilla/mux"
)


type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment,error)
	GetComment(context.Context, string) (comment.Comment,error)
	UpdateComment(context.Context,  string, comment.Comment) (comment.Comment,error)
	DeleteComment(context.Context, string) (error)
}

type Response struct{
	Message string
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
	 vars := mux.Vars(r)
	 id := vars["id"]
	 if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return 
	 }

	 cmt, err := h.Service.GetComment(r.Context(),id)
	 if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	 }

	 if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	 }
}


func  (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return 
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	cmt, err := h.Service.UpdateComment(r.Context(),id,cmt);
	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}


func  (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	err := h.Service.DeleteComment(r.Context(),id)
	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	if err := json.NewEncoder(w).Encode(Response{Message : "Sucessfully deleted"}); err != nil {
		panic(err)
	}


}