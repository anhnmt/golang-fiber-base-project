package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/schema/mixin"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("slug").
			NotEmpty(),

		field.Int("status").
			Default(1),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
		mixin.TimeMixin{},
	}
}

func (Permission) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("slug", "status"),
	}
}
