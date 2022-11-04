package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Todo struct to describe todo object
type Todo struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid4"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid4"`
	Content   string    `db:"content" json:"content" validate:"required,lte=255"`
	Complete  bool      `db:"complete" json:"complete"`
	TodoAttrs TodoAttrs `db:"todo_attrs" json:"todo_attrs" validate:"required,dive"`
}

// TodoAttrs struct to describe todo attributes
type TodoAttrs struct {
	Color string `json:"color"`
}

// Value make the TodoAttrs struct implement the driver.Value interface.
// This method simples returns the JSON-encoded representation of the struct.
func (t TodoAttrs) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// Scan make the TodoAttrs struct implement the sql.Scanner interface.
// This method decodes a JSON-encoded value into the struct fields.
func (t *TodoAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &t)
}
