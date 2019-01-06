package parser

import (
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"reflect"
)

type Name struct {
	NameValue *kmipbin.KmipTextString `kmip:"420055"`
	NameType  *kmipbin.KmipEnum       `kmip:"420054"`
}

type Template struct {
	Attribute []*Attribute `kmip:"420008"`
}

type TemplateAttribute struct {
	Name      *Name        `kmip:"420053"`
	Attribute []*Attribute `kmip:"420008"`
}

type PublicKeyTemplateAttribute struct {
	Name      *Name        `kmip:"420053"`
	Attribute []*Attribute `kmip:"420008"`
}

type PrivateKeyTemplateAttribute struct {
	Name      *Name        `kmip:"420053"`
	Attribute []*Attribute `kmip:"420008"`
}


type DataAttributesKmip struct {
	NameTag string
	NameType uint32
	NameValue kmipbin.KmipTextString
	ValueTag string
	ValueType uint32
	ValueValue interface{}
}


type Attribute struct {
	Data []byte
}

func MakeType(val uint32) interface{} {

	switch val {
	case 2:
		return new(kmipbin.KmipInt)
	case 3:
		return new(kmipbin.KmipLongInt)
	case 4:
		return new(kmipbin.KmipBigInt)
	case 5:
		return new(kmipbin.KmipEnum)
	case 6:
		return new(kmipbin.KmipBoolean)
	case 7:
		return new(kmipbin.KmipTextString)
	case 8:
		return new(kmipbin.KmipByteString)
	case 9:
		return new(kmipbin.KmipDate)
	case 10:
		return new(kmipbin.KmipInterval)
	}

	return nil
}



func (a *DataAttributesKmip) Unpack(b []byte) {
	b = b[8:]
	for {
		if len(b) <= 0 {
			break
		}
		tty := kmipbin.KmipTagType{}
		tty.UnMarshalBin(b[:4])
		a.NameTag = fmt.Sprintf("%06X",tty.Tag)
		a.NameType = tty.Type
		var l kmipbin.KmipLength
		l.UnMarshalBin(b[4:8])
		le := kmipbin.PadLength(int(l))
		b = b[8:]
		a.NameValue.UnMarshalBin(b[:le],int(l))
		fmt.Println(a.NameTag , a.NameType , a.NameValue)
		b = b[le:]

		/////

		tty.UnMarshalBin(b[:4])
		a.ValueTag = fmt.Sprintf("%06X",tty.Tag)
		a.ValueType = tty.Type
		l.UnMarshalBin(b[4:8])
		le = kmipbin.PadLength(int(l))
		b = b[8:]


		switch a.ValueType {
		case 2,3,5,6,9,10:
			a.ValueValue = MakeType(a.ValueType).(kmipbin.BaseMarshal)
			a.ValueValue.(kmipbin.BaseMarshal).UnMarshalBin(b[:le])
			b = b[le:]
			fmt.Println(a.ValueTag , a.ValueType, reflect.ValueOf(a.ValueValue).Elem() )
		case 4,7,8:
			a.ValueValue = MakeType(a.ValueType).(kmipbin.BaseMarshalString)
			a.ValueValue.(kmipbin.BaseMarshalString).UnMarshalBin(b[:le] , int(l))
			b = b[le:]
			fmt.Println(a.ValueTag , a.ValueType, reflect.ValueOf(a.ValueValue).Elem() )
		default:
			fmt.Println(a.ValueTag , a.ValueType, string(b[:le]) )
		}

		b = b[len(b):]
	}
}

func (a *Attribute) Unmarshal(bet *[]byte) {
	var l kmipbin.KmipLength
	l.UnMarshalBin((*bet)[4:8])
	le := kmipbin.PadLength(int(l))
	////////////////////////////////////////////

	a.Data = make([]byte, 8+le)
	copy(a.Data, (*bet)[:8+le])

	/////////////////////////////////////////////
	*bet = (*bet)[8+le:]

}

func (a *Attribute) Marshal() []byte {
	size := len(a.Data)
	tmp := make([]byte, size)
	// remove extra copy id not needed
	copy(tmp, a.Data)
	return tmp
}
