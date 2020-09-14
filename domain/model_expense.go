package domain

import (
    "time"
)

//Expense represents a model
type Expense struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
    Name string `json:"name"`
}