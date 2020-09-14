package domain

import (
    "time"
)

//WalletPosition represents a model
type WalletPosition struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
    Total int `json:"total"`
    ExpenseID string `json:"expenseId"`
}