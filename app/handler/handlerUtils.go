package handler

import (
	"encoding/json"
	"io"
	"net/url"
)

type InvalidURL struct {
	message string
}

func (i InvalidURL) Error() string {
	return i.message
}

type GenericError struct {
	message string
}

func (g GenericError) Error() string {
	return g.message
}

func validURL(testUrl string) bool {
	_, err := url.ParseRequestURI(testUrl)
	if err != nil {
		return false
	}
	return true
}

func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
