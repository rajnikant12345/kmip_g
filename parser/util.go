package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"reflect"
	"strings"
)



func IsCompound(value reflect.Value) bool {
	switch value.Type().String() {
	case "kmipbin.KmipInt":
		return false
	case "kmipbin.KmipBoolean":
		return false
	case "kmipbin.KmipEnum":
		return false
	case "kmipbin.KmipDate":
		return false
	case "kmipbin.KmipLongInt":
		return false
	case "kmipbin.KmipInterval":
		return false
	case "kmipbin.KmipTextString":
		return false
	case "kmipbin.KmipByteString":
		return false
	}

	return true
}

func forwadBytes(b []byte) []byte {
	return nil
}


func UnmaeshalAllT(a *KmipStruct, b []byte) {
	//ty := reflect.TypeOf(*a)
	v := reflect.ValueOf(a)

	//fmt.Println(reflect.TypeOf(v.Elem().Field(0).Interface()))
	typ := reflect.TypeOf(v.Elem().Field(0).Interface()).Elem()

	k := reflect.New(typ)
	//	fmt.Println(k.Type())

	b = b[8:]
	Dummy(&k, &b)

	v.Elem().Field(0).Set(k)

}

func Dummy(v *reflect.Value, bet *[]byte) {

	//b := *bet

	tagmap := make(map[string]reflect.Value)

	tagmapmulti := make(map[string]bool)

	ty := reflect.TypeOf(v.Elem().Interface())

	//fmt.Println("lallu",ty.String() , ty.Kind().String())

	for i := 0; i < v.Elem().NumField(); i++ {
		field := v.Elem().Field(i)
		tag := ty.Field(i).Tag.Get("kmip")

		//fmt.Println(ty.Field(i) , field)

		if tag != "" {
			tagmap[tag] = field
			tagm := ty.Field(i).Tag.Get("multi")
			if tagm == "true" {
				tagmapmulti[tag] = true
			} else {
				tagmapmulti[tag] = false
			}
		}
	}

	for {
		if len((*bet)) < 8 {
			break
		}
		tag := strings.ToUpper(hex.EncodeToString((*bet)[:3]))
		//fmt.Println(tag)
		f, ok := tagmap[tag]
		if !ok {
			return
			//b = forwadBytes(b)
		} else {

			// handle attribute saperatr as it comes with many flavours
			if tag == "420008" {

				a := &Attribute{}
				a.Unmarshal(bet)
				f.Set(reflect.Append(f, reflect.ValueOf(a)))
				return
			}

			//fmt.Println("rajni",f.Type().String())
			switch f.Type().String() {
			case "*kmipbin.KmipInt":
				*bet = (*bet)[8:]
				var k kmipbin.KmipInt
				k.UnMarshalBin((*bet)[:8])
				f.Set(reflect.ValueOf(&k))
				(*bet) = (*bet)[8:]
			case "*kmipbin.KmipEnum":
				*bet = (*bet)[8:]
				var k kmipbin.KmipEnum
				k.UnMarshalBin((*bet)[:8])
				f.Set(reflect.ValueOf(&k))
				(*bet) = (*bet)[8:]
			case "*kmipbin.KmipByteString":
				var l kmipbin.KmipLength
				l.UnMarshalBin((*bet)[4:8])
				*bet = (*bet)[8:]
				le := kmipbin.PadLength(int(l))
				var uvi kmipbin.KmipByteString
				uvi.UnMarshalBin((*bet)[:le], int(l))
				*bet = (*bet)[le:]
				f.Set(reflect.ValueOf(&uvi))
			case "*kmipbin.KmipTextString":
				var l kmipbin.KmipLength
				l.UnMarshalBin((*bet)[4:8])
				*bet = (*bet)[8:]
				le := kmipbin.PadLength(int(l))
				var uvi kmipbin.KmipTextString
				uvi.UnMarshalBin((*bet)[:le], int(l))
				*bet = (*bet)[le:]
				f.Set(reflect.ValueOf(&uvi))
			case "*kmipbin.KmipBoolean":
				*bet = (*bet)[8:]
				var k kmipbin.KmipBoolean
				k.UnMarshalBin((*bet)[:8])
				f.Set(reflect.ValueOf(&k))
				(*bet) = (*bet)[8:]
			case "*kmipbin.KmipDate":
				*bet = (*bet)[8:]
				var k kmipbin.KmipDate
				k.UnMarshalBin((*bet)[:8])
				f.Set(reflect.ValueOf(&k))
				(*bet) = (*bet)[8:]
			case "*kmipbin.KmipLongInt":
				*bet = (*bet)[8:]
				var k kmipbin.KmipLongInt
				k.UnMarshalBin((*bet)[:8])
				f.Set(reflect.ValueOf(&k))
				(*bet) = (*bet)[8:]
			case "*kmipbin.KmipInterval":
				*bet = (*bet)[8:]
				var k kmipbin.KmipInterval
				k.UnMarshalBin((*bet)[:8])
				f.Set(reflect.ValueOf(&k))
				(*bet) = (*bet)[8:]
			default:

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
				fmt.Println(l)
				k,_ := hex.DecodeString(tag+"02"+l)
				b.Write(k)
				b.Write(p.MarshalBin())
			default:
				tag := ty.Field(i).Tag.Get("kmip")
				if tag == "420008" {

				}else {
					bt := DummyMarshal(&field)
					l := fmt.Sprintf("%08x",len(bt))
					fmt.Println(l)
					k,_ := hex.DecodeString(tag+"01"+l)
					b.Write(k)
					b.Write(bt)
				}

			}
		}
	}
	return b.Bytes()
}
