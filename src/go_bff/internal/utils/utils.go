package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

const jsonMaxBytes = 1024 * 10

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(jsonMaxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, data any) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}
