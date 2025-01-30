package dto

type CreatePostRequest struct {
	UserID  int32  `json:"userID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
