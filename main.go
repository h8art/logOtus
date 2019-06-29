package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type OtusEvent interface {
	Log() []byte
}

func (h *HwAccepted) Log() []byte {
	return []byte(time.Now().Format("2006-01-02") + " accepted " + strconv.Itoa(h.Id) + " " + strconv.Itoa(h.Grade) + "\n")
}
func (h *HwSubmitted) Log() []byte {
	return []byte(time.Now().Format("2006-01-02") + " submitted " + strconv.Itoa(h.Id) + " " + h.Comment + "\n")
}
func LogOtusEvent(e OtusEvent, w io.Writer) {
	_, err := w.Write(e.Log())
	if err != nil {
		panic(err)
	}
}

func main() {
	var logAccepted OtusEvent
	var logSubmitted OtusEvent
	var writer bytes.Buffer

	logAccepted = &HwAccepted{12, 3}
	logSubmitted = &HwSubmitted{11, "my code", "great code"}

	LogOtusEvent(logAccepted, &writer)
	LogOtusEvent(logSubmitted, &writer)

	fmt.Printf("%s", writer.String())
}
