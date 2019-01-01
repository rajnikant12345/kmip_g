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
	Uid                            string                          `xml:",omitempty"`
	Name                           Name                            `xml:",omitempty"`
	ObjectType                     int                             `xml:",omitempty"` //enum
	CryptographicAlgo              int                             `xml:",omitempty"` //enum
	CryptoGraphicLength            int                             `xml:",omitempty"` //enum
	CryptoParams                   CryptoParams                    `xml:",omitempty"`
	CryptoDomainParams             CryptoDomainParams              `xml:",omitempty"`
	CertificateType                int                             `xml:",omitempty"` //enum
	CertificacteLength             int                             `xml:",omitempty"`
	X509CertificateId              X509CertificateId               `xml:",omitempty"`
	X509CertificateSub             X509CertificateSub              `xml:",omitempty"`
	X509CertificateIssuer          X509CertificateIssuer           `xml:",omitempty"`
	CertificateId                  CertificateId                   `xml:",omitempty"`
	CertificateSub                 CertificateSub                  `xml:",omitempty"`
	CertificateIssuer              CertificateIssuer               `xml:",omitempty"`
	DigitalSigAlgo                 int                             `xml:",omitempty"` //enum
	Digest                         Digest                          `xml:",omitempty"`
	OperationPolicyName            string                          `xml:",omitempty"`
	CryptoUsageMask                int                             `xml:",omitempty"`
	LeaseTime                      int                             `xml:",omitempty"` //date
	UsageLimit                     UsageLimit                      `xml:",omitempty"`
	State                          int                             `xml:",omitempty"` //enum
	Initialdate                    int64                           `xml:",omitempty"`
	Activationdate                 int64                           `xml:",omitempty"`
	ProcessStartDate               int64                           `xml:",omitempty"`
	ProcessStopDate                int64                           `xml:",omitempty"`
	Deactivationdate               int64                           `xml:",omitempty"`
	DestroyDate                    int64                           `xml:",omitempty"`
	CompromiseOccuranceDate        int64                           `xml:",omitempty"`
	CompromiseDate                 int64                           `xml:",omitempty"`
	RevocationReason               RevocationReason                `xml:",omitempty"`
	ArchiveDate                    int64                           `xml:",omitempty"`
	ObjectGroup                    string                          `xml:",omitempty"`
	Fresh                          int64                           `xml:",omitempty"` //bool
	Link                           Link                            `xml:",omitempty"`
	ApplicationSpecificInformation ApplicationSpecificInformation  `xml:",omitempty"`
	ContactInformation             string                          `xml:",omitempty"`
	LastChangeDate                 int64                           `xml:",omitempty"`
	CustomAttribute                map[string]CustomAttributeValue `xml:",omitempty"`
	AlternativeName                AlternativeName                 `xml:",omitempty"`
	KeyValuePresent                int64                           `xml:",omitempty"` //bool
	KeyValueLocation               KeyValueLocation                `xml:",omitempty"`
	OriginalCreationDate           int64                           `xml:",omitempty"`
}

type Attribute struct {
	AttributeName  *kmipbin.KmipTextString `kmip:"42000A"`
	AttributeValue interface{}             `kmip:"42000B" multi:"true"`
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
			if *a.AttributeName == "Activation Date" || *a.AttributeName == "Deactivation Date" {
				b = b[8:]
				a.AttributeValue = new(kmipbin.KmipDate)
				var k kmipbin.KmipDate
				k.UnMarshalBin(b)
				a.AttributeValue = k
				b = b[8:]
			}

		}
	}

	/////////////////////////////////////////////
	*bet = (*bet)[le:]

}
