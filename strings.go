package writies

import (
	"io"
	"unsafe"
)

// WriteString writes string to writer.
// It does not allocate memory, so it can be used
// in performance critical code instead of io.WriteString.
//
// Returns number of bytes written and error, if any occurred.
func WriteString(dst io.Writer, str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	head := unsafe.StringData(str)
	p := unsafe.Slice(head, len(str))
	return dst.Write(p)
}

// WriteStrings writes multiple strings to writer.
// It does not allocate memory, so it can be used in performance critical code.
//
// Returns number of bytes written and error, if any occurred.
// It stops writing on first error.
func WriteStrings(dst io.Writer, strs ...string) (int, error) {
	var written int
	for _, str := range strs {
		n, err := WriteString(dst, str)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

// JoinStrings joins strings with separator and writes result to writer.
// It does not allocate memory, so it can be used in performance critical code.
//
// Returns number of bytes written and error, if any occurred.
// It stops writing on first error.
// It does not write anything if strs is empty.
func JoinStrings(dst io.Writer, strs []string, sep string) (int, error) {
	var written int
	for i, str := range strs {
		if i > 0 {
			n, err := WriteString(dst, sep)
			written += n
			if err != nil {
				return written, err
			}
		}
		n, err := WriteString(dst, str)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

// RepeatString writes string n times to writer.
// It does not allocate memory, so it can be used in performance critical code.
//
// Returns number of bytes written and error, if any occurred.
// It stops writing on first error.
// It does not write anything if n <= 0.
func RepeatString(dst io.Writer, str string, n int) (int, error) {
	if str == "" {
		return 0, nil
	}
	var written int
	for i := 0; i < n; i++ {
		n, err := WriteString(dst, str)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}
