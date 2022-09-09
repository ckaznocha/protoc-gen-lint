package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/ckaznocha/protoc-gen-lint/linter"
)

// SortImports represents the parameter, which can be specified to the tool invocation
// to enable checking, whether the proto file imports are sorted alphabetically.
const SortImports = "sort_imports"

func main() {
	var (
		totalErrors      int
		generatorRequest pluginpb.CodeGeneratorRequest
		parameters       struct {
			SortImports bool
		}
	)

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := proto.Unmarshal(data, &generatorRequest); err != nil {
		panic(err)
	}

	for _, p := range strings.Split(generatorRequest.GetParameter(), ",") {
		switch strings.TrimSpace(p) {
		case "":
			continue
		case SortImports:
			parameters.SortImports = true
		default:
			fmt.Fprintf(os.Stderr, "Unmatched parameter: %s", p)
			os.Exit(1)
		}
	}

	for _, file := range generatorRequest.GetProtoFile() {
		numErrors, err := linter.LintProtoFile(linter.Config{
			ProtoFile:   file,
			OutFile:     os.Stderr,
			SortImports: parameters.SortImports,
		})
		if err != nil {
			panic(err)
		}

		totalErrors += numErrors
	}

	os.Exit(totalErrors)
}
