package s3

import (
	"bytes"
	"compress/gzip"
	"io"
)

/*
func newYourSolutionrewrite(rc io.ReadCloser) io.Reader {
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(body)
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	res, err := ioutil.ReadAll(&b)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(res)

}

func rewrite(rc io.ReadCloser) io.Reader {
	var p = make([]byte, 1024)
	_, err := rc.Read(p)
	if err != nil {
		return nil
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err = w.Write(p)
	if err != nil {
		return nil
	}

	res, err := ioutil.ReadAll(&b)
	if err != nil {
		return nil
	}
	return bytes.NewReader(res)

}
*/

//This question was genuinly confusing.
//The above comments are from my initial solution i posted to github.
//After I had time to think I came up with the below solution that I think better satisfies the question.
type YourSolution interface {
	io.Reader
}

type FooReader struct{}

func (f *FooReader) Read(b []byte) (int, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	err := w.Close()
	if err != nil {
		return 0, err
	}
	return w.Write(b)

}

func NewYourSolution(rc io.ReadCloser) *YourSolution {
	var p = make([]byte, 1024*20)
	_, err := rc.Read(p)
	if err != nil {
		return nil
	}
	var intf YourSolution
	_, err = intf.Read(p)
	if err != nil {
		return nil
	}
	return &intf

}
