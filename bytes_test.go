package writies_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ninedraft/writies"
)

func TestWriteMany(test *testing.T) {
	test.Parallel()
	input := [][]byte{
		[]byte("Hello, "),
		[]byte("World!"),
	}
	expected := []byte("Hello, World!")

	got := &bytes.Buffer{}
	n, err := writies.WriteMany(got, input...)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if !bytes.Equal(got.Bytes(), expected) {
		test.Fatal("unexpected output", got)
	}
}

func TestWriteMany_Empty(test *testing.T) {
	test.Parallel()

	got := &bytes.Buffer{}
	n, err := writies.WriteMany(got, nil)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.Len() != 0 {
		test.Fatal("unexpected output", got)
	}
}

func TestWriteMany_Error(test *testing.T) {
	test.Parallel()
	input := [][]byte{
		[]byte("Hello, "),
		[]byte("World!"),
	}

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.WriteMany(wr, input...)

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

func TestJoin(test *testing.T) {
	test.Parallel()
	input := [][]byte{
		[]byte("Hello,"),
		[]byte("World!"),
	}
	expected := []byte("Hello,_World!")

	got := &bytes.Buffer{}
	n, err := writies.Join(got, input, []byte("_"))

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if !bytes.Equal(got.Bytes(), expected) {
		test.Fatal("unexpected output", got)
	}
}

func TestJoin_EmptySep(test *testing.T) {
	test.Parallel()
	input := [][]byte{
		[]byte("Hello,"),
		[]byte("World!"),
	}
	expected := []byte("Hello,World!")

	got := &bytes.Buffer{}
	n, err := writies.Join(got, input, nil)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if !bytes.Equal(got.Bytes(), expected) {
		test.Fatal("unexpected output", got)
	}
}

func TestJoin_Empty(test *testing.T) {
	test.Parallel()

	got := &bytes.Buffer{}
	n, err := writies.Join(got, nil, []byte("_"))

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.Len() != 0 {
		test.Fatal("unexpected output", got)
	}
}

func TestJoin_Error(test *testing.T) {
	test.Parallel()
	input := [][]byte{
		[]byte("Hello,"),
		[]byte("World!"),
	}

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.Join(wr, input, []byte("_"))

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

func TestWriteN(test *testing.T) {
	test.Parallel()
	input := []byte("Hello, World!")
	const N = 100
	expected := bytes.Repeat([]byte("Hello, World!"), N)

	got := &bytes.Buffer{}
	n, err := writies.WriteN(got, input, N)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != len(expected) {
		test.Fatal("unexpected written bytes count", n)
	}
	if !bytes.Equal(got.Bytes(), expected) {
		test.Fatal("unexpected output", got)
	}
}

func TestWriteN_Negative(test *testing.T) {
	test.Parallel()
	input := []byte("Hello, World!")

	got := &bytes.Buffer{}
	n, err := writies.WriteN(got, input, -1)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.Len() != 0 {
		test.Fatal("unexpected output", got)
	}
}

func TestWriteN_Zero(test *testing.T) {
	test.Parallel()
	input := []byte("Hello, World!")

	got := &bytes.Buffer{}
	n, err := writies.WriteN(got, input, 0)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.Len() != 0 {
		test.Fatal("unexpected output", got)
	}
}

func TestWriteN_Empty(test *testing.T) {
	test.Parallel()

	got := &bytes.Buffer{}
	n, err := writies.WriteN(got, nil, 100)

	if err != nil {
		test.Fatal("unexpected error", err)
	}
	if n != 0 {
		test.Fatal("unexpected written bytes count", n)
	}
	if got.Len() != 0 {
		test.Fatal("unexpected output", got)
	}
}

func TestWriteN_Error(test *testing.T) {
	test.Parallel()
	input := []byte("Hello, World!")

	wr := &errWriter{
		err: errors.New("test error"),
		n:   -1000,
	}

	n, err := writies.WriteN(wr, input, 100)

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
