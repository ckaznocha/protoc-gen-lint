package main

import (
	"io/ioutil"
	"os"

	"github.com/ckaznocha/protoc-gen-lint/linter"
	"github.com/golang/protobuf/proto"
	protoc "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var (
		totalErrors      int
		generatorRequest protoc.CodeGeneratorRequest
	)

	data, err := ioutil.ReadAll(os.Stdin)
	panicOnError(err)

	panicOnError(proto.Unmarshal(data, &generatorRequest))

	for _, file := range generatorRequest.GetProtoFile() {
		numErrors, err := linter.LintProtoFile(file, os.Stderr)
		panicOnError(err)
		totalErrors += numErrors
	}

	os.Exit(totalErrors)
}
