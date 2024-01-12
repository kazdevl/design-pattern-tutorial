package main

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
)

// component
type logger interface {
	write(input []byte) error
}

// concret component
type stdLogger struct {
	output io.Writer
}

func newStdLogger(output io.Writer) *stdLogger {
	return &stdLogger{output: output}
}

func (l *stdLogger) write(b []byte) error {
	_, err := l.output.Write(b)
	return err
}

// decorator1
type maskedLogger struct {
	l logger
}

func newMaskedLogger(l logger) *maskedLogger {
	return &maskedLogger{l: l}
}

func (l *maskedLogger) write(input []byte) error {
	var info userInfo
	if err := json.Unmarshal(input, &info); err != nil {
		l.l.write(input)
		l.l.write([]byte(fmt.Sprintf("\n想定しない形式のデータが設定されました: %v\n", err)))
		return nil
	}
	info.Email = "********"
	b, _ := json.Marshal(info)
	return l.l.write(b)
}

// decorator2
type compressLogger struct {
	l logger
}

func newCompressLogger(l logger) *compressLogger {
	return &compressLogger{l: l}
}

func (l *compressLogger) write(input []byte) error {
	fmt.Println("圧縮前のサイズ:", len(input))
	buf := new(bytes.Buffer)
	zw := zlib.NewWriter(buf)
	defer zw.Close()
	if _, err := zw.Write(input); err != nil {
		return err
	}
	compressBytes := buf.Bytes()
	fmt.Println("圧縮後のサイズ:", len(compressBytes))
	l.l.write(compressBytes)
	return nil
}
