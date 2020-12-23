package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
)
// Chat is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Chat struct {
    ID uuid.UUID `json:"id" db:"id"`
    TextStr nulls.String `json:"text_str" db:"text_str"`
    UserStr nulls.String `json:"user_str" db:"user_str"`
    ChatroomStr nulls.String `json:"chatroom_str" db:"chatroom_str"`
    UrlStr nulls.String `json:"url_str" db:"url_str"`
    IDInt nulls.Int `json:"id_int" db:"id_int"`
    DateDate nulls.Time `json:"date_date" db:"date_date"`
    Spare1Str nulls.String `json:"spare1_str" db:"spare1_str"`
    Spare2Str nulls.String `json:"spare2_str" db:"spare2_str"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Chat) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Chats is not required by pop and may be deleted
type Chats []Chat

// String is not required by pop and may be deleted
func (c Chats) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Chat) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Chat) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Chat) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
