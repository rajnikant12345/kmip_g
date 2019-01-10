package resultreason

//resultstatus
const (
	Success          = 0
	OperationFailed  = 1
	OperationPending = 2
	OperationUndone  = 3
)


///Cancellation Result

const (
	Canceled = 0x00000001

	UnabletoCancel = 0x00000002

	Completed = 0x00000003

	Failed = 0x00000004

	Unavailable = 0x00000005
)
