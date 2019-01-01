package parser

import (
	"encoding/hex"
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"strings"
)

type Name struct {
	NameValue string `xml:",omitempty"`
	NameType  int    `xml:",omitempty"` //enum

}

type CryptoParams struct {
	BlockCipherMode     int   `xml:",omitempty"` //enum
	PaddingMethod       int   `xml:",omitempty"` //enum
	HashASlgo           int   `xml:",omitempty"` //enum
	KeyRoleType         int   `xml:",omitempty"` //enum
	DigitalSigAlgo      int   `xml:",omitempty"` //enum
	CryptoAlgo          int   `xml:",omitempty"` //enum
	RandomIv            int64 `xml:",omitempty"` //bool
	IvLen               int   `xml:",omitempty"`
	TagLen              int   `xml:",omitempty"`
	FixedFieldLen       int   `xml:",omitempty"`
	InvocationFieldLen  int   `xml:",omitempty"`
	CounterLength       int   `xml:",omitempty"`
	InitialCounterValue int   `xml:",omitempty"`
}

type CryptoDomainParams struct {
	Qlen             int `xml:",omitempty"`
	RecommendedCurve int `xml:",omitempty"` //enum
}

type X509CertificateId struct {
	IssuerDistinguishedName []byte `xml:",omitempty"`
	CertificateSerialNumber []byte `xml:",omitempty"`
}

type X509CertificateSub struct {
	SubjectDistinguishedName []byte `xml:",omitempty"`
	SubjectAlternativeName   []byte `xml:",omitempty"`
}

type X509CertificateIssuer struct {
	IssuerDistinguishedName []byte `xml:",omitempty"`
	IssuerAlternativeName   []byte `xml:",omitempty"`
}

type CertificateId struct {
	Issuer       string `xml:",omitempty"`
	SerialNumber string `xml:",omitempty"`
}

type CertificateSub struct {
	CertificateSubjectDistinguishedName string `xml:",omitempty"`
	CertificateSubjectAlternativeName   string `xml:",omitempty"`
}

type CertificateIssuer struct {
	CertificateIssuerDistinguishedName string `xml:",omitempty"`
	CertificateIssuerAlternativeName   string `xml:",omitempty"`
}

type Digest struct {
	HashAlgo      int    `xml:",omitempty"` //enum
	DigestValue   []byte `xml:",omitempty"`
	KeyFormatType int    `xml:",omitempty"` //enum
}

type UsageLimit struct {
	UsageLimit      int64 `xml:",omitempty"`
	UsageLimitCount int64 `xml:",omitempty"`
	UsageLimitUnit  int   `xml:",omitempty"` //enum

}

type RevocationReason struct {
	Code    int    `xml:",omitempty"` //enum
	Message string `xml:",omitempty"` //enum
}

type Link struct {
	Type         int    `xml:",omitempty"` //enum
	LinkedObject string `xml:",omitempty"`
}

type ApplicationSpecificInformation struct {
	ApplicationNamespace string `xml:",omitempty"`
	ApplicationData      string `xml:",omitempty"`
}

type AlternativeName struct {
	AlternativeNameValue string `xml:",omitempty"`
	AlternativeNameType  int    `xml:",omitempty"` //enum
}

type KeyValueLocation struct {
	KeyValueLocationValue string `xml:",omitempty"`
	KeyValueLocationType  string `xml:",omitempty"`
}

type CustomAttributeValue struct {
	Type  int
	Value interface{}
}

type AttributeValue struct {
	ApplicationNamespace *kmipbin.KmipTextString
	ApplicationData		 *kmipbin.KmipByteString

}

type TemplateAttribute struct {
	Attribute        []*Attribute            `kmip:"420008"`
}

type Attribute struct {
	AttributeName  *kmipbin.KmipTextString `kmip:"42000A"`
	AttributeValue interface{}          `kmip:"42000B"`
}

func IsEnum(name kmipbin.KmipTextString) bool {
	if name == "Cryptographic Algorithm" {
		return true
	}
	return false
}

func IsInt(name kmipbin.KmipTextString) bool {
	if name == "Cryptographic Length" ||
		name == "Cryptographic Usage Mask" {
		return true
	}
	return false
}


func IsBigInt(name kmipbin.KmipTextString) bool {
	return false
}

func IsLong(name kmipbin.KmipTextString) bool {
	return false
}


func IsBoolean(name kmipbin.KmipTextString) bool {
	return false
}


func IsInterval(name kmipbin.KmipTextString) bool {
	return false
}


func IsDate(name kmipbin.KmipTextString) bool {
	if name == "Activation Date" ||
		name == "Deactivation Date" {
		return true
	}
	return false
}


func IsTextString(name kmipbin.KmipTextString) bool {
	if name =="Object Group" {
		return true
	}
	return false
}


func IsByteString(name kmipbin.KmipTextString) bool {
	return false
}

func (a *Attribute) Unmarshal(bet *[]byte) {
	var l kmipbin.KmipLength
	l.UnMarshalBin((*bet)[4:8])
	*bet = (*bet)[8:]
	le := kmipbin.PadLength(int(l))
	////////////////////////////////////////////

	b := (*bet)[:le]

	for {
		if len(b) <= 0 {
			break
		}
		tag := strings.ToUpper(hex.EncodeToString(b[:3]))
		switch tag {
		case "42000A":
			var l kmipbin.KmipLength

			l.UnMarshalBin(b[4:8])
			b = b[8:]
			le := kmipbin.PadLength(int(l))
			var uvi kmipbin.KmipTextString
			uvi.UnMarshalBin(b[:le], int(l))
			b = b[le:]
			a.AttributeName = &uvi
		case "42000B":
			switch {
			case IsInt(*a.AttributeName):
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipInt)
				var k kmipbin.KmipInt
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			case IsDate(*a.AttributeName):
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipDate)
				var k kmipbin.KmipDate
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			case IsEnum(*a.AttributeName):
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipEnum)
				var k kmipbin.KmipEnum
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			case IsLong(*a.AttributeName):
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipLongInt)
				var k kmipbin.KmipLongInt
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			case IsBoolean(*a.AttributeName):
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipBoolean)
				var k kmipbin.KmipBoolean
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			case IsInterval(*a.AttributeName):
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipInterval)
				var k kmipbin.KmipInterval
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			case IsBigInt(*a.AttributeName):
				var l kmipbin.KmipLength
				l.UnMarshalBin(b[4:8])
				b = b[8:]
				le := kmipbin.PadLength(int(l))
				a.AttributeValue = new(kmipbin.KmipBigInt)
				var k kmipbin.KmipBigInt
				k.UnMarshalBin(b , int(l))
				a.AttributeValue = k
				b = b[le:]
			case IsTextString(*a.AttributeName):
				var l kmipbin.KmipLength
				l.UnMarshalBin(b[4:8])
				b = b[8:]
				le := kmipbin.PadLength(int(l))
				a.AttributeValue = new(kmipbin.KmipTextString)
				var k kmipbin.KmipTextString
				k.UnMarshalBin(b , int(l))
				a.AttributeValue = k
				b = b[le:]
			case IsByteString(*a.AttributeName):
				var l kmipbin.KmipLength
				l.UnMarshalBin(b[4:8])
				b = b[8:]
				le := kmipbin.PadLength(int(l))
				a.AttributeValue = new(kmipbin.KmipByteString)
				var k kmipbin.KmipByteString
				k.UnMarshalBin(b , int(l))
				a.AttributeValue = k
				b = b[le:]
			case *a.AttributeName == "Application Specific Information":
				

			}
		}
	}

	/////////////////////////////////////////////
	*bet = (*bet)[le:]

}
