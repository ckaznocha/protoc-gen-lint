package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/ckaznocha/protoc-gen-lint/linter"
)

type BoolFlag interface {
	IsBoolFlag() bool
}

var errLint = errors.New("encountered lint errors")

func main() {
	var flags flag.FlagSet
	sortImports := flags.Bool("sort_imports", false, "check whether or not the proto "+
		"file imports are sorted alphabetically")

	protogen.Options{
		ParamFunc: func(param, value string) error {
			// For backwards compatibility, treat a present flag with an empty
			// string value as boolean true. The Go flag module parsing
			// handles this, but the protogen flag parse behavior doesn't,
			// because it doesn't parse flags the same way.
			f := flags.Lookup(param)
			if f != nil {
				if _, ok := f.Value.(BoolFlag); ok {
					value = "true"
				}
			}

			return flags.Set(param, value)
		},
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		totalErrors := 0
		for _, f := range gen.Files {
			numErrors, err := linter.LintProtoFile(linter.Config{
				ProtoFile:   f.Proto,
				OutFile:     os.Stderr,
				SortImports: *sortImports,
			})
			if err != nil {
				return err
			}

			totalErrors += numErrors
		}

		if totalErrors > 0 {
			return fmt.Errorf("%w: %d total", errLint, totalErrors)
		}

		return nil
	})
}
