package main

import (
	"fmt"
	"bytes"
	"io"
)

func main(){
	newBufferedWriterCloser := &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
	var wc WriterCloser = newBufferedWriterCloser
	wc.Write([]byte("Hello Golang interfaces, This is a test"))
	wc.Close()
	// you can do type conversions like this (handle failures to avoid panics)
	reader,ok := wc.(io.Reader)
	if ok {
		fmt.Println(reader)
	} else {
		fmt.Println("Type conversion failed")
	}
}	

type Writer interface {
	Write([]byte)(int,error)
}

type Closer interface {
	Close() error
}

// This is composed of Writer & Closer interface
// This is done exactly how you would do embeding with structs
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte)(int,error) {
	n,err := bwc.buffer.Write(data)
	if err != nil {
		return 0,err
	}
	v := make([]byte,8)
	for bwc.buffer.Len() > 8 {
		_,err := bwc.buffer.Read(v)
		if err != nil{
			return 0,err
		}
		fmt.Println(string(v)) 
	}
	return n,nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_,err := fmt.Println(string(data))
		if err != nil{
			return err
		}
	}
	return nil
}