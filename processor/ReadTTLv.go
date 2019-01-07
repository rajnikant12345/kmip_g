package processor

import (
	"io"
	"errors"
)

func ReadBytes(len int , b io.ReadWriter) ([]byte, error) {
	out := make([]byte,len)
	n, e :=  b.Read(out)

	if e != nil {
		return nil , e
	}

	if n != len {
		return nil,  errors.New("ReadBytes insuggicient bytes")
	}

	return out ,  nil

}


