package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"reflect"
	"strings"
	"github.com/rajnikant12345/kmip_g/objects"
)

func GetTypeString(value reflect.Value) string {
	switch value.Type().String() {
	case "*kmipbin.KmipInt":
		return "02"
	case "*kmipbin.KmipBoolean":
		return "06"
	case "*kmipbin.KmipEnum":
		return "05"
	case "*kmipbin.KmipDate":
		return "09"
	case "*kmipbin.KmipLongInt":
		return "03"
	case "*kmipbin.KmipInterval":
		return "0A"
	case "*kmipbin.KmipTextString":
		return "07"
	case "*kmipbin.KmipByteString":
		return "08"
	default:
		return "-1"
	}
}

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

func GetKmipInt(value reflect.Value) interface{} {
	switch value.Type().String() {
	case "*kmipbin.KmipInt":
		return new(kmipbin.KmipInt)
	case "*kmipbin.KmipBoolean":
		return new(kmipbin.KmipBoolean)
	case "*kmipbin.KmipEnum":
		return new(kmipbin.KmipEnum)
	case "*kmipbin.KmipDate":
		return new(kmipbin.KmipDate)
	case "*kmipbin.KmipLongInt":
		return new(kmipbin.KmipLongInt)
	case "*kmipbin.KmipInterval":
		return new(kmipbin.KmipInterval)
	}

	return nil
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


func GetKMIPString(value reflect.Value) interface{} {
	switch value.Type().String() {
	case "*kmipbin.KmipTextString":
		return new(kmipbin.KmipTextString)
	case "*kmipbin.KmipByteString":
		return new(kmipbin.KmipByteString)
	}

	return nil
}

func UnmaeshalAllT(a *KmipStruct, b []byte) {
	v := reflect.ValueOf(a)

	typ := reflect.TypeOf(v.Elem().Field(0).Interface()).Elem()

	k := reflect.New(typ)

	b = b[8:]
	Dummy(&k, &b)

	v.Elem().Field(0).Set(k)

}

func ReadKmipInt(f *reflect.Value, bet *[]byte) {
	*bet = (*bet)[8:]
	typ := reflect.TypeOf(f.Interface()).Elem()
	k := reflect.New(typ)
	k.MethodByName("UnMarshalBin").Call([]reflect.Value{reflect.ValueOf((*bet)[:8])})
	f.Set(k)
	(*bet) = (*bet)[8:]
}

func ReadKmipString(f *reflect.Value, bet *[]byte) {
	var l kmipbin.KmipLength
	l.UnMarshalBin((*bet)[4:8])
	*bet = (*bet)[8:]
	le := kmipbin.PadLength(int(l))
	typ := reflect.TypeOf(f.Interface()).Elem()
	k := reflect.New(typ)
	k.MethodByName("UnMarshalBin").Call([]reflect.Value{reflect.ValueOf((*bet)[:8]), reflect.ValueOf(int(l))})
	*bet = (*bet)[le:]
	f.Set(k)
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

		field := v.Elem().Field(i)

		tag := Tags[ty.Field(i).Name]

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
			if tag == "420040" {
				a := &objects.KeyBlock{}
				a.Unmarshal(bet)
				f.Set(reflect.ValueOf(a))
			}else if tag == "420008" {
				a := &objects.Attribute{}
				a.Unmarshal(bet)
				f.Set(reflect.Append(f, reflect.ValueOf(a)))
			} else if IsKmipInt(f) {
				ReadKmipInt(&f, bet)
			} else if IsKMIPString(f) {
				ReadKmipString(&f, bet)
			} else {

				typ := reflect.TypeOf(f.Interface())
				//(*bet) = (*bet)[8:]
				if typ.Kind() == reflect.Ptr {
					typ = reflect.TypeOf(f.Interface()).Elem()
				}
				var k reflect.Value
				if typ.Kind() == reflect.Slice {
					tt := typ.Elem()
					k = reflect.New(tt.Elem())
					if k.Elem().Kind() != reflect.Struct {
						if IsKmipInt(k) {
							p := GetKmipInt(k)
							h := reflect.ValueOf(&p).Elem()
							ReadKmipInt(&h, bet)
							k = reflect.ValueOf(h.Interface())
						} else if IsKMIPString(k) {
							//fmt.Println(k.Type(), k.Elem().Type())
							p := GetKMIPString(k)
							h := reflect.ValueOf(&p).Elem()
							ReadKmipString(&h, bet)
							k = reflect.ValueOf(h.Interface())
						}
					} else {
						(*bet) = (*bet)[8:]
						Dummy(&k, bet)
					}
				} else {
					(*bet) = (*bet)[8:]
					k = reflect.New(typ)
					Dummy(&k, bet)
				}
				if typ.Kind() == reflect.Slice {
					f.Set(reflect.Append(f, k))
				} else {
					f.Set(k)
				}
			}
		}

	}
}
/*
func MarshalAllT(a *KmipStructResponse) []byte {
	//ty := reflect.TypeOf(*a)
	v := reflect.ValueOf(a)

	return DummyMarshal(&v, "")

}*/

