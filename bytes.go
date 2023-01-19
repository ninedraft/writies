package writies

import "io"

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
