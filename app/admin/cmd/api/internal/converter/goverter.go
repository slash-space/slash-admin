//go:build ignore
// +build ignore

package main

import (
	"fmt"
	goverter "github.com/jmattheis/goverter"
	"os"
)

func main() {

	err := goverter.GenerateConverterFile("./generated/generated.go", goverter.GenerateConfig{
		PackageName:             "generated",
		ScanDir:                 "slash-admin/app/admin/cmd/api/internal/converter",
		ExtendMethods:           []string{},
		PackagePath:             "",
		WrapErrors:              false,
		IgnoredUnexportedFields: true,
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
