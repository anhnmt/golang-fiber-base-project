package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/xdorro/golang-fiber-base-project/pkg/ent/schema/mixin"
)

// UserPermission holds the schema definition for the UserPermission entity.
type UserPermission struct {
	ent.Schema
}

// Fields of the UserPermission.
func (UserPermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			NotEmpty(),

		field.String("permission_id").
			NotEmpty(),

		field.Int("status").
			Default(1),
	}
}

func (UserPermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
	}
}

// Edges of the UserPermission.
func (UserPermission) Edges() []ent.Edge {
	return nil
}

func (UserPermission) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("user_id", "status"),
	}
}
