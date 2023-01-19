package writies_test

import (
	"bytes"
	"fmt"

	"github.com/ninedraft/writies"
)

func ExampleWriteString() {
	buf := &bytes.Buffer{}
	_, _ = writies.WriteString(buf, "Hello, world!")

	fmt.Println(buf.String()) // Hello, world!
}

func ExampleWriteStrings() {
	buf := &bytes.Buffer{}
	_, _ = writies.WriteStrings(buf, "Hello", ", ", "world!")

	fmt.Println(buf.String()) // Hello, world!
}

func ExampleJoinStrings() {
	buf := &bytes.Buffer{}
	_, _ = writies.JoinStrings(buf, []string{"Hello", "world"}, ", ")

	fmt.Println(buf.String()) // Hello, world
}

func ExampleRepeatString() {
	buf := &bytes.Buffer{}
	_, _ = writies.RepeatString(buf, "a!", 5)

	fmt.Println(buf.String()) // a!a!a!a!a!
}

func ExampleWriteMany() {
	buf := &bytes.Buffer{}
	_, _ = writies.WriteMany(buf, []byte("Hello"), []byte(", "), []byte("world!"))

	fmt.Println(buf.String()) // Hello, world!
}

func ExampleJoin() {
	buf := &bytes.Buffer{}
	_, _ = writies.Join(buf,
		[][]byte{[]byte("Hello"), []byte("world")},
		[]byte(", "),
	)

	fmt.Println(buf.String()) // Hello, world
}

func ExampleRepeat() {
	buf := &bytes.Buffer{}
	_, _ = writies.Repeat(buf, []byte("a!"), 5)

	fmt.Println(buf.String()) // a!a!a!a!a!
}
