package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/schema/mixin"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			NotEmpty(),

		field.String("role_id").
			NotEmpty(),

		field.Int("status").
			Default(1),
	}
}

func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return nil
}

func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("user_id", "status"),
	}
}
