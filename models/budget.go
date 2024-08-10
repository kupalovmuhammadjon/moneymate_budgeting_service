package models

type CreateBudget struct {
	UserID     string  `json:"user_id" bson:"user_id"`
	CategoryID string  `json:"category_id" bson:"category_id"`
	Amount     float64 `json:"amount" bson:"amount"`
	Period     string  `json:"period" bson:"period"`
	StartDate  string  `json:"start_date" bson:"start_date"`
	EndDate    string  `json:"end_date" bson:"end_date"`
}

type Budget struct {
	ID         string  `json:"id" bson:"_id,omitempty"`
	UserID     string  `json:"user_id" bson:"user_id"`
	CategoryID string  `json:"category_id" bson:"category_id"`
	Amount     float64 `json:"amount" bson:"amount"`
	Period     string  `json:"period" bson:"period"`
	StartDate  string  `json:"start_date" bson:"start_date"`
	EndDate    string  `json:"end_date" bson:"end_date"`
	CreatedAt  string  `json:"created_at" bson:"created_at"`
	UpdatedAt  string  `json:"updated_at" bson:"updated_at"`
}

type BudgetFilter struct {
	Page       int32   `json:"page" bson:"page"`
	Limit      int32   `json:"limit" bson:"limit"`
	UserID     string  `json:"user_id" bson:"user_id"`
	CategoryID string  `json:"category_id" bson:"category_id"`
	Amount     float64 `json:"amount" bson:"amount"`
	Period     string  `json:"period" bson:"period"`
	StartDate  string  `json:"start_date" bson:"start_date"`
	EndDate    string  `json:"end_date" bson:"end_date"`
}

type Budgets struct {
	Budgets []*Budget `json:"budgets" bson:"budgets"`
}
