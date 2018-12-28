package parser

import (
	"bytes"
	"io"
	"github.com/rajnikant12345/kmip_g/kmipbin"
)


func TTLVReader(reader io.ReadWriter) ([]byte, error) {
	out  :=  bytes.Buffer{}

	initialBytes := make([]byte,8)

	tmp := initialBytes

	var expected kmipbin.KmipLength
	// we expect 8 bytes initially
	expected = 8
	for {
		// read first 8 bytes
		b,err := reader.Read(tmp)
		if err != nil {
			return nil, err
		}
		// if 8 bytes are not recieved, just try tp read more
		if b < int(expected) {
			tmp = tmp[b:]
		}else {
			break
		}
		expected = expected - kmipbin.KmipLength(b)
	}

	out.Write(initialBytes)
	// read expected bytes from KMIP buffer
	expected.UnMarshalBin(initialBytes[4:])

	// create a buffer to contain full kmip bytes
	fullBytes := make([]byte,expected)

	// use temp buffer to hold partial data
	tmp = fullBytes


	for {
		b,err := reader.Read(tmp)
		if err != nil {
			return nil, err
		}
		if b < int(expected) {
			tmp = tmp[b:]
		}else {
			break
		}
		expected = expected - kmipbin.KmipLength(b)
	}
	out.Write(fullBytes)

	return out.Bytes(), nil
}



