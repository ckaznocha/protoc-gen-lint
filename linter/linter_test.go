//go:build go1.18
// +build go1.18

package linter

import (
	"io"
	"os"
	"testing"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/descriptorpb"
)

func FuzzLintProtoFile(f *testing.F) {
	testProto, err := os.ReadFile("testdata/test.textproto")
	if err != nil {
		f.Fatal(err)
	}

	f.Add(testProto)

	f.Fuzz(func(t *testing.T, data []byte) {
		descriptor := &descriptorpb.FileDescriptorProto{}
		if err := prototext.Unmarshal(data, descriptor); err != nil {
			t.Skip()
		}

		if _, err := LintProtoFile(Config{
			ProtoFile:   descriptor,
			OutFile:     io.Discard,
			SortImports: false,
		}); err != nil {
			t.Error(err)
		}
	})
}
