package main

type MyReader struct {
}

func (m MyReader) Read(chunk []byte) (n int, err error) {
	return 0, nil
}

type MyWriter struct {
}

func (m MyWriter) Write(chunk []byte) (n int, err error) {
	return 0, nil
}

func main() {

}
