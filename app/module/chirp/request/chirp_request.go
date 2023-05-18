package request

type PostRequest struct {
	Content  string `json:"content" validate:"required,max=255"`
	AuthorId string `json:"authorId" validate:"required"`
}
