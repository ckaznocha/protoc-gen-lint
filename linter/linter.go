package linter

import (
	"fmt"
	"io"

	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Path Types.
	pathMessage        = 4
	pathEnumType       = 5
	pathEnumValue      = 2
	pathService        = 6
	pathRPCMethod      = 2
	pathMessageName    = 1
	pathMessageField   = 2
	pathMessageMessage = 3
	pathMessageEnum    = 4
)

type (
	errorCode        int
	errorDescription string
)

const (
	// Error Types.
	errorImportOrder errorCode = iota
	errorMessageCase
	errorFieldCase
	errorEnumTypeCase
	errorEnumValueCase
	errorServiceCase
	errorRPCMethodCase
)

var linterErrors = []errorDescription{ //nolint:gochecknoglobals // Global enum values.
	"Sort import statements alphabetically.",
	"Use CamelCase (with an initial capital) for message names.",
	"Use underscore_separated_names for field names.",
	"Use CamelCase (with an initial capital) for enum type names.",
	"Use CAPITALS_WITH_UNDERSCORES  for enum value names.",
	"Use CamelCase (with an initial capital) for service names.",
	"Use CamelCase (with an initial capital) for RPC method names.",
}

type Config struct {
	ProtoFile   *descriptorpb.FileDescriptorProto
	OutFile     io.Writer
	SortImports bool
}

// LintProtoFile takes a file name, proto file description, and a file.
// It checks the file for errors and writes them to the output file.
func LintProtoFile(conf Config) (int, error) {
	var (
		errors      = protoBufErrors{}
		protoSource = conf.ProtoFile.GetSourceCodeInfo()
	)

	if conf.SortImports {
		errors.lintImportOrder(conf.ProtoFile.GetDependency())
	}

	for i, v := range conf.ProtoFile.GetMessageType() {
		errors.lintProtoMessage(int32(i), pathMessage, []int32{}, v)
	}

	for i, v := range conf.ProtoFile.GetEnumType() {
		errors.lintProtoEnumType(int32(i), pathEnumType, []int32{}, v)
	}

	for i, v := range conf.ProtoFile.GetService() {
		errors.lintProtoService(int32(i), v)
	}

	for _, v := range errors {
		line, col := v.getSourceLineNumber(protoSource)
		fmt.Fprintf(
			conf.OutFile,
			"%s:%d:%d: '%s' - %s\n",
			conf.ProtoFile.GetName(),
			line,
			col,
			v.errorString,
			linterErrors[v.errorCode],
		)
	}

	return len(errors), nil
}