func MarshalAllRequest(a *KmipStruct) []byte {
	//ty := reflect.TypeOf(*a)
	v := reflect.ValueOf(a)

	return DummyMarshal(&v, "")

}

func WriteKmipInt(field reflect.Value, tag string, b *bytes.Buffer) {
	byt := field.MethodByName("MarshalBin").Call(nil)
	l := fmt.Sprintf("%08x", 4)
	k, _ := hex.DecodeString(tag + GetTypeString(field) + l)
	b.Write(k)
	b.Write(byt[0].Bytes())
}

func WriteKmipString(field reflect.Value, tag string, b *bytes.Buffer) {
	//fmt.Println(field.Elem())
	length := field.Elem().Len()
	byt := field.MethodByName("MarshalBin").Call(nil)
	l := fmt.Sprintf("%08x", length)
	k, _ := hex.DecodeString(tag + GetTypeString(field) + l)
	b.Write(k)
	b.Write(byt[0].Bytes())
}

func DummyMarshal(v *reflect.Value, tagin string) []byte {
	b := bytes.Buffer{}
	ty := reflect.TypeOf(v.Elem().Interface())

	for i := 0; i < v.Elem().NumField(); i++ {
		field := v.Elem().Field(i)

		if !field.IsNil() {

			if IsKmipInt(field) {
				var tag string
				if tagin != "" {
					tag = tagin
				} else {
					tag = Tags[ty.Field(i).Name]
				}
				if tag == "" {
					continue
				}
				WriteKmipInt(field, tag, &b)

			} else if IsKMIPString(field) {
				var tag string
				if tagin != "" {
					tag = tagin
				} else {
					tag = Tags[ty.Field(i).Name]
				}
				if tag == "" {
					continue
				}

				WriteKmipString(field, tag, &b)

			} else {
				tag := Tags[ty.Field(i).Name]
				if tag == "420008" {
					length := field.Len()
					for i := 0; i < length; i++ {
						el := field.Index(i).Interface().(*Attribute)
						b.Write(el.Marshal())
					}
				} else {
					var bt []byte
					if field.Kind() == reflect.Slice {
						length := field.Len()
						for i := 0; i < length; i++ {
							el := field.Index(i)
							if IsKMIPString(el) {
								WriteKmipString(el, tag, &b)
							} else if IsKmipInt(el) {
								WriteKmipInt(el, tag, &b)
							} else {
								bty := DummyMarshal(&el, "")
								l := fmt.Sprintf("%08x", len(bty))
								k, _ := hex.DecodeString(tag + "01" + l)
								bt = append(bt, k...)
								bt = append(bt, bty...)
							}
						}
						b.Write(bt)
					} else {
						bt = DummyMarshal(&field, "")
						l := fmt.Sprintf("%08x", len(bt))
						k, _ := hex.DecodeString(tag + "01" + l)
						b.Write(k)
						b.Write(bt)
					}
				}
			}
		}
	}
	return b.Bytes()
}
