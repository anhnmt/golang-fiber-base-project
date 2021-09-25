package mixin

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xdorro/golang-fiber-base-project/internal/common"
)

type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				//return utils.UUID()
				return common.GenerateUUID()
			}).
			Annotations(
				entgql.OrderField("ID"),
			),
	}
}

// Edges of the BaseMixin.
func (BaseMixin) Edges() []ent.Edge {
	return nil
}
