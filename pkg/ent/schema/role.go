package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/schema/mixin"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("slug").
			NotEmpty(),

		field.Int("status").
			Default(1),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
		mixin.TimeMixin{},
	}
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("slug", "status"),
	}
}
