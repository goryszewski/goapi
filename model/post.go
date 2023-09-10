package model

type Post struct {
	ID int 
	title string
}

var (
	posts []*Post
	nextIDP = 1 
)

func GetPosts () []*Post {
	return posts
}