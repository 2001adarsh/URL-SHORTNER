package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *ShortURL) MapHandler(writer http.ResponseWriter, request *http.Request) {
	//pathsToUrls := map[string]string{
	//	"urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	//	"yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	//}
	vars := mux.Vars(request)
	url := vars["url"]
	var err error
	//actualURL, ok := pathsToUrls[url]
	actualURL, notOk := s.db.Get(url)
	if notOk != nil {
		err = InvalidURL{message: "Invalid URL. URL not found in the database."}
	}

	switch err.(type) {
	case nil:
	case InvalidURL:
		s.log.Println("[ERROR] ", err, actualURL)
		writer.WriteHeader(http.StatusBadRequest)
		_ = ToJSON(&InvalidURL{message: err.Error()}, writer)
		return
	default:
		s.log.Println("[ERROR] Can't be redirected to the actual URL", actualURL)
		writer.WriteHeader(http.StatusInternalServerError)
		_ = ToJSON(&GenericError{message: "Something went wrong, please check and try again."}, writer)
		return
	}

	s.log.Println("[INFO] Redirecting to the actual URL", actualURL)
	http.Redirect(writer, request, actualURL, http.StatusMovedPermanently)
}
