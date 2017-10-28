package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTPHandlerUtil implement interface
type HTTPHandlerUtil struct {
	adapter *PredictionAdapter
}

// HTTPHandlerFactory creates a interface for HTTPHandlerUtil
func HTTPHandlerFactory() *HTTPHandlerUtil {
	thisHandler := new(HTTPHandlerUtil)
	thisHandler.adapter = new(PredictionAdapter)
	return thisHandler
}

func (handler *HTTPHandlerUtil) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.Method == "GET" {
		if r.URL.Path == "/Demands" {
			handler.adapter.HandleDemandsRequest(w, 0, 0)
			return
		}
	}
	path := r.URL.Path
	if path == "/" {
		path = "/StartPage.html"
	}

	modifiedPath, contentType := handler.parseRequest(&path)
	data, err := ioutil.ReadFile(string(modifiedPath))
	log.Println(modifiedPath)
	if err == nil {
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
	}
}

func (handler *HTTPHandlerUtil) parseRequest(path *string) (string, string) {
	var contentType string
	modifiedSourcePath := "../../Webview" + *(path)
	if strings.HasSuffix(*path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(*path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(*path, ".js") {
		modifiedSourcePath = "../../ViewModel" + *(path)
		contentType = "application/javascript"
	} else if strings.HasSuffix(*path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	return modifiedSourcePath, contentType
}