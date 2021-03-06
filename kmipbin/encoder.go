package kmipbin

import (
	"encoding/binary"
	"errors"
)

func PadLength(in int) int {
	if in%8 == 0 {
		return in
	}
	return in + 8 - ((in + 8) % 8)
}

type KmipTagType struct {
	Tag  uint32
	Type uint32
}


type BaseMarshal interface {
	UnMarshalBin([]byte) error
	MarshalBin() []byte
}

type BaseMarshalString interface {
	UnMarshalBin([]byte , int) error
	MarshalBin() []byte
}

func (k *KmipTagType) UnMarshalBin(b []byte) error {
	if len(b) != 4 {
		return errors.New("Invalid Kmip Tag")
	}
	k.Tag = binary.BigEndian.Uint32(b)
	k.Type = k.Tag & 0x000000ff
	k.Tag = k.Tag >> 8
	return nil
}

type KmipLength uint32

func (k *KmipLength) MarshalBin() []byte {
	i := *k
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b[:4], uint32(i))
	return b
}

func (k *KmipLength) UnMarshalBin(b []byte) error {
	if len(b) != 4 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipLength(binary.BigEndian.Uint32(b))
	return nil
}

type KmipInt uint32

func (k *KmipInt) MarshalBin() []byte {
	i := *k
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b[:4], uint32(i))
	return b
}

func (k *KmipInt) UnMarshalBin(b []byte) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipInt(binary.BigEndian.Uint32(b[:4]))
	return nil
}


type KmipEnum uint32

func (k *KmipEnum) MarshalBin() []byte {
	i := *k
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b[:4], uint32(i))
	return b
}

func (k *KmipEnum) UnMarshalBin(b []byte) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipEnum(binary.BigEndian.Uint32(b[:4]))
	return nil
}


type KmipInterval uint32

func (k *KmipInterval) MarshalBin() []byte {
	i := *k
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b[:4], uint32(i))
	return b
}

func (k *KmipInterval) UnMarshalBin(b []byte) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipInterval(binary.BigEndian.Uint32(b[:4]))
	return nil
}



type KmipLongInt uint64

func (k *KmipLongInt) MarshalBin() []byte {
	i := *k
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func (k *KmipLongInt) UnMarshalBin(b []byte) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipLongInt(binary.BigEndian.Uint64(b))
	return nil
}



type KmipBoolean uint64

func (k *KmipBoolean) MarshalBin() []byte {
	i := *k
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func (k *KmipBoolean) UnMarshalBin(b []byte) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipBoolean(binary.BigEndian.Uint64(b))
	return nil
}





type KmipDate uint64

func (k *KmipDate) MarshalBin() []byte {
	i := *k
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func (k *KmipDate) UnMarshalBin(b []byte) error {
	if len(b) != 8 {
		return errors.New("Invalid KmipInt")
	}
	*k = KmipDate(binary.BigEndian.Uint64(b))
	return nil
}





type KmipTextString string

func (k *KmipTextString) MarshalBin() []byte {
	i := *k
	// create a mulitple of 8 padded with zero
	l := PadLength(len(i))
	b := make([]byte, l)
	copy(b, []byte(*k))
	return b
}

func (k *KmipTextString) UnMarshalBin(b []byte, length int) error {
	if len(b)%8 != 0 {
		return errors.New("Invalid Kmip string, must be multiple of 8")
	}
	*k = KmipTextString(b[:length])
	return nil
}



type KmipByteString []byte

func (k *KmipByteString) MarshalBin() []byte {
	i := *k
	l := PadLength(len(i))
	b := make([]byte, l)
	copy(b, []byte(*k))
	return b
}

func (k *KmipByteString) UnMarshalBin(b []byte, length int) error {
	if len(b)%8 != 0 {
		return errors.New("Invalid Kmip byte string, must be multiple of 8")
	}
	*k = KmipByteString(b[:length])
	return nil
}



type KmipBigInt []byte


func (k *KmipBigInt) MarshalBin() []byte {
	i := *k
	l := PadLength(len(i))
	b := make([]byte, l)
	copy(b, []byte(*k))
	return b
}

func (k *KmipBigInt) UnMarshalBin(b []byte, length int) error {
	if len(b)%8 != 0 {
		return errors.New("Invalid Kmip byte string, must be multiple of 8")
	}
	*k = KmipBigInt(b[:length])
	return nil
}

type KmipStruct interface {
	MarsalBin() []byte
	UnMarshalBin() error
}


