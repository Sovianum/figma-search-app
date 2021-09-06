package util

import (
	"encoding/json"
	"io"
)

func UnmarshalFromReaderCloser(rc io.ReadCloser, body interface{}) error {
	b, err := io.ReadAll(rc)
	if err != nil {
		return err
	}

	defer rc.Close()

	return json.Unmarshal(b, body)
}
