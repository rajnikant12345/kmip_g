package objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type X509CertificateIdentifier struct {
	IssuerDistinguishedName *kmipbin.KmipByteString
	CertificateSerialNumber *kmipbin.KmipByteString
}

type X509CertificateSubject struct {
	SubjectDistinguishedName *kmipbin.KmipByteString
	SubjectAlternativeName   *kmipbin.KmipByteString
}

type X509CertificateIssuer struct {
	IssuerDistinguishedName *kmipbin.KmipByteString
	IssuerAlternativeName   *kmipbin.KmipByteString
}

type CertificateIdentifier struct {
	Issuer       *kmipbin.KmipTextString
	SerialNumber *kmipbin.KmipTextString
}

type CertificateSubject struct {
	CertificateSubjectDistinguishedName *kmipbin.KmipTextString
	CertificateSubjectAlternativeName   *kmipbin.KmipTextString
}

type CertificateIssuer struct {
	CertificateIssuerDistinguishedName *kmipbin.KmipTextString
	CertificateIssuerAlternativeName   *kmipbin.KmipTextString
}

type Name struct {
	NameValue *kmipbin.KmipTextString
	NameType  *kmipbin.KmipEnum
}

type Nonce struct {
	NonceID    *kmipbin.KmipByteString
	NonceValue *kmipbin.KmipByteString
}

type CryptographicDomainParameters struct {
	Qlength          *kmipbin.KmipInt
	RecommendedCurve *kmipbin.KmipEnum
}

type CredentialValue struct {
	DeviceSerialNumber     *kmipbin.KmipTextString
	Username               *kmipbin.KmipTextString
	Password               *kmipbin.KmipTextString
	DeviceIdentifier       *kmipbin.KmipTextString
	NetworkIdentifier      *kmipbin.KmipTextString
	MachineIdentifier      *kmipbin.KmipTextString
	MediaIdentifier        *kmipbin.KmipTextString
	Nonce                  *Nonce
	AttestationType        *kmipbin.KmipEnum
	AttestationMeasurement *kmipbin.KmipByteString
	AttestationAssertion   *kmipbin.KmipByteString
}

type Credential struct {
	CredentialType  *kmipbin.KmipEnum
	CredentialValue *CredentialValue
}

type KeyBlock struct {
	KeyFormatType          *kmipbin.KmipEnum
	KeyCompressionType     *kmipbin.KmipEnum
	KeyValue               interface{}
	CryptographicAlgorithm *kmipbin.KmipEnum
	CryptographicLength    *kmipbin.KmipInt
	KeywrappingData        *KeyWrappingData
}

func (k *KeyBlock) Unmarshal(bet *[]byte, f func(interface{}, []byte)) {
	var l kmipbin.KmipLength
	l.UnMarshalBin((*bet)[4:8])
	le := kmipbin.PadLength(int(l))
	////////////////////////////////////////////

	//	k.RawData = make([]byte, 8+le)
	//	copy(k.RawData, (*bet)[:8+le])

	/////////////////////////////////////////////
	*bet = (*bet)[8+le:]
}

type KeyValue struct {
	KeyMaterial interface{}
	Attribute   *[]Attribute
}

type KeyWrappingData struct {
	WrappingMethod             *kmipbin.KmipEnum
	EncryptionKeyInformation   *EncryptionKeyInformation
	MACSignatureKeyInformation *MACSignatureKeyInformation
	MACSignature               *kmipbin.KmipByteString
	IVCounterNonce             *kmipbin.KmipByteString
	EncodingOption             *kmipbin.KmipEnum
}

type CryptographicParameters struct {
	BlockCipherMode           *kmipbin.KmipEnum
	PaddingMethod             *kmipbin.KmipEnum
	HashingAlgorithm          *kmipbin.KmipEnum
	KeyRoleType               *kmipbin.KmipEnum
	DigitalSignatureAlgorithm *kmipbin.KmipEnum
	CryptographicAlgorithm    *kmipbin.KmipEnum
	RandomIV                  *kmipbin.KmipBoolean
	IVLength                  *kmipbin.KmipInt
	TagLength                 *kmipbin.KmipInt
	FixedFieldLength          *kmipbin.KmipInt
	InvocationFieldLength     *kmipbin.KmipInt
	CounterLength             *kmipbin.KmipInt
	InitialCounterValue       *kmipbin.KmipInt
}

// MACSignatureKeyInformation 2.1.5 Table 11
type MACSignatureKeyInformation struct {
	UniqueIdentifier        *kmipbin.KmipTextString
	CryptographicParameters *CryptographicParameters
}

