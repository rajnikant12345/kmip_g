package resultreason

//ResultReasonEnumeration
const (
	ItemNotFound                     = 1
	ResponseTooLarge                 = 2
	AuthenticationNotSuccessful      = 3
	InvalidMessage                   = 4
	OperationNotSupported            = 5
	MissingData                      = 6
	InvalidField                     = 7
	FeatureNotSupported              = 8
	OperationCanceledByRequester     = 9
	CryptographicFailure             = 10
	IllegalOperation                 = 11
	PermissionDenied                 = 12
	ObjectArchived                   = 13
	IndexOutofBounds                 = 14
	ApplicationNamespaceNotSupported = 15
	KeyFormatTypeNotSupported        = 16
	KeyCompressionTypeNotSupported   = 17
	EncodingOptionError              = 18
	KeyValueNotPresent               = 19
	AttestationRequired              = 20
	AttestationFailed                = 21
	GeneralFailure                   = 0x100
)

const (
	Continue = 0x00000001

	Stop = 0x00000002

	Undo = 0x00000003
)

//Encoding Option Enumeration

const (
	NoEncoding = 0x00000001

	TTLVEncoding = 0x00000002
)
