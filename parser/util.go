package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"reflect"
	"strings"
)



func IsKmipInt(value reflect.Value) bool {
	switch value.Type().String() {
	case "*kmipbin.KmipInt":
		return true
	case "*kmipbin.KmipBoolean":
		return true
	case "*kmipbin.KmipEnum":
		return true
	case "*kmipbin.KmipDate":
		return true
	case "*kmipbin.KmipLongInt":
		return true
	case "*kmipbin.KmipInterval":
		return true
	}

	return false
}

func IsKMIPString(value reflect.Value) bool {
	switch value.Type().String() {
	case "*kmipbin.KmipTextString":
		return true
	case "*kmipbin.KmipByteString":
		return true
	}

	return false
}


func UnmaeshalAllT(a *KmipStruct, b []byte) {
	v := reflect.ValueOf(a)

	typ := reflect.TypeOf(v.Elem().Field(0).Interface()).Elem()

	k := reflect.New(typ)

	b = b[8:]
	Dummy(&k, &b)

	v.Elem().Field(0).Set(k)

}

func Dummy(v *reflect.Value, bet *[]byte) {


	// this map consisit of mapping between kmip tag and
	// and a reflect.value
	tagmap := make(map[string]reflect.Value)


	// get the type of element inside v
	ty := reflect.TypeOf(v.Elem().Interface())


	// it must be a struct type, hence all the elements must be iterated and
	// must be checked against a valid kmip tag
	for i := 0; i < v.Elem().NumField(); i++ {
		// het the field at index
		field := v.Elem().Field(i)

		// get the kmip tag value
		tag := ty.Field(i).Tag.Get("kmip")

		// if it's a valid kmip element then , add it to tagmap
		if tag != "" {
			tagmap[tag] = field
		}
	}

	// this loop is infinite as we will be iterating till we reach end of input buffer
	// or an unwanted element is encountered.
	for {

		// check the length of input buffer
		if len((*bet)) < 8 {
			break
		}

		// get the tag from buffer
		tag := strings.ToUpper(hex.EncodeToString((*bet)[:3]))

		// find the tag in tagmap
		f, ok := tagmap[tag]
		if !ok {
			return
		} else {

			// handle attribute separate as it comes with many flavours
			if tag == "420008" {

				a := &Attribute{}
				a.Unmarshal(bet)
				f.Set(reflect.Append(f, reflect.ValueOf(a)))
				return
			}

			if IsKmipInt(f) {
				*bet = (*bet)[8:]
				typ := reflect.TypeOf(f.Interface()).Elem()
				k := reflect.New(typ)
				k.MethodByName("UnMarshalBin").Call([]reflect.Value{reflect.ValueOf((*bet)[:8])})
				f.Set(k)
				(*bet) = (*bet)[8:]
			} else if IsKMIPString(f) {
				var l kmipbin.KmipLength
				l.UnMarshalBin((*bet)[4:8])
				*bet = (*bet)[8:]
				le := kmipbin.PadLength(int(l))
				typ := reflect.TypeOf(f.Interface()).Elem()
				k := reflect.New(typ)
				k.MethodByName("UnMarshalBin").Call([]reflect.Value{reflect.ValueOf((*bet)[:8]),reflect.ValueOf(int(l))})
				*bet = (*bet)[le:]
				f.Set(k)
			} else {
				typ := reflect.TypeOf(f.Interface())
				(*bet) = (*bet)[8:]

				if typ.Kind() == reflect.Ptr {
					typ = reflect.TypeOf(f.Interface()).Elem()
				}

				//	fmt.Println("default",typ , typ.Kind())

				var k reflect.Value

				if typ.Kind() == reflect.Slice {
					tt := typ.Elem()
					k = reflect.New(tt.Elem())
					//fmt.Println(k.Type().String())
				} else {
					k = reflect.New(typ)
				}

				//	fmt.Println("default 1",k.Type().String() )

				Dummy(&k, bet)

				if typ.Kind() == reflect.Slice {
					//	fmt.Println(k)
					f.Set(reflect.Append(f, k))
					//	fmt.Println(f)
				} else {
					f.Set(k)
				}
			}
		}

	}
}



func MarshalAllT(a *KmipStructResponse) []byte {
	//ty := reflect.TypeOf(*a)
	v := reflect.ValueOf(a)

	return DummyMarshal(&v)


}


func DummyMarshal(v *reflect.Value) []byte {

	b := bytes.Buffer{}

	ty := reflect.TypeOf(v.Elem().Interface())

	for i := 0; i < v.Elem().NumField(); i++ {
		field := v.Elem().Field(i)

		if !field.IsNil() {
			switch field.Type().String() {
			case "*kmipbin.KmipInt":
				tag := ty.Field(i).Tag.Get("kmip")
				if tag == "" {
					continue
				}
				p := field.Interface().(*kmipbin.KmipInt)
				l := fmt.Sprintf("%08x",4)
				k,_ := hex.DecodeString(tag+"02"+l)
				b.Write(k)
				b.Write(p.MarshalBin())
			default:
				tag := ty.Field(i).Tag.Get("kmip")
				if tag == "420008" {

				}else {
					bt := DummyMarshal(&field)
					l := fmt.Sprintf("%08x",len(bt))
					k,_ := hex.DecodeString(tag+"01"+l)
					b.Write(k)
					b.Write(bt)
				}

			}
		}
	}
	return b.Bytes()
}
