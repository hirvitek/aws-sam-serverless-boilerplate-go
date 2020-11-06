package logger

import (
	"encoding/json"
	"os"
)

const (
	errorLvl  = "ERROR"
	normalLvl = "NORMAL"
)

type logStructure struct {
	Level   string `json:"level"`
	Name    string `json:"name"`
	Metadata map[string]string `json:"metadata"`
	Message interface{} `json:"message"`
}

type Port interface {
	Info(name string, message interface{})
	Failure(err error, params ...string)
	SetMetadata(meta map[string]string)
}

type logger struct {
	level    string
	message  interface{}
	metadata map[string]string
	enc      *json.Encoder
}

func New() *logger {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	return &logger{
		enc: enc,
		metadata: map[string]string{},
	}
}

func (l *logger) SetMetadata(meta map[string]string)  {
	for key, value := range meta {
		l.metadata[key] = value
	}
}

func (l logger) Info(name string, message interface{}) {
	fullMsg := logStructure{
		Level:   normalLvl,
		Name:    name,
		Message: message,
		Metadata: l.metadata,
	}
	l.enc.Encode(fullMsg)
}

// This is called Failure because if I called it Error it would implement the error type and therefore confusing
func (l logger) Failure(err error, params ...string) {
	// Prevent index out of range if no params are passed
	if len(params) == 0 {
		params = append(params, "")
	}
	
	msg := logStructure{
		Level:   errorLvl,
		Name:    params[0],
		Message: err.Error(),
		Metadata: l.metadata,
	}
	l.enc.Encode(msg)
}
