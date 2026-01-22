package main

import "strings"

type Reader interface {
	Read() []byte
}

type Writer interface {
	Write(p []byte) int
}

type ReaderWriter interface {
	Reader
	Writer
}

type UpperReaderWriter struct {
	UpperString string
}

func (rw *UpperReaderWriter) Read() []byte {
	return []byte(rw.UpperString)
}

func (rw *UpperReaderWriter) Write(p []byte) int {
	rw.UpperString = strings.ToUpper(string(p))
	return len(p)
}
