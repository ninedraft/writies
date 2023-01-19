package writies_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ninedraft/writies"
)

func TestWriteString(test *testing.T) {
	test.Parallel()
	const input = "Hello, World!"

	got := &bytes.Buffer{}
	n, err := writies.WriteString(got, input)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(input) {
		test.Fatal("unexpected written bytes count", n)
	}
}

func TestWriteString_Empty(test *testing.T) {
	test.Parallel()

	got := &bytes.Buffer{}
	n, err := writies.WriteString(got, "")

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
}

func TestWriteString_Error(test *testing.T) {
	test.Parallel()
	const input = "Hello, World!"

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.WriteString(wr, input)

	if !errors.Is(err, wr.err) {
		test.Fatal("unexpected error", err)
	}
	if n != wr.n {
		test.Fatal("unexpected written bytes count", n)
	}
}

func TestWriteStrings(test *testing.T) {
	test.Parallel()
	input := []string{"Hello", ", ", "World", "!"}
	const expected = "Hello, World!"

	got := &bytes.Buffer{}
	n, err := writies.WriteStrings(got, input...)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != expected {
		test.Fatal("unexpected result", got.String())
	}
}

func TestWriteStrings_Empty(test *testing.T) {
	test.Parallel()

	got := &bytes.Buffer{}
	n, err := writies.WriteStrings(got)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != "" {
		test.Fatal("unexpected result", got.String())
	}
}

func TestWriteStrings_Error(test *testing.T) {
	test.Parallel()
	input := []string{"Hello", ", ", "World", "!"}

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.WriteStrings(wr, input...)

	if !errors.Is(err, wr.err) {
		test.Fatal("unexpected error", err)
	}
	if n != wr.n {
		test.Fatal("unexpected written bytes count", n)
	}
}

func TestJoinStrings(test *testing.T) {
	test.Parallel()
	input := []string{"a", "b", "c"}
	const sep = "1"
	expected := strings.Join(input, sep)

	got := &bytes.Buffer{}
	n, err := writies.JoinStrings(got, input, "1")

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != expected {
		test.Fatal("unexpected result", got)
	}
}

func TestJoinStrings_Empty(test *testing.T) {
	test.Parallel()

	got := &bytes.Buffer{}
	n, err := writies.JoinStrings(got, nil, "1")

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != "" {
		test.Fatal("unexpected result", got)
	}
}

func TestJoinStrings_Error(test *testing.T) {
	test.Parallel()
	input := []string{"a", "b", "c"}
	const expected = "a1b1c"

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.JoinStrings(wr, input, "1")

	if !errors.Is(err, wr.err) {
		test.Fatal("unexpected error", err)
	}
	if n != wr.n {
		test.Fatal("unexpected written bytes count", n)
	}
	if wr.calls != 1 {
		test.Fatal("unexpected calls count", wr.calls)
	}
}

func TestWriteStringN(test *testing.T) {
	test.Parallel()
	const N = 100
	const input = "Hello, World!"
	var expected = strings.Repeat(input, N)

	got := &bytes.Buffer{}
	n, err := writies.WriteStringN(got, input, N)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != expected {
		test.Fatal("unexpected result", got.String())
	}
}

func TestWriteStringN_Negative(test *testing.T) {
	test.Parallel()
	const input = "Hello, World!"

	got := &bytes.Buffer{}
	n, err := writies.WriteStringN(got, input, -1)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != "" {
		test.Fatal("unexpected result", got.String())
	}
}

func TestWriteStringN_Zero(test *testing.T) {
	test.Parallel()
	const input = "Hello, World!"

	got := &bytes.Buffer{}
	n, err := writies.WriteStringN(got, input, 0)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.String() != "" {
		test.Fatal("unexpected result", got.String())
	}
}

func TestWriteStringN_Error(test *testing.T) {
	test.Parallel()
	const N = 100
	const input = "Hello, World!"

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.WriteStringN(wr, input, N)

	if !errors.Is(err, wr.err) {
		test.Fatal("unexpected error", err)
	}
	if n != wr.n {
		test.Fatal("unexpected written bytes count", n)
	}
	if wr.calls != 1 {
		test.Fatal("unexpected calls count", wr.calls)
	}
}

type errWriter struct {
	calls int
	err   error
	n     int
}

func (ew *errWriter) Write([]byte) (int, error) {
	ew.calls++
	return ew.n, ew.err
}
