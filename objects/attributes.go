package objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
	"strings"
	"fmt"
	"encoding/hex"
)

type Attribute struct {
	AttributeName  *kmipbin.KmipTextString
	AttributeIndex *kmipbin.KmipInt
	AttributeValue interface{}
	data           []byte	`ignore:"true"`
	function       func(interface{} , []byte)	`ignore:"true"`
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


func (a *Attribute) Unpack(b []byte) {
	b = b[8:]
	//for {
	if len(b) <= 0 {
		return
	}
	tty := kmipbin.KmipTagType{}
	var l kmipbin.KmipLength
	l.UnMarshalBin(b[4:8])
	le := kmipbin.PadLength(int(l))
	b = b[8:]
	var AttributeName kmipbin.KmipTextString
	AttributeName.UnMarshalBin(b[:le],int(l))
	a.AttributeName = &AttributeName
	b = b[le:]

	/////

	tty.UnMarshalBin(b[:4])
	l.UnMarshalBin(b[4:8])
	le = kmipbin.PadLength(int(l))
	b = b[8:]

	a.CreateAttribute()

	switch tty.Type {
	case 2,3,5,6,9,10:
		if strings.HasPrefix(string(*a.AttributeName),"x-") {
			val := MakeType(tty.Type).(kmipbin.BaseMarshal)
			val.UnMarshalBin(b[:le])
			a.AttributeValue = val
		}else {
			a.AttributeValue.(kmipbin.BaseMarshal).UnMarshalBin(b)
			b = b[le:]
		}

	case 4,7,8:
		if strings.HasPrefix(string(*a.AttributeName),"x-") {
			var str kmipbin.KmipTextString
			str.UnMarshalBin(b[:le] , int(l))
			a.AttributeValue = &str
		}else {
			a.AttributeValue.(kmipbin.BaseMarshalString).UnMarshalBin(b[:le] , int(l))
			b = b[le:]
			//fmt.Println(*a.AttributeName)
		}
	default:
		a.function(a.AttributeValue , b)
	}

	b = b[len(b):]
	//}
}


func (a *Attribute) Unmarshal(bet *[]byte , f func(interface{} , []byte)) {
	var l kmipbin.KmipLength
	l.UnMarshalBin((*bet)[4:8])
	le := kmipbin.PadLength(int(l))
	////////////////////////////////////////////

	a.data = make([]byte, 8+le)
	copy(a.data, (*bet)[:8+le])

	/////////////////////////////////////////////
	*bet = (*bet)[8+le:]

	/////////////////////////////////////////////////////
	a.function = f

	a.Unpack(a.data)
}

func (a *Attribute) Marshal(f func(interface{}) []byte ) []byte {
	fmt.Println("============================")
	b := f(a)
	fmt.Println(hex.EncodeToString(b))
	fmt.Println("============================")
	return b
}

func (a *Attribute) CreateAttribute() {

	if a.AttributeName == nil {
		return
	}
	if *a.AttributeName == "Unique Identifier" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Name" {
		a.AttributeValue = new(Name)
	}
	if *a.AttributeName == "Object Type" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Cryptographic Algorithm" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Cryptographic Length" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Cryptographic Parameters" {
		a.AttributeValue = new(CryptographicParameters)
	}
	if *a.AttributeName == "Cryptographic Domain Parameters" {
		a.AttributeValue = new(CryptographicDomainParameters)
	}
	if *a.AttributeName == "Certificate Type" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Certificate Length" {
		a.AttributeValue = new(kmipbin.KmipInt)
	}
	if *a.AttributeName == "X.509 Certificate Identifier" {
		a.AttributeValue = new(X509CertificateIdentifier)
	}
	if *a.AttributeName == "X.509 Certificate Subject" {
		a.AttributeValue = new(X509CertificateSubject)
	}
	if *a.AttributeName == "X.509 Certificate Issuer" {
		a.AttributeValue = new(X509CertificateIssuer)
	}
	if *a.AttributeName == "Certificate Identifier" {
		a.AttributeValue = new(CertificateIdentifier)
	}
	if *a.AttributeName == "Certificate Subject" {
		a.AttributeValue = new(CertificateSubject)
	}
	if *a.AttributeName == "Certificate Issuer" {
		a.AttributeValue = new(CertificateIssuer)
	}
	if *a.AttributeName == "Digital Signature Algorithm" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Digest" {
		a.AttributeValue = new(Digest)
	}
	if *a.AttributeName == "Operation Policy Name" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Cryptographic Usage Mask" {
		a.AttributeValue = new(kmipbin.KmipInt)
	}
	if *a.AttributeName == "Lease Time" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Usage Limits" {
		a.AttributeValue = new(UsageLimits)
	}
	if *a.AttributeName == "State" {
		a.AttributeValue = new(kmipbin.KmipEnum)
	}
	if *a.AttributeName == "Initial Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Activation Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Process Start Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Protect Stop Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Deactivation Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Destroy Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Compromise Occurrence Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Compromise Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Revocation Reason" {
		a.AttributeValue = new(RevocationReason)
	}
	if *a.AttributeName == "Archive Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	if *a.AttributeName == "Object Group" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Fresh" {
		a.AttributeValue = new(kmipbin.KmipBoolean)
	}
	if *a.AttributeName == "Link" {
		a.AttributeValue = new(Link)
	}
	if *a.AttributeName == "Application Specific Information" {
		a.AttributeValue = new(ApplicationSpecificInformation)
	}
	if *a.AttributeName == "Contact Information" {
		a.AttributeValue = new(kmipbin.KmipTextString)
	}
	if *a.AttributeName == "Last Change Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
	/*if strings.HasPrefix(string(*a.AttributeName),"x-") {
		a.AttributeValue = make([]byte,1000)
	}*/
	if *a.AttributeName == "Alternative Name" {
		a.AttributeValue = new(AlternativeName)
	}
	if *a.AttributeName == "Key Value Present" {
		a.AttributeValue = new(kmipbin.KmipBoolean)
	}
	if *a.AttributeName == "Key Value Location" {
		a.AttributeValue = new(KeyValueLocation)
	}
	if *a.AttributeName == "Original Creation Date" {
		a.AttributeValue = new(kmipbin.KmipDate)
	}
}
