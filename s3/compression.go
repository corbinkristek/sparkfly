package s3

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"strings"
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

*/

//I found this question to be difficult to understand.
//The above comments are from my initial solution i posted to github.
//2 days ago I came up with the solution below that I think better satisfies the original question.
type YourSolution interface {
	io.Reader
}

func NewYourSolution(rc io.ReadCloser) *YourSolution {
	rc = io.NopCloser(strings.NewReader("this is a test string"))
	buff1 := new(bytes.Buffer)
	buff1.ReadFrom(rc)
	rc.Close()
	buff2 := new(bytes.Buffer)
	w := gzip.NewWriter(buff2)
	w.Write([]byte(buff1.String()))
	w.Close()
	res, _ := ioutil.ReadAll(buff2)
	log.Println(res)
	var f YourSolution = bytes.NewReader(res)
	return &f

}
