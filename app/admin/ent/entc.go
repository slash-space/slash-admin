//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
)

func main() {
	ex, err := entgql.NewExtension()

	if err != nil {
		panic("creating entgql extension")
	}

	opts := []entc.Option{
		entc.TemplateDir("./template"),
		entc.Extensions(ex),
	}

	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureModifier,
			gen.FeatureUpsert,
		},
	}, opts...); err != nil {
		fmt.Println("running entgo codegen:", err)
		return
	}
	fmt.Println("Code generation completed successfully")

}
