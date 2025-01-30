package dto

type CreatePostResponse struct {
	ID      int32  `json:"id"`
	UserID  int32  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
