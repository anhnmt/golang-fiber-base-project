//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../../gqlgen.yml"),
		entgql.WithSchemaPath("../graph/schema/ent.graphqls"),
	)

	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
	}

	if err = entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
