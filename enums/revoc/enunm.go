package revoc


// Revocation reasom
const (
	Unspecified = 0x00000001

	KeyCompromise = 0x00000002

	CACompromise = 0x00000003

	AffiliationChanged = 0x00000004

	Superseded = 0x00000005

	CessationofOperation = 0x00000006

	PrivilegeWithdrawn = 0x00000007
)

