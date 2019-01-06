package objects

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type X509CertificateIdentifier struct {
	IssuerDistinguishedName *kmipbin.KmipByteString
	CertificateSerialNumber	*kmipbin.KmipByteString
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
	Username               *kmipbin.KmipTextString
	Password               *kmipbin.KmipTextString
	DeviceSerialNumber     *kmipbin.KmipTextString
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
	KeyCompressionTypem    *kmipbin.KmipEnum
	KeyValue               interface{}
	CryptographicAlgorithm *kmipbin.KmipEnum
	CryptographicLength    *kmipbin.KmipInt
	KeywrappingData        *KeyWrappingData
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

type TransparentRSAPrivateKey struct {
	Modulus, PrivateExponent       *kmipbin.KmipBigInt
	PublicExponent, p, Q           *kmipbin.KmipBigInt
	PrimeExponentP, PrimeExponentQ *kmipbin.KmipBigInt
	CRTCoefficient                 *kmipbin.KmipBigInt
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
	Attribute *Attribute
}

type PublicKeyTemplateAttribute struct {
	Name      *Name
	Attribute *Attribute
}

type CommonTemplateAttribute struct {
	Name      *Name
	Attribute *Attribute
}

type TemplateAttribute struct {
	Name      *Name
	Attribute *Attribute
}

type ExtensionInformation struct {
	ExtensionName *kmipbin.KmipTextString
	ExtensionTag  *kmipbin.KmipInt
	ExtensionType *kmipbin.KmipInt
}
