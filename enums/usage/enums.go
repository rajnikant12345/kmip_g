package usage

const (
	Sign = 0x00000001

	Verify = 0x00000002

	Encrypt = 0x00000004

	Decrypt = 0x00000008

	WrapKey = 0x00000010

	UnwrapKey = 0x00000020

	Export = 0x00000040

	MACGenerate = 0x00000080

	MACVerify = 0x00000100

	DeriveKey = 0x00000200

	ContentCommitment = 0x00000400

	KeyAgreement = 0x00000800

	CertificateSign = 0x00001000

	CRLSign = 0x00002000

	GenerateCryptogram = 0x00004000

	ValidateCryptogram = 0x00008000

	TranslateEncrypt = 0x00010000

	TranslateDecrypt = 0x00020000

	TranslateWrap = 0x00040000

	TranslateUnwrap = 0x00080000
)
