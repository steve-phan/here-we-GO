package main

import (
	"bytes"
	"fmt"
)

const title = "Bookable24 a booking system for Restaurants and Nail salons"

const byteNum = 7

func main() {
	var bp = &BufferProgress{
		buffer: bytes.NewBuffer([]byte{}),
	}
	fmt.Printf("The buffer start is %v", bp.buffer)
	// bp.buffer  = "Bookabl" because go will move all declare var to top

	num, _ := bp.Print([]byte(title))
	bp.Close()
	fmt.Printf("The result is %v", num)

}

type BufferProgress struct {
	buffer *bytes.Buffer
}

func (bp *BufferProgress) Print(data []byte) (int, error) {
	// create an empty slice
	v := make([]byte, byteNum)
	n, err := bp.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	for bp.buffer.Len() > byteNum {
		// Read the next 8 bytes and assign to v
		_, err := bp.buffer.Read(v)

		if err != nil {
			return 0, err
		}
		// print it out ( do something with this 8 bytes)
		fmt.Println(string(v))
	}

	return n, nil
}

func (bp *BufferProgress) Close() error {
	for bp.buffer.Len() > 0 {
		data := bp.buffer.Next(byteNum)
		_, err := fmt.Println(string(data))

		if err != nil {
			return err
		}
	}

	return nil
}
