package golang

// MyReader - Custom reader implementation
type MyReader struct{}

func (reader MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}
