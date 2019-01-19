package state

const (
	PreActive = 0x00000001

	Active = 0x00000002

	Deactivated = 0x00000003

	Compromised = 0x00000004

	Destroyed = 0x00000005

	DestroyedCompromised = 0x00000006
)


//Validity Indicator
const (
	Valid = 0x00000001

	Invalid = 0x00000002

	Unknown = 0x00000003
)