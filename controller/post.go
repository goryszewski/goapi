package controller

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"web/api/model"
)

type postController struct {
	postIDPattern *regexp.Regexp
}

func (this postController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/posts" {
		this.getAll(w,r)
	}
}

func (this *postController ) getAll (w http.ResponseWriter , r *http.Request) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Return All posts")
	encodeResponseAsSJON(model.GetPosts(),w)
}

func newPostController() *postController{
	return &postController{ 
		postIDPattern: regexp.MustCompile(`^/posts/(\d+)/?`),
	}
}