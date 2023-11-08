package model

type Category struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
