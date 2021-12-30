package handler

import (
	"github.com/2001adarsh/url-shortner/app/storage"
	"log"
	"net/http"
)

type ShortURL struct {
	log *log.Logger
	db  storage.Database
}

func (s *ShortURL) ViewHandler(writer http.ResponseWriter, request *http.Request) {
	s.log.Println("ViewHandler")
}

func (s *ShortURL) ValidateMiddleware(handler http.Handler) http.Handler {
	return handler
}

func NewShortURL(log *log.Logger, db storage.Database) *ShortURL {
	return &ShortURL{log: log, db: db}
}
