//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	opts := []entc.Option{}

	if err := entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
