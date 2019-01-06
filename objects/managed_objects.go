package objects

import "github.com/rajnikant12345/kmip_g/kmipbin"

type Certificate struct {
	CertificateType  *kmipbin.KmipEnum
	CertificateValue *kmipbin.KmipByteString
}

type SymmetricKey struct {
	KeyBlock *KeyBlock
}

type PublicKey struct {
	KeyBlock *KeyBlock
}

type PrivateKey struct {
	KeyBlock *KeyBlock
}

type SplitKey struct {
	SplitKeyParts     *kmipbin.KmipInt
	KeyPartIdentifier *kmipbin.KmipInt
	SplitKeyThreshold *kmipbin.KmipInt
	SplitKeyMethod    *kmipbin.KmipEnum
	PrimeFieldSize    *kmipbin.KmipBigInt
	KeyBlock          *KeyBlock
}

type Template struct {
	Attribute []*Attribute
}

type SecretData struct {
	SecretDataType *kmipbin.KmipEnum
	KeyBlock       *KeyBlock
}

type OpaqueObject struct {
	OpaqueDataType  *kmipbin.KmipEnum
	OpaqueDataValue *kmipbin.KmipByteString
}

type PGPKey struct {
	PGPKeyVersion *kmipbin.KmipInt
	KeyBlock      *KeyBlock
}
