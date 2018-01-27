package linter

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

type protoBufError struct {
	path        []int32
	errorCode   errorCode
	errorString string
}

func (p *protoBufError) getSourceLineNumber(
	protoSource *descriptor.SourceCodeInfo,
) (int32, int32) {

	p.path = append(p.path, pathMessageName)
	var sourceLen = len(p.path)

	for _, v := range protoSource.GetLocation() {
		var (
			curPath = v.GetPath()
			curLen  = len(curPath)
			skip    bool
		)

		// Started out using reflect.DeepEqual.
		// This is more verbose but should be much faster.
		if curLen == sourceLen {
			for i := 0; i < sourceLen; i++ {
				if curPath[i] != p.path[i] {
					skip = true
					break
				}
			}

			if skip {
				continue
			}

			var (
				span = v.GetSpan()
				line = span[0] + 1
				col  = span[1] + 1
			)

			return line, col
		}
	}
	return 0, 0
}
