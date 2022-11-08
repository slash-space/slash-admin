//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
)

func main() {
	schemaPath := "./schema"

	graph, err := entc.LoadGraph(schemaPath, &gen.Config{})

	if err != nil {
		fmt.Println("entproto: failed loading ent graph: %v", err)
	}
	if err := entproto.Generate(graph); err != nil {
		fmt.Println("entproto: failed generating proto file: %s", err)
	}

	fmt.Println("Code generation completed successfully")

}
