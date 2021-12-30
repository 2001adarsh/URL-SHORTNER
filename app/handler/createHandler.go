package handler

import "net/http"

func (s *ShortURL) CreateHandler(writer http.ResponseWriter, request *http.Request) {
	s.log.Println("123")
}
