package writies

import "io"

// WriteMany writes multiple byte slices to writer.
// It does not allocate memory, so it can be used in performance critical code.
//
// Returns number of bytes written and error, if any occurred.
// It stops writing on first error.
// It does not write anything if strs is empty.
func WriteMany(dst io.Writer, strs ...[]byte) (int, error) {
	var written int
	for _, str := range strs {
		n, err := dst.Write(str)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

// Join joins byte slices with separator and writes result to writer.
// It does not allocate memory, so it can be used in performance critical code.
//
// Returns number of bytes written and error, if any occurred.
// It stops writing on first error.
// It does not write anything if strs is empty.
func Join(dst io.Writer, strs [][]byte, sep []byte) (int, error) {
	var written int
	for i, str := range strs {
		if i > 0 {
			n, err := dst.Write(sep)
			written += n
			if err != nil {
				return written, err
			}
		}
		n, err := dst.Write(str)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

// WriteN writes byte slice n times to writer.
// It does not allocate memory, so it can be used in performance critical code.
//
// Returns number of bytes written and error, if any occurred.
// It stops writing on first error.
// It does not write anything if n <= 0.
func WriteN(dst io.Writer, str []byte, n int) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	var written int
	for i := 0; i < n; i++ {
		n, err := dst.Write(str)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}
