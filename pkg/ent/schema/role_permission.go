package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/schema/mixin"
)

// RolePermission holds the schema definition for the RolePermission entity.
type RolePermission struct {
	ent.Schema
}

// Fields of the RolePermission.
func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_id").
			NotEmpty(),

		field.String("permission_id").
			NotEmpty(),

		field.Int("status").
			Default(1),
	}
}

func (RolePermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
	}
}

// Edges of the RolePermission.
func (RolePermission) Edges() []ent.Edge {
	return nil
}

func (RolePermission) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("role_id", "status"),
	}
}
