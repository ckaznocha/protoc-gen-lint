package linter

import (
	"fmt"
	"io"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const (
	//Path Types
	pathMessage        = 4
	pathEnumType       = 5
	pathEnumValue      = 2
	pathService        = 6
	pathRPCMethod      = 2
	pathMessageName    = 1
	pathMessageField   = 2
	pathMessageMessage = 3
	pathMessageEnum    = 4

	//Error Types
	errorMessageCase   = 1
	errorFieldCase     = 2
	errorEnumTypeCase  = 3
	errorEnumValueCase = 4
	errorServiceCase   = 5
	errorRPCMethodCase = 6
)

var linterErrors = map[int]string{
	1: "Use CamelCase (with an initial capital) for message names.",
	2: "Use underscore_separated_names for field names.",
	3: "Use CamelCase (with an initial capital) for enum type names.",
	4: "Use CAPITALS_WITH_UNDERSCORES  for enum value names.",
	5: "Use CamelCase (with an initial capital) for service names.",
	6: "Use CamelCase (with an initial capital) for RPC method names.",
}

// LintProtoFile takes a file name, proto file description, and a file.
// It checks the file for errors and writes them to the output file
func LintProtoFile(
	fileName string,
	protoFile *descriptor.FileDescriptorProto,
	outFile io.WriteCloser,
) (int, error) {

	defer outFile.Close()

	var (
		errors      = protoBufErrors{}
		protoSource = protoFile.GetSourceCodeInfo()
	)

	for i, v := range protoFile.GetMessageType() {
		errors.lintProtoMessage(int32(i), pathMessage, []int32{}, v)
	}

	for i, v := range protoFile.GetEnumType() {
		errors.lintProtoEnumType(int32(i), pathEnumType, []int32{}, v)
	}

	for i, v := range protoFile.GetService() {
		errors.lintProtoService(int32(i), v)
	}

	for _, v := range errors {
		line, col := v.getSourceLineNumber(protoSource)
		_, err := outFile.Write([]byte(
			fmt.Sprintf(
				"%s:%d:%d: '%s' - %s\n",
				fileName,
				line,
				col,
				v.errorString,
				linterErrors[v.errorCode],
			),
		))
		if err != nil {
			return len(errors), err
		}
	}

	return len(errors), nil

}
