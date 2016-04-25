package main

import (
	"io/ioutil"
	"os"

	"github.com/ckaznocha/protoc-gen-lint/linter"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoc "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var (
		err              error
		data             []byte
		protoFiles       []*descriptor.FileDescriptorProto
		totalErrors      int
		generatorRequest = new(protoc.CodeGeneratorRequest)
	)

	data, err = ioutil.ReadAll(os.Stdin)
	panicOnError(err)

	err = proto.Unmarshal(data, generatorRequest)
	panicOnError(err)

	protoFiles = generatorRequest.GetProtoFile()
	for _, file := range protoFiles {
		numErrors, err := linter.LintProtoFile(file, os.Stderr)
		panicOnError(err)
		totalErrors += numErrors
	}

	os.Exit(totalErrors)
}
