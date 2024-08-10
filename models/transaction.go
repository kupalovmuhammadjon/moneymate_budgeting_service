package models

type CreateTransaction struct {
	UserID      string  `json:"user_id" bson:"user_id"`
	AccountID   string  `json:"account_id" bson:"account_id"`
	CategoryID  string  `json:"category_id" bson:"category_id"`
	Amount      float64 `json:"amount" bson:"amount"`
	Type        string  `json:"type" bson:"type"`
	Description string  `json:"description" bson:"description"`
	Date        string  `json:"date" bson:"date"`
}

type Transaction struct {
	ID          string  `json:"id" bson:"_id,omitempty"` // Use `bson:"_id,omitempty"` for MongoDB ID field
	UserID      string  `json:"user_id" bson:"user_id"`
	AccountID   string  `json:"account_id" bson:"account_id"`
	CategoryID  string  `json:"category_id" bson:"category_id"`
	Amount      float64 `json:"amount" bson:"amount"`
	Type        string  `json:"type" bson:"type"`
	Description string  `json:"description" bson:"description"`
	Date        string  `json:"date" bson:"date"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
	UpdatedAt   string  `json:"updated_at" bson:"updated_at"`
}

type TransactionFilter struct {
	Page       int32   `json:"page" bson:"page"`
	Limit      int32   `json:"limit" bson:"limit"`
	UserID     string  `json:"user_id" bson:"user_id"`
	AccountID  string  `json:"account_id" bson:"account_id"`
	CategoryID string  `json:"category_id" bson:"category_id"`
	Amount     float64 `json:"amount" bson:"amount"`
	Type       string  `json:"type" bson:"type"`
	Date       string  `json:"date" bson:"date"`
}

type Transactions struct {
	Transactions []*Transaction `json:"transactions" bson:"transactions"`
}

type Spending struct {
	UserID           string   `json:"user_id" bson:"user_id"`
	TotalAmount      float64  `json:"total_amount" bson:"total_amount"`
	CategoryName     string   `json:"category_name" bson:"category_name"`
	TransactionTypes []string `json:"transaction_types" bson:"transaction_types"`
}

type Spendings struct {
	Spendings []*Spending `json:"spendings" bson:"spendings"`
}

type Income struct {
	UserID           string   `json:"user_id" bson:"user_id"`
	TotalAmount      float64  `json:"total_amount" bson:"total_amount"`
	CategoryName     string   `json:"category_name" bson:"category_name"`
	TransactionTypes []string `json:"transaction_types" bson:"transaction_types"`
}

type Incomes struct {
	Incomes []*Income `json:"incomes" bson:"incomes"`
}

type BudgetPerformance struct {
	Surplus              float64 `json:"surplus" bson:"surplus"`
	Loss                 float64 `json:"loss" bson:"loss"`
	SurplusPercentage    float32 `json:"surplus_percentage" bson:"surplus_percentage"`
	LossPercentage       float32 `json:"loss_percentage" bson:"loss_percentage"`
	IsInSurplus          bool    `json:"is_in_surplus" bson:"is_in_surplus"`
	CategoriesInSurplus  int32   `json:"categories_in_surplus" bson:"categories_in_surplus"`
	CategoriesInLoss     int32   `json:"categories_in_loss" bson:"categories_in_loss"`
	CategoriesInProgress int32   `json:"categories_in_progress" bson:"categories_in_progress"`
}

type GoalProgress struct {
	Surplus           float64 `json:"surplus" bson:"surplus"`
	WorkingPercentage float32 `json:"working_percentage" bson:"working_percentage"`
	GoalsAchieved     int32   `json:"goals_achieved" bson:"goals_achieved"`
	GoalsFailed       int32   `json:"goals_failed" bson:"goals_failed"`
	GoalsInProgress   int32   `json:"goals_in_progress" bson:"goals_in_progress"`
}