// EncryptionKeyInformation 2.1.5 Table 10
type EncryptionKeyInformation struct {
	UniqueIdentifier        *kmipbin.KmipTextString
	CryptographicParameters *CryptographicParameters
}

type KeyWrappingSpecification struct {
	WrappingMethod             *kmipbin.KmipEnum
	EncryptionKeyInformation   *EncryptionKeyInformation
	MACSignatureKeyInformation *MACSignatureKeyInformation
	AttributeName              *kmipbin.KmipTextString
	EncodingOption             *kmipbin.KmipEnum
}

type TransparentSymmetricKey struct {
	Key *kmipbin.KmipByteString
}

type TransparentDSAPrivateKey struct {
	P, Q, G, X *kmipbin.KmipBigInt
}

type TransparentDSAPublicKey struct {
	P, Q, G, Y *kmipbin.KmipBigInt
}

type TransparentKey struct {
	Modulus, PrivateExponent       *kmipbin.KmipBigInt
	PublicExponent, P, Q           *kmipbin.KmipBigInt
	PrimeExponentP, PrimeExponentQ *kmipbin.KmipBigInt
	CRTCoefficient                 *kmipbin.KmipBigInt
	X                              *kmipbin.KmipBigInt
	G, J, Y                        *kmipbin.KmipBigInt
	RecommendedCurve               *kmipbin.KmipEnum
	D                              *kmipbin.KmipBigInt
	QString                        *kmipbin.KmipByteString
}

type TransparentRSAPublicKey struct {
	Modulus        *kmipbin.KmipBigInt
	PublicExponent *kmipbin.KmipBigInt
}

type TransparentDHPrivateKey struct {
	P, Q, G, J, X *kmipbin.KmipBigInt
}

type TransparentDHPublicKey struct {
	P, Q, G, J, Y *kmipbin.KmipBigInt
}

type TransparentECDSAPrivateKey struct {
	RecommendedCurve *kmipbin.KmipEnum
	D                *kmipbin.KmipBigInt
}

type TransparentECDSAPublicKey struct {
	RecommendedCurve *kmipbin.KmipEnum
	QString          *kmipbin.KmipByteString
}

type TransparentECDHPrivateKey struct {
	RecommendedCurve *kmipbin.KmipEnum
	D                *kmipbin.KmipBigInt
}

type TransparentECDHPublicKey struct {
	RecommendedCurve *kmipbin.KmipEnum
	QString          *kmipbin.KmipByteString
}

type TransparentECMQVPrivateKey struct {
	RecommendedCurve *kmipbin.KmipEnum
	D                *kmipbin.KmipBigInt
}

type TransparentECMQVPublicKey struct {
	RecommendedCurve *kmipbin.KmipEnum
	QString          *kmipbin.KmipByteString
}

type PrivateKeyTemplateAttribute struct {
	Name      *Name
	Attribute []*Attribute
}

type PublicKeyTemplateAttribute struct {
	Name      *Name
	Attribute []*Attribute
}

type CommonTemplateAttribute struct {
	Name      *Name
	Attribute []*Attribute
}

type TemplateAttribute struct {
	Name      *Name
	Attribute []*Attribute
}

type ExtensionInformation struct {
	ExtensionName *kmipbin.KmipTextString
	ExtensionTag  *kmipbin.KmipInt
	ExtensionType *kmipbin.KmipInt
}

type Digest struct {
	HashingAlgorithm *kmipbin.KmipEnum
	DigestValue      *kmipbin.KmipByteString
	KeyFormatType    *kmipbin.KmipEnum
}

type UsageLimits struct {
	UsageLimitsTotal *kmipbin.KmipLongInt
	UsageLimitsCount *kmipbin.KmipLongInt
	UsageLimitsUnit  *kmipbin.KmipEnum
}

type RevocationReason struct {
	RevocationReasonCode *kmipbin.KmipEnum
	RevocationMessage    *kmipbin.KmipTextString
}

type Link struct {
	LinkType               *kmipbin.KmipEnum
	LinkedObjectIdentifier *kmipbin.KmipTextString
}

type ApplicationSpecificInformation struct {
	ApplicationNamespace *kmipbin.KmipTextString
	ApplicationData      *kmipbin.KmipTextString
}

type AlternativeName struct {
	AlternativeNameValue *kmipbin.KmipTextString
	AlternativeNameType  *kmipbin.KmipEnum
}

type KeyValueLocation struct {
	KeyValueLocationValue *kmipbin.KmipTextString
	KeyValueLocationType  *kmipbin.KmipEnum
}

type DerivationParameters struct {
	CryptographicParameters *CryptographicParameters
	InitializationVector    *kmipbin.KmipByteString
	DerivationData          *kmipbin.KmipByteString
}
