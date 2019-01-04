package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type SecretData struct {
	SecretDataType *kmipbin.KmipEnum
	KeyBlock       *KeyBlock
}


type PrivateKey struct {
	KeyBlock       *KeyBlock
}

type PublicKey struct {
	KeyBlock       *KeyBlock
}
