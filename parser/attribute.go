package parser

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
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

type Attribute struct {
	Data []byte
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
