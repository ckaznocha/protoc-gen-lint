package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ckaznocha/protoc-gen-lint/linter"
	"github.com/golang/protobuf/proto"
	protoc "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// SortImports represents the parameter, which can be specified to the tool invocation
// to enable checking, whether the proto file imports are sorted alphabetically.
const SortImports = "sort_imports"

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		totalErrors      int
		generatorRequest protoc.CodeGeneratorRequest
		parameters       struct {
			SortImports bool
		}
	)

	for _, p := range strings.Split(generatorRequest.GetParameter(), ",") {
		switch p {
		case SortImports:
			parameters.SortImports = true
		default:
			fmt.Fprintf(os.Stderr, "Unmatched parameter: %s", p)
			totalErrors++
		}
	}
	if totalErrors != 0 {
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(os.Stdin)
	panicOnError(err)

	panicOnError(proto.Unmarshal(data, &generatorRequest))

	for _, file := range generatorRequest.GetProtoFile() {
		numErrors, err := linter.LintProtoFile(linter.Config{
			ProtoFile:   file,
			OutFile:     os.Stderr,
			SortImports: parameters.SortImports,
		})
		panicOnError(err)
		totalErrors += numErrors
	}

	os.Exit(totalErrors)
}
