package main

import "net/http"

type ImageLoader struct{}

func (h *ImageLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// TODO: allows requesting proxy images from the 'tmp' folder

}
