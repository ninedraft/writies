package writies_test

import (
	"io"
	"testing"

	"github.com/ninedraft/writies"
)

var writeStringDst testWriter
var writeStringN int
var writeStringErr error

func BenchmarkWriteString(b *testing.B) {
	const str = "Hello, World!"

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeStringN, writeStringErr = writies.WriteString(&writeStringDst, str)
	}
}

var writeStringIODst testWriter
var writeStringION int
var writeStringIOErr error

func BenchmarkWriteStringIO(b *testing.B) {
	const str = "Hello, World!"

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeStringION, writeStringIOErr = io.WriteString(&writeStringIODst, str)
	}
}

type testWriter struct {
	value []byte
}

func (w *testWriter) Write(p []byte) (int, error) {
	w.value = p
	return len(p), nil
}
