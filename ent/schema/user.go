package schema

import (
	"errors"
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().
		Validate(func(email string) error {
			if _, err := mail.ParseAddress(email); err != nil {
				return errors.New("email format is incorrect")
			}
			return nil
		}),
		field.String("password").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("refresh_token").Optional(),
		field.Time("last_signin_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
