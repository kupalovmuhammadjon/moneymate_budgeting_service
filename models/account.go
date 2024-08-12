package models

type CreateAccount struct {
	UserID    string  `json:"user_id" bson:"user_id"`
	Name      string  `json:"name" bson:"name"`
	Type      string  `json:"type" bson:"type"`
	Balance   float64 `json:"balance" bson:"balance"`
	Currency  string  `json:"currency" bson:"currency"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
}

type Account struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	UserID    string  `json:"user_id" bson:"user_id"`
	Name      string  `json:"name" bson:"name"`
	Type      string  `json:"type" bson:"type"`
	Balance   float64 `json:"balance" bson:"balance"`
	Currency  string  `json:"currency" bson:"currency"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
}

type Accounts struct {
	Accounts []*Account `json:"accounts" bson:"accounts"`
}

type AccountFilter struct {
	Page     int32   `json:"page" bson:"page"`
	Limit    int32   `json:"limit" bson:"limit"`
	UserID   string  `json:"user_id" bson:"user_id"`
	Name     string  `json:"name" bson:"name"`
	Type     string  `json:"type" bson:"type"`
	Balance  float64 `json:"balance" bson:"balance"`
	Currency string  `json:"currency" bson:"currency"`
}
