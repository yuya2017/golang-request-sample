package post

type Post struct {
	PostCount int
	Signatures []string
}

func New(postCount int, signatures []string) *Post {
	return &Post{postCount, signatures}
}
