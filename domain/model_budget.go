package domain

import (
    "time"
)

//Budget represents a model
type Budget struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
    Max int `json:"max"`
    ExpenseID string `json:"expenseId"`
}