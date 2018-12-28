package kmipbin

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"math/big"
)

type KmipInt uint32

func (k *KmipInt) MarshalBin() []byte {
	i := *k
	b := make([]byte,8)
	binary.BigEndian.PutUint32(b[:4],uint32(i))
	return b
}


func (k *KmipInt) UnMarshalBin( b []byte  ) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipInt( binary.BigEndian.Uint32(b[:4]) )
	return nil
}

type KmipLongInt uint64


func (k *KmipLongInt) MarshalBin() []byte {
	i := *k
	b := make([]byte,8)
	binary.BigEndian.PutUint64(b,uint64(i))
	return b
}


func (k *KmipLongInt) UnMarshalBin( b []byte  ) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipLongInt( binary.BigEndian.Uint64(b) )
	return nil
}



type KmipTextString string


func (k *KmipTextString) MarshalBin() []byte {
	i := *k
	l := (len(i)+8) - (len(i)+8)%8
	b := make([]byte, l)
	copy(b, []byte(*k))
	return b
}


func (k *KmipTextString) UnMarshalBin( b []byte , length int  ) error {
	if len(b)%8 != 0 {
		return errors.New("Invalid Kmip string, must be multiple of 8")
	}
	*k = KmipTextString(b[:length])
	return nil
}


type KmipByteString []byte


func (k *KmipByteString) MarshalBin() []byte {
	i := *k
	l := (len(i)+8) - (len(i)+8)%8
	b := make([]byte, l)
	copy(b, []byte(*k))
	return b
}


func (k *KmipByteString) UnMarshalBin( b []byte , length int  ) error {
	if len(b)%8 != 0 {
		return errors.New("Invalid Kmip byte string, must be multiple of 8")
	}
	*k = KmipByteString(b[:length])
	return nil
}


type KmipBigInt big.Int

func (k *KmipBigInt) MarshalBin() []byte {
	i := big.Int(*k)
	b := i.Bytes()
	pad := make([]byte, len(b)%8)
	b = append(pad, b...)
	return b
}


func (k *KmipBigInt) UnMarshalBin( b []byte , length int  ) error {
	if len(b)%8 != 0 {
		return errors.New("Invalid Kmip byte string, must be multiple of 8")
	}
	i := new(big.Int)
	i.SetBytes(b[:length])
	*k = KmipBigInt(*i)

	return nil
}