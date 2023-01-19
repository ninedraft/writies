package writies

import (
	"io"
	"unsafe"
)

func WriteString(dst io.Writer, str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	head := unsafe.StringData(str)
	p := unsafe.Slice(head, len(str))
	return dst.Write(p)
}

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

func WriteStringN(dst io.Writer, str string, n int) (int, error) {
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
