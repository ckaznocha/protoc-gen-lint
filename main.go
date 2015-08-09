package main

import (
	"io/ioutil"
	"os"

	"github.com/ckaznocha/proc-gen-lint/linter"
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
		protoFileNames   []string
		totalErrors      int
		generatorRequest = new(protoc.CodeGeneratorRequest)
	)

	data, err = ioutil.ReadAll(os.Stdin)
	panicOnError(err)

	err = proto.Unmarshal(data, generatorRequest)
	panicOnError(err)

	protoFiles = generatorRequest.GetProtoFile()
	protoFileNames = generatorRequest.GetFileToGenerate()

	for i := 0; i < len(protoFiles); i++ {
		numErrors, err := linter.LintProtoFile(
			protoFileNames[i],
			protoFiles[i],
			os.Stderr,
		)
		panicOnError(err)
		totalErrors += numErrors
	}

	os.Exit(totalErrors)
}
