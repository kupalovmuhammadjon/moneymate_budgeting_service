package models

type CreateGoal struct {
	UserID        string  `json:"user_id" bson:"user_id"`
	Name          string  `json:"name" bson:"name"`
	TargetAmount  float64 `json:"target_amount" bson:"target_amount"`
	CurrentAmount float64 `json:"current_amount" bson:"current_amount"`
	Deadline      string  `json:"deadline" bson:"deadline"`
	Status        string  `json:"status" bson:"status"`
}

type Goal struct {
	ID            string  `json:"id" bson:"_id,omitempty"` 
	UserID        string  `json:"user_id" bson:"user_id"`
	Name          string  `json:"name" bson:"name"`
	TargetAmount  float64 `json:"target_amount" bson:"target_amount"`
	CurrentAmount float64 `json:"current_amount" bson:"current_amount"`
	Deadline      string  `json:"deadline" bson:"deadline"`
	Status        string  `json:"status" bson:"status"`
	CreatedAt     string  `json:"created_at" bson:"created_at"`
	UpdatedAt     string  `json:"updated_at" bson:"updated_at"`
}

type GoalFilter struct {
	Page          int32   `json:"page" bson:"page"`
	Limit         int32   `json:"limit" bson:"limit"`
	UserID        string  `json:"user_id" bson:"user_id"`
	Name          string  `json:"name" bson:"name"`
	TargetAmount  float64 `json:"target_amount" bson:"target_amount"`
	CurrentAmount float64 `json:"current_amount" bson:"current_amount"`
	Deadline      string  `json:"deadline" bson:"deadline"`
	Status        string  `json:"status" bson:"status"`
}

type Goals struct {
	Goals []*Goal `json:"goals" bson:"goals"`
}
