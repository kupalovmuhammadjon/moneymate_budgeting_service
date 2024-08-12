package models

type CreateCategory struct {
	UserID string `json:"user_id" bson:"user_id"`
	Name   string `json:"name" bson:"name"`
	Type   string `json:"type" bson:"type"`
}

type Category struct {
	ID        string `json:"id" bson:"_id,omitempty"` 
	UserID    string `json:"user_id" bson:"user_id"`
	Name      string `json:"name" bson:"name"`
	Type      string `json:"type" bson:"type"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

type CategoryFilter struct {
	Page   int32  `json:"page" bson:"page"`
	Limit  int32  `json:"limit" bson:"limit"`
	UserID string `json:"user_id" bson:"user_id"`
	Name   string `json:"name" bson:"name"`
	Type   string `json:"type" bson:"type"`
}

type Categories struct {
	Categories []*Category `json:"categories" bson:"categories"`
}
