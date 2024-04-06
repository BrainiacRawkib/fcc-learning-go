package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// Write Method
	fmt.Printf("\n")
	var w Writer = ConsoleWriter{}
	write, err := w.Write([]byte("Hello GO!"))
	if err != nil {
		return
	}
	fmt.Println(write)

	// Increment Method
	fmt.Printf("\n")
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// NewBufferedWriterCloser Method
	fmt.Printf("\n")
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube Listeners, this is a test"))
	wc.Close()

	// type conversion
	fmt.Printf("type conversion")
	//bwc := wc.(*BufferedWriterCloser)
	//bwc := wc.(io.Reader)
	//fmt.Println(bwc)
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed!")
	}

	// Empty Interface
	fmt.Printf("\n")
	fmt.Printf("Empty Interface")
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello Youtube Listeners, this is a test"))
		wc.Close()
	}
	emptyR, emptyOk := myObj.(io.Reader)
	if emptyOk {
		fmt.Println(emptyR)
	} else {
		fmt.Println("Conversion failed!!")
	}

	// Switching Interface
	fmt.Printf("\n")
	switchInterface()

	var wc2 WriterCloser = &myWriterCloser{}
	fmt.Println(wc2)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Printf("*** Write Method ***\n")
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	fmt.Printf("*** Increment Method ***\n")
	*ic++
	return int(*ic)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mwc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	fmt.Printf("*** NewBufferedWriterCloser Method ***\n")
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func switchInterface() {
	var i interface{} = true
	switch i.(type) {
	case int:
		fmt.Println("i is an integer!")
	case string:
		fmt.Println("i is a string!")
	case float32:
		fmt.Println("i is float32!")
	default:
		fmt.Println("i is unknown!")
	}
}
