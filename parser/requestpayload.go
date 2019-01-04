package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type RequestPayload struct {
	ObjectType                  *kmipbin.KmipEnum            `kmip:"420057"`
	UniqueIdentifier            *kmipbin.KmipTextString      `kmip:"420094"`
	TemplateAttribute           *TemplateAttribute           `kmip:"420091"`
	Attribute                   []*Attribute                 `kmip:"420008"`
	PublicKeyTemplateAttribute  *PublicKeyTemplateAttribute  `kmip:"42006E"`
	PrivateKeyTemplateAttribute *PrivateKeyTemplateAttribute `kmip:"420065"`
	Template                    *Template                    `kmip:"420090"`
}

func (p *RequestPayload) GetObjectType() *kmipbin.KmipEnum {
	return p.ObjectType
}

func (p *RequestPayload) GetUniqueIdentifier() *kmipbin.KmipTextString {
	return p.UniqueIdentifier
}

func (p *RequestPayload) GetAttributeCount() int {
	return len(p.Attribute)
}

func (p *RequestPayload) GetAttributeAt(index int) *Attribute {
	if p.GetAttributeCount()-1 > index {
		return nil
	}
	return p.Attribute[index]
}
