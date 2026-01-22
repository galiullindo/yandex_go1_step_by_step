package main

import (
	"fmt"
	"strings"
)

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Logger interface {
	Log(message string)
}

type Log struct {
	Level LogLevel
}

func (l *Log) Log(message string) {
	fmt.Printf("%s: %s\n", strings.ToUpper(string(l.Level)), message)
}
