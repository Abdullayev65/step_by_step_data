package ioPut

type Sign struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ArticleInput struct {
	Data string `json:"data"`
}

type CommentInput struct {
	Comment string `json:"comment"`
}
