package utils

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var Slogger *slog.Logger

func SloggerInit() *slog.Logger {
	Slogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return Slogger
}

const jsonMaxBytes int = 1024 * 10

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
		Slogger.Error(err.Error())
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		Slogger.Error(err.Error())
		return err
	}

	return nil
}
