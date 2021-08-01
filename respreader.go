package respreader

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/kothar/brotli-go.v0/dec"
	"gopkg.in/kothar/brotli-go.v0/enc"
)

const (
	TypeGZIP    = "gzip"
	TypeZLIB    = "zlib"
	TypeDeflate = "deflate"
	TypeBrotli  = "br"
)

func Decode(Response *http.Response) ([]byte, error) {
	var (
		reader io.ReadCloser
		err    error
		b      bytes.Buffer
	)

	reader = nil

	switch strings.ToLower(Response.Header.Get("Content-Encoding")) {

	case TypeGZIP:
		if reader, err = gzip.NewReader(Response.Body); err != nil {
			return nil, err
		}

	case TypeZLIB, TypeDeflate:
		if reader, err = zlib.NewReader(Response.Body); err != nil {
			return nil, err
		}

	case TypeBrotli:
		reader = dec.NewBrotliReader(Response.Body)

	default:
		return ioutil.ReadAll(Response.Body)
	}

	if reader == nil {
		return nil, errors.New("decode failed")
	}

	defer func() {
		_ = reader.Close()
	}()

	if _, err = b.ReadFrom(reader); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func DecodeBytes(buffer []byte, algo string) ([]byte, error) {
	var (
		reader io.ReadCloser
		err    error
		b      bytes.Buffer
	)

	reader = nil

	switch strings.ToLower(algo) {

	case TypeGZIP:
		if reader, err = gzip.NewReader(bytes.NewBuffer(buffer)); err != nil {
			return nil, err
		}

	case TypeZLIB, TypeDeflate:
		if reader, err = zlib.NewReader(bytes.NewBuffer(buffer)); err != nil {
			return nil, err
		}

	case TypeBrotli:
		reader = dec.NewBrotliReader(bytes.NewBuffer(buffer))

	default:
		return ioutil.ReadAll(bytes.NewBuffer(buffer))
	}

	if reader == nil {
		return nil, errors.New("decode failed")
	}

	defer func() {
		_ = reader.Close()
	}()

	if _, err = b.ReadFrom(reader); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func EncodeBytes(buffer []byte, algo string) ([]byte, error) {
	var (
		writer io.WriteCloser
	)

	if buffer == nil || len(buffer) == 0 {
		return nil, nil
	}

	buff := bytes.Buffer{}

	switch strings.ToLower(algo) {

	case TypeGZIP:
		writer = gzip.NewWriter(&buff)

	case TypeZLIB, TypeDeflate:
		writer = zlib.NewWriter(&buff)

	case TypeBrotli:
		writer = enc.NewBrotliWriter(nil, &buff)

	default:
		return nil, errors.New("encode failed")
	}

	if writer == nil {
		return nil, errors.New("encode failed")
	}

	if _, err := writer.Write(buffer); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
