package respreader

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"gopkg.in/kothar/brotli-go.v0/dec"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Decode(Response *http.Response) ([]byte, error) {
	switch strings.ToLower(Response.Header.Get("Content-Encoding")) {
	case "gzip":
		reader, err := gzip.NewReader(Response.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		var b bytes.Buffer
		if _, err = b.ReadFrom(reader); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	case "zlib", "deflate":
		reader, err := zlib.NewReader(Response.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		var b bytes.Buffer
		if _, err = b.ReadFrom(reader); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	case "br":
		reader, err := ioutil.ReadAll(Response.Body)
		if err != nil {
			return nil, err
		}
		decompressed, err := dec.DecompressBuffer(reader, make([]byte, 0))
		if err != nil {
			return nil, err
		}
		return decompressed, nil
	default:
		return ioutil.ReadAll(Response.Body)
	}
	log.Fatal()
}
