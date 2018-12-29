package kmipbin

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
)

func TestKmipInt_MarshalBin(t *testing.T) {
	var k KmipInt

	k = 0x20

	b := k.MarshalBin()
	if string(b) != string([]byte{0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00}) {
		t.Fail()
	}
}

func TestKmipInt_UnMarshalBin(t *testing.T) {
	b := []byte{0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00}
	var k KmipInt
	k.UnMarshalBin(b)
	if k != 32 {
		t.Fail()
	}
}

func TestKmipLongInt_MarshalBin(t *testing.T) {
	var k KmipLongInt

	k = 0x20

	b := k.MarshalBin()
	if string(b) != string([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20}) {
		t.Fail()
	}
}

func TestKmipLongInt_UnMarshalBin(t *testing.T) {
	b := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20}
	var k KmipLongInt
	k.UnMarshalBin(b)
	if k != 32 {
		t.Fail()
	}
}

func TestKmipTextString_UnMarshalBin(t *testing.T) {
	b, _ := hex.DecodeString("48656C6C6F20576F726C640000000000")
	var k KmipTextString
	k.UnMarshalBin(b, 11)
	if k != "Hello World" {
		t.Fail()
	}
	if len(k) != 11 {
		t.Fail()
	}
}

func TestKmipTextString_MarshalBin(t *testing.T) {

	k := KmipTextString("Hello World")

	b := k.MarshalBin()

	t.Log(hex.EncodeToString(b))

	if strings.ToUpper(hex.EncodeToString(b)) != "48656C6C6F20576F726C640000000000" {
		t.Fail()
	}

}

func TestKmipByteString_UnMarshalBin(t *testing.T) {
	b, _ := hex.DecodeString("48656C6C6F20576F726C640000000000")
	k := KmipByteString(b)
	k.UnMarshalBin(b, 11)
	if string(k) != string(b[:11]) {
		t.Fail()
	}
	if len(k) != 11 {
		t.Fail()
	}
}

func TestKmipByteString_MarshalBin(t *testing.T) {

	by, _ := hex.DecodeString("48656C6C6F20576F726C64")

	k := KmipByteString(by)

	b := k.MarshalBin()

	t.Log(hex.EncodeToString(b))

	if strings.ToUpper(hex.EncodeToString(b)) != "48656C6C6F20576F726C640000000000" {
		t.Fail()
	}

}

func TestKmipBigInt_UnMarshalBin(t *testing.T) {
	b, _ := hex.DecodeString("0000000003FD35EB6BC2DF4618080000")

	var k KmipBigInt

	k.UnMarshalBin(b, 16)

	h := big.Int(k)
	if h.String() != "1234567890000000000000000000" {
		t.Fail()
	}

}

func TestKmipBigInt_MarshalBin(t *testing.T) {
	b, _ := hex.DecodeString("0000000003FD35EB6BC2DF4618080000")

	var k KmipBigInt

	k.UnMarshalBin(b, 16)

	b = k.MarshalBin()

	t.Log(strings.ToUpper(hex.EncodeToString(b)))

	if strings.ToUpper(hex.EncodeToString(b)) != "0000000003FD35EB6BC2DF4618080000" {
		t.Fail()
	}

}
