package s3

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
)

func newYourSolution(rc io.ReadCloser) io.Reader {
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
