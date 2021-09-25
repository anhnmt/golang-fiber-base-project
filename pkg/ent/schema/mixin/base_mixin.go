package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/util"
)

type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable().
			DefaultFunc(func() int64 {
				conf := config.Config()

				return util.GenerateSnowflakeID(conf.AppNode).Int64()
			}),
	}
}

// Edges of the BaseMixin.
func (BaseMixin) Edges() []ent.Edge {
	return nil
}
