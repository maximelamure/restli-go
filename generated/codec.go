package generated

import (
	"encoding/json"
	"errors"
	"io"
)

type reader interface {
	Read(io.Reader, interface{}) error
	ContentType() string
}

// Can use different codecs in the future based on config
func getCodec() reader {
	return newJSONReader()
}

type jsonReader struct {
	contentType string
}

func newJSONReader() reader {
	return &jsonReader{
		contentType: "application/json",
	}
}

func (j *jsonReader) ContentType() string {
	return j.contentType
}

func (j *jsonReader) Read(r io.Reader, data interface{}) error {
	if errJSON := json.NewDecoder(r).Decode(data); errJSON != nil && errJSON != io.EOF {
		return errors.New("could not decode object" + errJSON.Error())
	}

	return nil
}
