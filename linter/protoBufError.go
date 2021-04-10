package linter

import "google.golang.org/protobuf/types/descriptorpb"

type protoBufError struct {
	errorString string
	path        []int32
	errorCode   errorCode
}

func (p *protoBufError) getSourceLineNumber(
	protoSource *descriptorpb.SourceCodeInfo,
) (line, col int32) {
	p.path = append(p.path, pathMessageName)
	sourceLen := len(p.path)

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
